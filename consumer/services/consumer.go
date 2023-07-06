package services

import "github.com/Shopify/sarama"

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return consumerHandler{eventHandler: eventHandler}
}

func (h consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.eventHandler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}
	return nil
}
