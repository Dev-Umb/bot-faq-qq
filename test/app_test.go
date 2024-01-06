package test

import (
	conf "bot-faq-qq/config"
	"bot-faq-qq/internal/pkg/handler"
	"context"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"log"
	"testing"
	"time"
)

func TestBot(t *testing.T) {
	botToken := token.BotToken(conf.AppID, conf.Token)
	//api := botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
	// 沙箱环境
	api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)
	ctx := context.Background()
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Printf("%+v, err:%v", ws, err)
	}

	process := handler.Processor{Api: api}
	intent := websocket.RegisterHandlers(
		// at 机器人事件，目前是在这个事件处理中有逻辑，会回消息，其他的回调处理都只把数据打印出来，不做任何处理
		process.ATMessageEventHandler(),
		// 如果想要捕获到连接成功的事件，可以实现这个回调
		process.ReadyHandler(),
		// 连接关闭回调
		process.ErrorNotifyHandler(),
		// 频道事件
		process.GuildEventHandler(),
		// 成员事件
		process.MemberEventHandler(),
		// 子频道事件
		process.ChannelEventHandler(),
		//// 私信，目前只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		//process.DirectMessageHandler(),
		//// 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		//process.CreateMessageHandler(),
		// 互动事件
		process.InteractionHandler(),
		// 发帖事件
		process.ThreadEventHandler(),
	)
	// 指定需要启动的分片数为 2 的话可以手动修改 wsInfo
	if err = botgo.NewSessionManager().Start(ws, botToken, &intent); err != nil {
		log.Fatalln(err)
	}
}
