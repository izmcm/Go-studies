package invoker

import (
	"fmt"
	"impl"
	"marshaller"
	"miop"
	// "proxies"
	"naming"
	"shared"
	"srh"
)

type CalculatorInvoker struct {
	instCalc impl.Calculadora
}

func NewCalculatorInvoker() CalculatorInvoker {
	return CalculatorInvoker{instCalc: impl.Calculadora{}}
}

func (calcInv *CalculatorInvoker) Invoke() {
	srhImpl := srh.SRH{ServerHost: shared.CALCULATOR_IP, ServerPort: shared.CALCULATOR_PORT, Conn: nil}
	marshallerImpl := marshaller.Marshaller{}
	// calculatorImpl := impl.Calculadora{}
	calculatorImpl := calcInv.instCalc
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	for {
		rcvMsgBytes := srhImpl.Receive()

		fmt.Printf("did receive: ")
		fmt.Println(rcvMsgBytes)

		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation
		// fmt.Println("operation detected: " + operation)

		// TODO: talvez esteja errado
		// TODO: fazer essas operações independentes da estrutura de operação
		switch operation {
		case "Add":
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Add(_p1, _p2)
		case "Sub":
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Sub(_p1, _p2)
		case "Mul":
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Mul(_p1, _p2)
		case "Div":
			_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			_p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			replParams[0] = calculatorImpl.Div(_p1, _p2)
		}
		// fmt.Println(replParams[0])

		repHeader := miop.ReplyHeader{Context: "", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		srhImpl.Send(msgToClientBytes)
		// (*srhImpl.Conn).Close()
	}
}

type NameServerInvoker struct {
	Register *naming.NamingService
}

func NewNameServerInvoker(service naming.NamingService) NameServerInvoker {
	return NameServerInvoker{Register: &service}
}

func (nameInv NameServerInvoker) Invoke() {
	fmt.Println("reading client data")
	srhImpl := srh.SRH{ServerHost: shared.NAMESERVER_IP, ServerPort: shared.NAMESERVER_PORT, Conn: nil}
	marshallerImpl := marshaller.Marshaller{}
	// calculatorImpl := impl.Calculadora{}
	service := *nameInv.Register
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)
	fmt.Println("reading client data")

	for {
		rcvMsgBytes := srhImpl.Receive()
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation
		// fmt.Println("operation detected: " + operation)

		// TODO: talvez esteja errado
		// TODO: fazer essas operações independentes da estrutura de operação
		switch operation {
		case "GetServer":
			_name := miopPacketRequest.Bd.ReqBody.Body[0].(string)
			replParams[0] = service.Lookup(_name)
			fmt.Println(_name)
		}
		fmt.Println(replParams[0])

		repHeader := miop.ReplyHeader{Context: "", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		srhImpl.Send(msgToClientBytes)
		// (*srhImpl.Conn).Close()
	}
}
