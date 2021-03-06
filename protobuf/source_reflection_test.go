package protobuf

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func TestReflection(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer ln.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	go s.Serve(ln)

	// Ensure that all streams are closed by the end of the test.
	defer s.GracefulStop()

	source, err := NewDescriptorProviderReflection(ReflectionArgs{
		Timeout: time.Second,
		Peers:   []string{ln.Addr().String()},
	})
	require.NoError(t, err)
	require.NotNil(t, source)

	// Close the streaming reflect call to ensure GracefulStop doesn't block.
	defer source.Close()

	result, err := source.FindSymbol("grpc.reflection.v1alpha.ServerReflectionRequest")
	assert.NoError(t, err)
	assert.NotNil(t, result)

	result, err = source.FindSymbol("wat")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Symbol not found: wat")
	assert.Nil(t, result)
}

func TestReflectionWithProtocolInPeer(t *testing.T) {
	source, err := NewDescriptorProviderReflection(ReflectionArgs{
		Timeout: time.Second,
		Peers:   []string{"grpc://127.0.0.1:12345"},
	})
	assert.Nil(t, source)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "peer contains scheme")
}

func TestReflectionMultiplePeers(t *testing.T) {
	listenRefuser, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err, "failed to listen on a port")
	defer listenRefuser.Close()

	noListen, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err, "failed to listen on a port")
	noListen.Close()

	got, err := NewDescriptorProviderReflection(ReflectionArgs{
		Timeout: time.Second,
		Peers:   []string{noListen.Addr().String(), listenRefuser.Addr().String()},
	})

	require.NoError(t, err)
	require.NotNil(t, got)

	go func() {
		conn, _ := listenRefuser.Accept()
		conn.Close()
	}()
	_, err = got.FindSymbol("some-symbol")
	require.Error(t, err)
}

func TestReflectionClosedPort(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err, "failed to listen on a port")
	ln.Close()

	got, err := NewDescriptorProviderReflection(ReflectionArgs{
		Timeout: time.Second,
		Peers:   []string{ln.Addr().String()},
	})

	assert.Contains(t, err.Error(), "could not reach reflection server")
	assert.Nil(t, got)
}
