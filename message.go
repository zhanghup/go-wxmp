package wxmp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type MsgType string
type MsgHeader struct {
	ToUserName   string  `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"   json:"CreateTime"`
	MsgType      MsgType `xml:"MsgType"      json:"MsgType"`
}

const (
	// 普通消息类型
	MsgTypeText       MsgType = "text"       // 文本消息
	MsgTypeImage      MsgType = "image"      // 图片消息
	MsgTypeVoice      MsgType = "voice"      // 语音消息
	MsgTypeVideo      MsgType = "video"      // 视频消息
	MsgTypeShortVideo MsgType = "shortvideo" // 小视频消息
	MsgTypeLocation   MsgType = "location"   // 地理位置消息
	MsgTypeLink       MsgType = "link"       // 链接消息
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
			_, _ = res.Write([]byte(nonce))
		case http.MethodPost:
			data, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Println(string(data))
			}

		}
	}
}

// 文本消息
type MsgText struct {
	MsgHeader
	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

// 图片消息
type MsgImage struct {
	MsgHeader
	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

// 语音消息
type MsgVoice struct {
	MsgHeader
	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 语音消息媒体id, 可以调用多媒体文件下载接口拉取该媒体
	Format  string `xml:"Format"  json:"Format"`  // 语音格式, 如amr, speex等

	// 语音识别结果, UTF8编码,
	// NOTE: 需要开通语音识别功能, 否则该字段为空, 即使开通了语音识别该字段还是有可能为空
	Recognition string `xml:"Recognition,omitempty" json:"Recognition,omitempty"`
}

// 视频消息
type MsgVideo struct {
	MsgHeader
	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据.
}

// 小视频消息
type MsgShortVideo struct {
	MsgHeader
	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据.
}

// 地理位置消息
type MsgLocation struct {
	MsgHeader
	MsgId     int64   `xml:"MsgId"      json:"MsgId"`      // 消息id, 64位整型
	LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
	LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
	Scale     int     `xml:"Scale"      json:"Scale"`      // 地图缩放大小
	Label     string  `xml:"Label"      json:"Label"`      // 地理位置信息
}

// 链接消息
type MsgLink struct {
	MsgHeader
	MsgId       int64  `xml:"MsgId"       json:"MsgId"`       // 消息id, 64位整型
	Title       string `xml:"Title"       json:"Title"`       // 消息标题
	Description string `xml:"Description" json:"Description"` // 消息描述
	URL         string `xml:"Url"         json:"Url"`         // 消息链接
}
