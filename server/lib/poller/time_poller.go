package poller

import (
	"fmt"
	"time"
)

type TimePoller struct {
	lastTS int64
}

func (self *TimePoller) String() string {
	return fmt.Sprintf("timepoller info. lastTS:%d time:%s", self.lastTS, time.Unix(self.lastTS, 0).Format("2006-01-02 15:04:05"))
}

func NewTimePoller(now int64) *TimePoller {
	poller := &TimePoller{}
	poller.Init(now)
	return poller
}

func (self *TimePoller) Init(now int64) {
	self.lastTS = now
}

func (self *TimePoller) PassedTime(now int64, period int64) bool {
	delta := now - self.lastTS
	if delta < 0 {
		delta = -delta
	}
	if delta < period {
		return false
	}

	self.lastTS = now
	return true
}
