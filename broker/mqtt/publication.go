package mqtt

import "github.com/bis0si0n/kratos-transport/broker"

type publication struct {
	topic string
	msg   *broker.Message
	err   error
}

func (p *publication) Ack() error {
	return nil
}

func (p *publication) Error() error {
	return p.err
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	return p.msg
}

func (p *publication) RawMessage() interface{} {
	return p.msg
}
