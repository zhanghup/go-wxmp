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
	RegisterText(fn func(msg MsgText))
	RegisterImage(fn func(msg MsgImage))
	RegisterVoice(fn func(msg MsgVoice))
	RegisterVideo(fn func(msg MsgVideo))
	RegisterShortVideo(fn func(msg MsgShortVideo))
	RegisterLocation(fn func(msg MsgLocation))
	RegisterLink(fn func(msg MsgLink))

	RegisterEventSubscribe(fn func(msg EventSubscribe))
	RegisterEventUnsubscribe(fn func(msg EventUnsubscribe))
	RegisterEventScan(fn func(msg EventScan))
	RegisterEventLocation(fn func(msg EventLocation))
}

var msg *message

func (this *context) Message() Imessage {
	if msg != nil {
		return msg
	}
	msg = &message{
		context:        this,
		errors:         make([]func(err error), 0),
		msgTexts:       make([]func(msg MsgText), 0),
		msgImages:      make([]func(msg MsgImage), 0),
		msgVoices:      make([]func(msg MsgVoice), 0),
		msgVideos:      make([]func(msg MsgVideo), 0),
		msgShortVideos: make([]func(msg MsgShortVideo), 0),
		msgLocations:   make([]func(msg MsgLocation), 0),
		msgLinks:       make([]func(msg MsgLink), 0),

		eventSubscribe:   make([]func(msg EventSubscribe), 0),
		eventUnsubscribe: make([]func(msg EventUnsubscribe), 0),
		eventScan:        make([]func(msg EventScan), 0),
		eventLocation:    make([]func(msg EventLocation), 0),
	}
	return msg
}

type message struct {
	context *context

	errors         []func(err error)
	msgTexts       []func(msg MsgText)
	msgImages      []func(msg MsgImage)
	msgVoices      []func(msg MsgVoice)
	msgVideos      []func(msg MsgVideo)
	msgShortVideos []func(msg MsgShortVideo)
	msgLocations   []func(msg MsgLocation)
	msgLinks       []func(msg MsgLink)

	eventSubscribe   []func(msg EventSubscribe)
	eventUnsubscribe []func(msg EventUnsubscribe)
	eventScan        []func(msg EventScan)
	eventLocation    []func(msg EventLocation)
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
func (this *message) RegisterText(fn func(msg MsgText)) {
	this.msgTexts = append(this.msgTexts, fn)
}
func (this *message) RegisterImage(fn func(msg MsgImage)) {
	this.msgImages = append(this.msgImages, fn)
}
func (this *message) RegisterVoice(fn func(msg MsgVoice)) {
	this.msgVoices = append(this.msgVoices, fn)
}
func (this *message) RegisterVideo(fn func(msg MsgVideo)) {
	this.msgVideos = append(this.msgVideos, fn)
}
func (this *message) RegisterShortVideo(fn func(msg MsgShortVideo)) {
	this.msgShortVideos = append(this.msgShortVideos, fn)
}
func (this *message) RegisterLocation(fn func(msg MsgLocation)) {
	this.msgLocations = append(this.msgLocations, fn)
}
func (this *message) RegisterLink(fn func(msg MsgLink)) {
	this.msgLinks = append(this.msgLinks, fn)
}

func (this *message) RegisterEventSubscribe(fn func(msg EventSubscribe)) {
	this.eventSubscribe = append(this.eventSubscribe, fn)
}
func (this *message) RegisterEventUnsubscribe(fn func(msg EventUnsubscribe)) {
	this.eventUnsubscribe = append(this.eventUnsubscribe, fn)
}
func (this *message) RegisterEventScan(fn func(msg EventScan)) {
	this.eventScan = append(this.eventScan, fn)
}
func (this *message) RegisterEventLocation(fn func(msg EventLocation)) {
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

			switch MsgType(msgType) {
			case MsgTypeText:
				o := MsgText{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_文本消息")
				}
				for _, f := range this.msgTexts {
					f(o)
				}
			case MsgTypeImage:
				o := MsgImage{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_图片消息")
				}
				for _, f := range this.msgImages {
					f(o)
				}
			case MsgTypeVoice:
				o := MsgVoice{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_语音消息")
				}
				for _, f := range this.msgVoices {
					f(o)
				}
			case MsgTypeVideo:
				o := MsgVideo{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_视频消息")
				}
				for _, f := range this.msgVideos {
					f(o)
				}
			case MsgTypeShortVideo:
				o := MsgShortVideo{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_小视频消息")
				}
				for _, f := range this.msgShortVideos {
					f(o)
				}
			case MsgTypeLocation:
				o := MsgLocation{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_地理位置消息")
				}
				for _, f := range this.msgLocations {
					f(o)
				}
			case MsgTypeLink:
				o := MsgLink{}
				err := xml.Unmarshal(data, o)
				if err != nil {
					this.error(err, "HttpServer_链接消息")
				}
				for _, f := range this.msgLinks {
					f(o)
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
						f(o)
					}
				case MsgEventUnsubscribe: // 取消关注事件
					o := EventUnsubscribe{}
					err := xml.Unmarshal(data, o)
					if err != nil {
						this.error(err, "HttpServer_链接消息")
					}
					for _, f := range this.eventUnsubscribe {
						f(o)
					}
				case MsgEventScan: // 已经关注的用户扫描带参数二维码事件
					o := EventScan{}
					err := xml.Unmarshal(data, o)
					if err != nil {
						this.error(err, "HttpServer_链接消息")
					}
					for _, f := range this.eventScan {
						f(o)
					}
				case MsgEventLocation: // 上报地理位置事件
					o := EventLocation{}
					err := xml.Unmarshal(data, o)
					if err != nil {
						this.error(err, "HttpServer_链接消息")
					}
					for _, f := range this.eventLocation {
						f(o)
					}
				}
			}

		}
	}
}
