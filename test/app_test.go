package test

import (
	conf "bot-faq-qq/config"
	"context"
	"fmt"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"log"
	"testing"
	"time"
)

func TestBot(t *testing.T) {
	token := token.BotToken(conf.AppID, conf.Token)
	api := botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
	ctx := context.Background()
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Printf("%+v, err:%v", ws, err)
	}

	// 监听哪类事件就需要实现哪类的 handler，定义：websocket/event_handler.go
	var atMessage = func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		fmt.Println(event, data)
		return nil
	}
	intent := websocket.RegisterHandlers(atMessage)
	// 启动 session manager 进行 ws 连接的管理，如果接口返回需要启动多个 shard 的连接，这里也会自动启动多个
	botgo.NewSessionManager().Start(ws, token, &intent)

}
