package wxmp

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Imessage interface {
	HttpServer() func(res http.ResponseWriter, req *http.Request)
	RegisterError(fn func(err error))
	RegisterText(fn func(msg MsgText) interface{})
	RegisterImage(fn func(msg MsgImage) interface{})
	RegisterVoice(fn func(msg MsgVoice) interface{})
	RegisterVideo(fn func(msg MsgVideo) interface{})
	RegisterShortVideo(fn func(msg MsgShortVideo) interface{})
	RegisterLocation(fn func(msg MsgLocation) interface{})
	RegisterLink(fn func(msg MsgLink) interface{})

	RegisterEventSubscribe(fn func(msg EventSubscribe) interface{})
	RegisterEventUnsubscribe(fn func(msg EventUnsubscribe) interface{})
	RegisterEventScan(fn func(msg EventScan) interface{})
	RegisterEventLocation(fn func(msg EventLocation) interface{})
}

var msg *message

func (this *context) Message() Imessage {
	if msg != nil {
		return msg
	}
	msg = &message{
		context:        this,
		errors:         make([]func(err error) interface{}, 0),
		msgTexts:       make([]func(msg MsgText) interface{}, 0),
		msgImages:      make([]func(msg MsgImage) interface{}, 0),
		msgVoices:      make([]func(msg MsgVoice) interface{}, 0),
		msgVideos:      make([]func(msg MsgVideo) interface{}, 0),
		msgShortVideos: make([]func(msg MsgShortVideo) interface{}, 0),
		msgLocations:   make([]func(msg MsgLocation) interface{}, 0),
		msgLinks:       make([]func(msg MsgLink) interface{}, 0),

		eventSubscribe:   make([]func(msg EventSubscribe) interface{}, 0),
		eventUnsubscribe: make([]func(msg EventUnsubscribe) interface{}, 0),
		eventScan:        make([]func(msg EventScan) interface{}, 0),
		eventLocation:    make([]func(msg EventLocation) interface{}, 0),
	}
	return msg
}

type message struct {
	context *context

	errors         []func(err error)
	msgTexts       []func(msg MsgText) interface{}
	msgImages      []func(msg MsgImage) interface{}
	msgVoices      []func(msg MsgVoice) interface{}
	msgVideos      []func(msg MsgVideo) interface{}
	msgShortVideos []func(msg MsgShortVideo) interface{}
	msgLocations   []func(msg MsgLocation) interface{}
	msgLinks       []func(msg MsgLink) interface{}

	eventSubscribe   []func(msg EventSubscribe) interface{}
	eventUnsubscribe []func(msg EventUnsubscribe) interface{}
	eventScan        []func(msg EventScan) interface{}
	eventLocation    []func(msg EventLocation) interface{}
}

func (this *message) error(err interface{}, fn string) {
	s := this.context.error(err)
	if len(s) == 0 {
		return
	}
	if len(this.errors) > 0 {
		for _, f := range this.errors {
			f(fmt.Errorf("微信公众号 - 消息管理 - %s - %s", fn, s))
		}
	}
}

func (this *message) RegisterError(fn func(err error)) {
	this.errors = append(this.errors, fn)
}
func (this *message) RegisterText(fn func(msg MsgText) interface{}) {
	this.msgTexts = append(this.msgTexts, fn)
}
func (this *message) RegisterImage(fn func(msg MsgImage) interface{}) {
	this.msgImages = append(this.msgImages, fn)
}
func (this *message) RegisterVoice(fn func(msg MsgVoice) interface{}) {
	this.msgVoices = append(this.msgVoices, fn)
}
func (this *message) RegisterVideo(fn func(msg MsgVideo) interface{}) {
	this.msgVideos = append(this.msgVideos, fn)
}
func (this *message) RegisterShortVideo(fn func(msg MsgShortVideo) interface{}) {
	this.msgShortVideos = append(this.msgShortVideos, fn)
}
func (this *message) RegisterLocation(fn func(msg MsgLocation) interface{}) {
	this.msgLocations = append(this.msgLocations, fn)
}
func (this *message) RegisterLink(fn func(msg MsgLink) interface{}) {
	this.msgLinks = append(this.msgLinks, fn)
}

func (this *message) RegisterEventSubscribe(fn func(msg EventSubscribe) interface{}) {
	this.eventSubscribe = append(this.eventSubscribe, fn)
}
func (this *message) RegisterEventUnsubscribe(fn func(msg EventUnsubscribe) interface{}) {
	this.eventUnsubscribe = append(this.eventUnsubscribe, fn)
}
func (this *message) RegisterEventScan(fn func(msg EventScan) interface{}) {
	this.eventScan = append(this.eventScan, fn)
}
func (this *message) RegisterEventLocation(fn func(msg EventLocation) interface{}) {
	this.eventLocation = append(this.eventLocation, fn)
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
			_, _ = res.Write([]byte(nonce))
		case http.MethodPost:
			data, err := ioutil.ReadAll(req.Body)
			if err != nil {
				this.error(err, "HttpServer_0")
				return
			}
			hd := map[string]string{}
			err = xml.Unmarshal(data, hd)
			if err != nil {
				this.error(err, "HttpServer_1")
				return
			}

			msgType := hd["MsgType"]
			msgEvent := hd["Event"]

			var response interface{}
			switch MsgType(msgType) {
			case MsgTypeText:
				o := MsgText{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_文本消息")
				}
				for _, f := range this.msgTexts {
					response = f(o)
				}
			case MsgTypeImage:
				o := MsgImage{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_图片消息")
				}
				for _, f := range this.msgImages {
					response = f(o)
				}
			case MsgTypeVoice:
				o := MsgVoice{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_语音消息")
				}
				for _, f := range this.msgVoices {
					response = f(o)
				}
			case MsgTypeVideo:
				o := MsgVideo{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_视频消息")
				}
				for _, f := range this.msgVideos {
					response = f(o)
				}
			case MsgTypeShortVideo:
				o := MsgShortVideo{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_小视频消息")
				}
				for _, f := range this.msgShortVideos {
					response = f(o)
				}
			case MsgTypeLocation:
				o := MsgLocation{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_地理位置消息")
				}
				for _, f := range this.msgLocations {
					response = f(o)
				}
			case MsgTypeLink:
				o := MsgLink{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_链接消息")
				}
				for _, f := range this.msgLinks {
					response = f(o)
				}
			case MsgTypeEvent: // 事件消息
				switch MsgEvent(msgEvent) {
				case MsgEventSubscribe: // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
					o := EventSubscribe{}
					err := xml.Unmarshal(data, o)
					if err != nil {
						this.error(err, "HttpServer_链接消息")
					}
					for _, f := range this.eventSubscribe {
						response = f(o)
					}
				case MsgEventUnsubscribe: // 取消关注事件
					o := EventUnsubscribe{}
					err := xml.Unmarshal(data, o)
					if err != nil {
						this.error(err, "HttpServer_链接消息")
					}
					for _, f := range this.eventUnsubscribe {
						response = f(o)
					}
				case MsgEventScan: // 已经关注的用户扫描带参数二维码事件
					o := EventScan{}
					err := xml.Unmarshal(data, o)
					if err != nil {
						this.error(err, "HttpServer_链接消息")
					}
					for _, f := range this.eventScan {
						response = f(o)
					}
				case MsgEventLocation: // 上报地理位置事件
					o := EventLocation{}
					err := xml.Unmarshal(data, o)
					if err != nil {
						this.error(err, "HttpServer_链接消息")
					}
					for _, f := range this.eventLocation {
						response = f(o)
					}
				}
			}

			// 消息回复
			if response != nil {

			} else {
				_, _ = res.Write([]byte(""))
			}
		}
	}
}
