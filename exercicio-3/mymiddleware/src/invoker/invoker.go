package invoker

import (
	// "fmt"
	"marshaller"
	"miop"
	"proxies"
	"shared"
	"srh"
)

type CalculatorInvoker struct{}

func NewCalculatorInvoker() CalculatorInvoker {
	// TODO
}

func (CalculatorInvoker) Invoke() {
	srhImpl := srh.SRH{ServerHost: "localhost", ServerPort: shared.CALCULATOR_PORT}
	marshallerImpl := marshaller.Marshaller{}
	calculatorImpl := impl.Calculadora{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	for {
		rcvMsgBytes := srhImpl.Receive()

		miopPacketRequest := mashallerImpl.Unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		// TODO: talvez esteja errado
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

		repHeader := miop.ReplyHeader{Context: "", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status: 1}
		RepBody := miop.ReplyBody{OperationResult: replParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		srhImpl.Send(msgToClientBytes)
	}
}
