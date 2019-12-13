package wxmp

import (
	"fmt"
	"net/http"
)

func (this *context) Message() Imessage {
	return &message{
		context: this,
	}
}

type Imessage interface {
	HttpServer() func(res http.ResponseWriter, req *http.Request)
}

type message struct {
	context *context

	getActions  []func()
	postActions []func()
}

func (this *message) error(err interface{}, fn string) error {
	s := this.context.error(err)
	if len(s) == 0 {
		return nil
	}
	return fmt.Errorf("微信公众号 - 消息管理 - %s - %s", fn, s)
}

func (this *message) HttpServer() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		values := req.URL.Query()
		signature := values.Get("signature")
		timestamp := values.Get("timestamp")
		nonce := values.Get("nonce")
		echostr := values.Get("echostr")
		if len(signature) == 0 || len(timestamp) == 0 || len(nonce) == 0 || len(echostr) == 0 {
			return
		}

		if !this.context.sign(timestamp, nonce, echostr, signature) {
			return
		}

		switch req.Method {
		case http.MethodGet:

		case http.MethodPost:
		}
	}
}

func (this *message) HttpPost() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

	}
}
