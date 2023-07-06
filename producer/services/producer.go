package services

import (
	"encoding/json"
	"events"
	"reflect"

	"github.com/Shopify/sarama"
)

type EventProducer interface {
	Produce(event events.Event) error
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) eventProducer {
	return eventProducer{
		producer: producer,
	}
}

func (e eventProducer) Produce(event events.Event) error {
	topic := reflect.TypeOf(event).Name()

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}
	_, _, err = e.producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	return nil

}
