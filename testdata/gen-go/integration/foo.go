// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package integration

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type Foo interface {
	// Parameters:
	//  - Arg
	Bar(arg int32) (r int32, err error)
}

type FooClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewFooClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FooClient {
	return &FooClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewFooClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FooClient {
	return &FooClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Arg
func (p *FooClient) Bar(arg int32) (r int32, err error) {
	if err = p.sendBar(arg); err != nil {
		return
	}
	return p.recvBar()
}

func (p *FooClient) sendBar(arg int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("bar", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FooBarArgs{
		Arg: arg,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *FooClient) recvBar() (value int32, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "bar" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "bar failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "bar failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "bar failed: invalid message type")
		return
	}
	result := FooBarResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.NotFound != nil {
		err = result.NotFound
		return
	}
	value = result.GetSuccess()
	return
}

type FooProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Foo
}

func (p *FooProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *FooProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *FooProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewFooProcessor(handler Foo) *FooProcessor {

	self2 := &FooProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["bar"] = &fooProcessorBar{handler: handler}
	return self2
}

func (p *FooProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type fooProcessorBar struct {
	handler Foo
}

func (p *fooProcessorBar) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FooBarArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("bar", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FooBarResult{}
	var retval int32
	var err2 error
	if retval, err2 = p.handler.Bar(args.Arg); err2 != nil {
		switch v := err2.(type) {
		case *NotFound:
			result.NotFound = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing bar: "+err2.Error())
			oprot.WriteMessageBegin("bar", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("bar", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Arg
type FooBarArgs struct {
	Arg int32 `thrift:"arg,1" db:"arg" json:"arg"`
}

func NewFooBarArgs() *FooBarArgs {
	return &FooBarArgs{}
}

func (p *FooBarArgs) GetArg() int32 {
	return p.Arg
}
func (p *FooBarArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FooBarArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Arg = v
	}
	return nil
}

func (p *FooBarArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("bar_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FooBarArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("arg", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:arg: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Arg)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.arg (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:arg: ", p), err)
	}
	return err
}

func (p *FooBarArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FooBarArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - NotFound
type FooBarResult struct {
	Success  *int32    `thrift:"success,0" db:"success" json:"success,omitempty"`
	NotFound *NotFound `thrift:"notFound,1" db:"notFound" json:"notFound,omitempty"`
}

func NewFooBarResult() *FooBarResult {
	return &FooBarResult{}
}

var FooBarResult_Success_DEFAULT int32

func (p *FooBarResult) GetSuccess() int32 {
	if !p.IsSetSuccess() {
		return FooBarResult_Success_DEFAULT
	}
	return *p.Success
}

var FooBarResult_NotFound_DEFAULT *NotFound

func (p *FooBarResult) GetNotFound() *NotFound {
	if !p.IsSetNotFound() {
		return FooBarResult_NotFound_DEFAULT
	}
	return p.NotFound
}
func (p *FooBarResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FooBarResult) IsSetNotFound() bool {
	return p.NotFound != nil
}

func (p *FooBarResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FooBarResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *FooBarResult) ReadField1(iprot thrift.TProtocol) error {
	p.NotFound = &NotFound{}
	if err := p.NotFound.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.NotFound), err)
	}
	return nil
}

func (p *FooBarResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("bar_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FooBarResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *FooBarResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetNotFound() {
		if err := oprot.WriteFieldBegin("notFound", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:notFound: ", p), err)
		}
		if err := p.NotFound.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.NotFound), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:notFound: ", p), err)
		}
	}
	return err
}

func (p *FooBarResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FooBarResult(%+v)", *p)
}