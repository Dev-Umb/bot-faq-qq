package handler

import (
	"context"
	"log"

	"github.com/tencent-connect/botgo/dto"
)

func (processor Processor) setEmoji(ctx context.Context, channelID string, messageID string) {
	err := processor.Api.CreateMessageReaction(
		ctx, channelID, messageID, dto.Emoji{
			ID:   "307",
			Type: 1,
		},
	)
	if err != nil {
		log.Println(err)
	}
}

func (processor Processor) setPins(ctx context.Context, channelID, msgID string) {
	_, err := processor.Api.AddPins(ctx, channelID, msgID)
	if err != nil {
		log.Println(err)
	}
}

func (processor Processor) setAnnounces(ctx context.Context, data *dto.WSATMessageData) {
	if _, err := processor.Api.CreateChannelAnnounces(
		ctx, data.ChannelID,
		&dto.ChannelAnnouncesToCreate{MessageID: data.ID},
	); err != nil {
		log.Println(err)
	}
}

func (processor Processor) sendReply(ctx context.Context, channelID string, toCreate *dto.MessageToCreate) {
	if _, err := processor.Api.PostMessage(ctx, channelID, toCreate); err != nil {
		log.Println(err)
	}
}
