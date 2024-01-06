package handler

import (
	"fmt"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
)

// ThreadEventHandler 论坛主贴事件
func (processor Processor) ThreadEventHandler() event.ThreadEventHandler {
	return func(event *dto.WSPayload, data *dto.WSThreadData) error {
		fmt.Println(event, data)
		return nil
	}
}
