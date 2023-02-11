package executor

import (
	"context"

	"www.velocidex.com/golang/velociraptor/actions"
	crypto_proto "www.velocidex.com/golang/velociraptor/crypto/proto"
	"www.velocidex.com/golang/velociraptor/responder"
)

type _TestExecutor struct {
	flow_manager  *responder.FlowManager
	event_manager *actions.EventTable
}

func NewTestExecutor() *_TestExecutor {
	return &_TestExecutor{
		event_manager: &actions.EventTable{},
	}
}

func NewClientExecutorForTests() *ClientExecutor {
	return &ClientExecutor{
		Outbound:      make(chan *crypto_proto.VeloMessage),
		Inbound:       make(chan *crypto_proto.VeloMessage),
		event_manager: &actions.EventTable{},
	}
}

func (self *_TestExecutor) FlowManager() *responder.FlowManager {
	return nil
}

func (self *_TestExecutor) EventManager() *actions.EventTable {
	return nil
}

func (self *_TestExecutor) ClientId() string {
	return ""
}

func (self *_TestExecutor) ReadFromServer() *crypto_proto.VeloMessage {
	return nil
}
func (self *_TestExecutor) SendToServer(message *crypto_proto.VeloMessage) {}
func (self *_TestExecutor) ProcessRequest(ctx context.Context,
	message *crypto_proto.VeloMessage) {
}
func (self *_TestExecutor) ReadResponse() <-chan *crypto_proto.VeloMessage {
	return nil
}
