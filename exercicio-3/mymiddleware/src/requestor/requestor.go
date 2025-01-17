package requestor

import (
	"crh"
	"fmt"
	"marshaller"
	"miop"
	"shared"
)

type Requestor struct{}

func (Requestor) Invoke(inv shared.Invocation) []interface{} {
	marshallerInst := marshaller.Marshaller{}

	crhInst := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}
	fmt.Printf("instance: ")
	fmt.Println(crhInst)

	// Create request packet
	// TODO: Randomize the requestId or make an algorithm to make it work
	reqHeader := miop.RequestHeader{Context: "Context", RequestId: 1000, ResponseExpected: true, ObjectKey: 2000, Operation: inv.Request.Op}
	reqBody := miop.RequestBody{Body: inv.Request.Params}
	header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: shared.MIOP_REQUEST}
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	miopPacketRequest := miop.Packet{Hdr: header, Bd: body}

	// Serialize Request packet
	msgToClientBytes := marshallerInst.Marshall(miopPacketRequest)

	// Send Request Packet and receive reply
	fmt.Println("preparando pra enviar")
	msgFromServerBytes := crhInst.SendReceive(msgToClientBytes)
	fmt.Println("recebidos")
	miopPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)

	r := miopPacketReply.Bd.RepBody.OperationResult

	fmt.Printf("\ntype: %T\n", r)
	fmt.Println("value:\n", r)

	return r.([]interface{})
}
