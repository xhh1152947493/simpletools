package sink

import "simpletools/internal/defs"

type EventSink struct {
	SinkQ             chan interface{}
	exitAfterPakCount int64
}

func (self *EventSink) SendSinkQ(s defs.EEventSinkSource, in interface{}) {
	self.SinkQ <- in
}

func (self *EventSink) CatchException(recoverResult, in interface{}, desc string) {
	if recoverResult == nil {
		return
	}

}

func (self *EventSink) ConsumePak(in interface{}) {
	//defer func() {
	//	self.CatchException(recover(), in, "consume pak catch can recover error")
	//}()

	//switch input := in.(type) {
	//
	//}
}
