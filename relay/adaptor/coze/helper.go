package coze

import "github.com/9688101/HX/relay/adaptor/coze/constant/event"

func event2StopReason(e *string) string {
	if e == nil || *e == event.Message {
		return ""
	}
	return "stop"
}
