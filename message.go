package wxmp

import (
	ctx "context"
	"encoding/xml"
	"fmt"
	"github.com/zhanghup/go-tools"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

type Imessage interface {
	HttpServer() func(res http.ResponseWriter, req *http.Request)
	ContextTimeout(t int64) Imessage
	RegisterError(fn func(err error))

	RegisterText(fn func(ctx ctx.Context, msg MsgText) (ctx.Context, interface{}))
	RegisterImage(fn func(ctx ctx.Context, msg MsgImage) (ctx.Context, interface{}))
	RegisterVoice(fn func(ctx ctx.Context, msg MsgVoice) (ctx.Context, interface{}))
	RegisterVideo(fn func(ctx ctx.Context, msg MsgVideo) (ctx.Context, interface{}))
	RegisterShortVideo(fn func(ctx ctx.Context, msg MsgShortVideo) (ctx.Context, interface{}))
	RegisterLocation(fn func(ctx ctx.Context, msg MsgLocation) (ctx.Context, interface{}))
	RegisterLink(fn func(ctx ctx.Context, msg MsgLink) (ctx.Context, interface{}))

	RegisterEventSubscribe(fn func(ctx ctx.Context, msg EventSubscribe) (ctx.Context, interface{}))
	RegisterEventUnsubscribe(fn func(ctx ctx.Context, msg EventUnsubscribe) (ctx.Context, interface{}))
	RegisterEventScan(fn func(ctx ctx.Context, msg EventScan) (ctx.Context, interface{}))
	RegisterEventLocation(fn func(ctx ctx.Context, msg EventLocation) (ctx.Context, interface{}))
	RegisterEventTemplateFinish(fn func(ctx ctx.Context, msg EventTemplateFinish) (ctx.Context, interface{}))
}

var msg *message

func (this *context) Message() Imessage {
	if msg != nil {
		return msg
	}
	msg = &message{
		context:        this,
		ctx:            tools.CacheCreate(true),
		cancels:        tools.CacheCreate(true),
		ctxTimeout:     600,
		errors:         make([]func(err error), 0),
		msgTexts:       make([]func(ctx ctx.Context, msg MsgText) (ctx.Context, interface{}), 0),
		msgImages:      make([]func(ctx ctx.Context, msg MsgImage) (ctx.Context, interface{}), 0),
		msgVoices:      make([]func(ctx ctx.Context, msg MsgVoice) (ctx.Context, interface{}), 0),
		msgVideos:      make([]func(ctx ctx.Context, msg MsgVideo) (ctx.Context, interface{}), 0),
		msgShortVideos: make([]func(ctx ctx.Context, msg MsgShortVideo) (ctx.Context, interface{}), 0),
		msgLocations:   make([]func(ctx ctx.Context, msg MsgLocation) (ctx.Context, interface{}), 0),
		msgLinks:       make([]func(ctx ctx.Context, msg MsgLink) (ctx.Context, interface{}), 0),

		eventSubscribe:   make([]func(ctx ctx.Context, msg EventSubscribe) (ctx.Context, interface{}), 0),
		eventUnsubscribe: make([]func(ctx ctx.Context, msg EventUnsubscribe) (ctx.Context, interface{}), 0),
		eventScan:        make([]func(ctx ctx.Context, msg EventScan) (ctx.Context, interface{}), 0),
		eventLocation:    make([]func(ctx ctx.Context, msg EventLocation) (ctx.Context, interface{}), 0),
	}
	return msg
}

type message struct {
	context    *context
	ctx        tools.ICache
	cancels    tools.ICache
	ctxTimeout int64

	errors         []func(err error)
	msgTexts       []func(ctx ctx.Context, msg MsgText) (ctx.Context, interface{})
	msgImages      []func(ctx ctx.Context, msg MsgImage) (ctx.Context, interface{})
	msgVoices      []func(ctx ctx.Context, msg MsgVoice) (ctx.Context, interface{})
	msgVideos      []func(ctx ctx.Context, msg MsgVideo) (ctx.Context, interface{})
	msgShortVideos []func(ctx ctx.Context, msg MsgShortVideo) (ctx.Context, interface{})
	msgLocations   []func(ctx ctx.Context, msg MsgLocation) (ctx.Context, interface{})
	msgLinks       []func(ctx ctx.Context, msg MsgLink) (ctx.Context, interface{})

	eventSubscribe      []func(ctx ctx.Context, msg EventSubscribe) (ctx.Context, interface{})
	eventUnsubscribe    []func(ctx ctx.Context, msg EventUnsubscribe) (ctx.Context, interface{})
	eventScan           []func(ctx ctx.Context, msg EventScan) (ctx.Context, interface{})
	eventLocation       []func(ctx ctx.Context, msg EventLocation) (ctx.Context, interface{})
	eventTemplateFinish []func(ctx ctx.Context, msg EventTemplateFinish) (ctx.Context, interface{})
}

func (this *message) error(err interface{}) {
	s := this.context.error(err, "消息管理")
	if s == nil {
		return
	}

	if len(this.errors) > 0 {
		for _, f := range this.errors {
			f(s)
		}
	}
}

func (this *message) RegisterError(fn func(err error)) {
	this.errors = append(this.errors, fn)
}
func (this *message) RegisterText(fn func(ctx ctx.Context, msg MsgText) (ctx.Context, interface{})) {
	this.msgTexts = append(this.msgTexts, fn)
}
func (this *message) RegisterImage(fn func(ctx ctx.Context, msg MsgImage) (ctx.Context, interface{})) {
	this.msgImages = append(this.msgImages, fn)
}
func (this *message) RegisterVoice(fn func(ctx ctx.Context, msg MsgVoice) (ctx.Context, interface{})) {
	this.msgVoices = append(this.msgVoices, fn)
}
func (this *message) RegisterVideo(fn func(ctx ctx.Context, msg MsgVideo) (ctx.Context, interface{})) {
	this.msgVideos = append(this.msgVideos, fn)
}
func (this *message) RegisterShortVideo(fn func(ctx ctx.Context, msg MsgShortVideo) (ctx.Context, interface{})) {
	this.msgShortVideos = append(this.msgShortVideos, fn)
}
func (this *message) RegisterLocation(fn func(ctx ctx.Context, msg MsgLocation) (ctx.Context, interface{})) {
	this.msgLocations = append(this.msgLocations, fn)
}
func (this *message) RegisterLink(fn func(ctx ctx.Context, msg MsgLink) (ctx.Context, interface{})) {
	this.msgLinks = append(this.msgLinks, fn)
}

func (this *message) RegisterEventSubscribe(fn func(ctx ctx.Context, msg EventSubscribe) (ctx.Context, interface{})) {
	this.eventSubscribe = append(this.eventSubscribe, fn)
}
func (this *message) RegisterEventUnsubscribe(fn func(ctx ctx.Context, msg EventUnsubscribe) (ctx.Context, interface{})) {
	this.eventUnsubscribe = append(this.eventUnsubscribe, fn)
}
func (this *message) RegisterEventScan(fn func(ctx ctx.Context, msg EventScan) (ctx.Context, interface{})) {
	this.eventScan = append(this.eventScan, fn)
}
func (this *message) RegisterEventLocation(fn func(ctx ctx.Context, msg EventLocation) (ctx.Context, interface{})) {
	this.eventLocation = append(this.eventLocation, fn)
}

func (this *message) RegisterEventTemplateFinish(fn func(ctx ctx.Context, msg EventTemplateFinish) (ctx.Context, interface{})) {
	this.eventTemplateFinish = append(this.eventTemplateFinish, fn)
}

func (this *message) xml(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("<xml>%s</xml>", this.reflectInterface(reflect.TypeOf(v), reflect.ValueOf(v)))
}

func (this *message) xmlItem(wrap string, ty reflect.Type, vl reflect.Value) string {
	switch ty.Kind() {
	case reflect.Map:
		return fmt.Sprintf("<%s>%s</%s>", wrap, this.reflectInterface(ty, vl), wrap)
	case reflect.Struct:
		return fmt.Sprintf("<%s>%s</%s>", wrap, this.reflectInterface(ty, vl), wrap)
	case reflect.String:
		return fmt.Sprintf("<%s><![CDATA[%s]]></%s>", wrap, vl.String(), wrap)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return fmt.Sprintf("<%s>%d</%s>", wrap, vl.Uint(), wrap)
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return fmt.Sprintf("<%s>%d</%s>", wrap, vl.Int(), wrap)
	}
	return ""
}

func (this *message) reflectInterface(t reflect.Type, v reflect.Value) string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	result := ""
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			vl := v.Field(i)
			ty := t.Field(i).Type
			field := t.Field(i)

			wrap := field.Tag.Get("xml")
			if len(wrap) == 0 {
				wrap = field.Tag.Get("json")
			}
			if len(wrap) == 0 {
				wrap = field.Name
			}

			if ty.Kind() == reflect.Ptr {
				if vl.Pointer() == 0 {
					result += fmt.Sprintf("<%s></%s>", wrap, wrap)
					continue
				}
				vl = vl.Elem()
				ty = ty.Elem()
			}
			if ty.Kind() == reflect.Struct && field.Anonymous {
				result += this.reflectInterface(ty, vl)
			} else {
				result += this.xmlItem(wrap, ty, vl)
			}
		}
	} else if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			if k.Kind() != reflect.String {
				continue
			}
			vl := v.MapIndex(k)
			ty := vl.Type()
			result += this.xmlItem(k.String(), ty, vl)
		}
	}
	return result
}

func (this *message) HttpServer() func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		values := req.URL.Query()
		signature := values.Get("signature")
		timestamp := values.Get("timestamp")
		nonce := values.Get("nonce")
		echostr := values.Get("echostr")
		if len(signature) == 0 || len(timestamp) == 0 || len(nonce) == 0 {
			return
		}

		if !this.context.sign(timestamp, nonce, signature) {
			return
		}

		switch req.Method {
		case http.MethodGet:
			_, _ = res.Write([]byte(echostr))
		case http.MethodPost:
			data, err := ioutil.ReadAll(req.Body)
			if err != nil {
				this.error(err)
				return
			}
			hd := struct {
				MsgHeader
				MsgType string `xml:"MsgType"`
				Event   string `xml:"Event"`
			}{}
			err = xml.Unmarshal(data, &hd)
			if err != nil {
				this.error(err)
				return
			}

			msgCtx := this.MsgContext(hd.FromUserName)

			var response interface{}
			switch MsgType(hd.MsgType) {
			case MsgTypeText:
				o := MsgText{}
				err := xml.Unmarshal(data, &o)
				if err != nil {
					this.error(err)
				}
				for _, f := range this.msgTexts {
					msgCtx, response = f(msgCtx, o)
				}
			case MsgTypeImage:
				o := MsgImage{}
				err := xml.Unmarshal(data, &o)
				if err != nil {
					this.error(err)
				}
				for _, f := range this.msgImages {
					msgCtx, response = f(msgCtx, o)
				}
			case MsgTypeVoice:
				o := MsgVoice{}
				err := xml.Unmarshal(data, &o)
				if err != nil {
					this.error(err)
				}
				for _, f := range this.msgVoices {
					msgCtx, response = f(msgCtx, o)
				}
			case MsgTypeVideo:
				o := MsgVideo{}
				err := xml.Unmarshal(data, &o)
				if err != nil {
					this.error(err)
				}
				for _, f := range this.msgVideos {
					msgCtx, response = f(msgCtx, o)
				}
			case MsgTypeShortVideo:
				o := MsgShortVideo{}
				err := xml.Unmarshal(data, &o)
				if err != nil {
					this.error(err)
				}
				for _, f := range this.msgShortVideos {
					msgCtx, response = f(msgCtx, o)
				}
			case MsgTypeLocation:
				o := MsgLocation{}
				err := xml.Unmarshal(data, &o)
				if err != nil {
					this.error(err)
				}
				for _, f := range this.msgLocations {
					msgCtx, response = f(msgCtx, o)
				}
			case MsgTypeLink:
				o := MsgLink{}
				err := xml.Unmarshal(data, &o)
				if err != nil {
					this.error(err)
				}
				for _, f := range this.msgLinks {
					msgCtx, response = f(msgCtx, o)
				}
			case MsgTypeEvent: // 事件消息
				switch MsgEvent(hd.Event) {
				case MsgEventSubscribe: // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
					o := EventSubscribe{}
					err := xml.Unmarshal(data, &o)
					if err != nil {
						this.error(err)
					}
					for _, f := range this.eventSubscribe {
						msgCtx, response = f(msgCtx, o)
					}
				case MsgEventUnsubscribe: // 取消关注事件
					o := EventUnsubscribe{}
					err := xml.Unmarshal(data, &o)
					if err != nil {
						this.error(err)
					}
					for _, f := range this.eventUnsubscribe {
						msgCtx, response = f(msgCtx, o)
					}
				case MsgEventScan: // 已经关注的用户扫描带参数二维码事件
					o := EventScan{}
					err := xml.Unmarshal(data, &o)
					if err != nil {
						this.error(err)
					}
					for _, f := range this.eventScan {
						msgCtx, response = f(msgCtx, o)
					}
				case MsgEventLocation: // 上报地理位置事件
					o := EventLocation{}
					err := xml.Unmarshal(data, &o)
					if err != nil {
						this.error(err)
					}
					for _, f := range this.eventLocation {
						msgCtx, response = f(msgCtx, o)
					}
				case MsgEventTemplateFinish:
					o := EventTemplateFinish{}
					err := xml.Unmarshal(data, &o)
					if err != nil {
						this.error(err)
					}
					for _, f := range this.eventTemplateFinish {
						msgCtx, response = f(msgCtx, o)
					}
				}
			}

			if msgCtx != nil {
				this.MsgContextUpdate(hd.FromUserName, msgCtx)
			} else {
				this.MsgContextClose(hd.FromUserName)
			}

			// 消息回复
			rr := this.xml(response)
			_, _ = res.Write([]byte(rr))
		}
	}
}

func (this *message) ContextTimeout(t int64) Imessage {
	this.ctxTimeout = t
	return this
}

func (this *message) MsgContext(openid string) ctx.Context {
	var c ctx.Context
	var cancel func()
	if this.ctx.Get(openid) == nil {
		c, cancel = ctx.WithCancel(ctx.Background())
	} else {
		c = this.ctx.Get(openid).(ctx.Context)
		cancel = this.cancels.Get(openid).(func())

	}
	this.cancels.Set(openid, cancel, this.ctxTimeout+time.Now().Unix())
	this.ctx.Set(openid, c, this.ctxTimeout+time.Now().Unix())
	return c
}

func (this *message) MsgContextUpdate(openid string, c ctx.Context) {
	this.ctx.Set(openid, c, this.ctxTimeout+time.Now().Unix())
}
func (this *message) MsgContextClose(openid string) {
	this.ctx.Delete(openid)
	fn := this.ctx.Get(openid)
	if fn != nil {
		fn.(func())()
	}
}
