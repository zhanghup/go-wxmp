package wxmp

import "fmt"

func (this *context) Message() Imessage {
	return &message{
		context: this,
	}
}

type Imessage interface {
}

type message struct {
	context *context
}

func (this *message) error(err interface{}, fn string) error {
	s := this.context.error(err)
	if len(s) == 0 {
		return nil
	}
	return fmt.Errorf("微信公众号 - 消息管理 - %s - %s", fn, s)
}
