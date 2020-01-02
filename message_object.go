package wxmp

import (
	"fmt"
	"strings"
	"time"
)

type MsgType string
type MsgEvent string

type MsgHeader struct {
	ToUserName   string  `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"   json:"CreateTime"`
	MsgType      MsgType `xml:"MsgType"      json:"MsgType"`
}

func (this MsgHeader) Response() MsgHeader {
	return MsgHeader{
		ToUserName:   this.FromUserName,
		FromUserName: this.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      this.MsgType,
	}
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
	MsgTypeEvent      MsgType = "event"      // 事件消息

	MsgEventSubscribe      MsgEvent = "subscribe"             // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
	MsgEventUnsubscribe    MsgEvent = "unsubscribe"           // 取消关注事件
	MsgEventScan           MsgEvent = "SCAN"                  // 已经关注的用户扫描带参数二维码事件
	MsgEventLocation       MsgEvent = "LOCATION"              // 上报地理位置事件
	MsgEventTemplateFinish MsgEvent = "TEMPLATESENDJOBFINISH" // 模板消息发送结束
)

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

// 关注事件
type EventSubscribe struct {
	MsgHeader
	EventType string `xml:"Event" json:"Event"` // subscribe

	// 下面两个字段只有在扫描带参数二维码进行关注时才有值, 否则为空值!
	EventKey string `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 格式为: qrscene_二维码的参数值
	Ticket   string `xml:"Ticket,omitempty"   json:"Ticket,omitempty"`   // 二维码的ticket, 可用来换取二维码图片
}

// 关注事件 - 获取二维码参数
func (event *EventSubscribe) Scan() (scene string, err error) {
	const prefix = "qrscene_"
	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %s 为前缀: %s", prefix, event.EventKey)
		return
	}
	scene = event.EventKey[len(prefix):]
	return
}

// 取消关注事件
type EventUnsubscribe struct {
	MsgHeader
	EventType string `xml:"Event"              json:"Event"`              // unsubscribe
	EventKey  string `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 空值
}

// 用户已关注时, 扫描带参数二维码的事件
type EventScan struct {
	MsgHeader
	EventType string `xml:"Event"    json:"Event"`    // SCAN
	EventKey  string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 二维码的参数值(scene_id, scene_str)
	Ticket    string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket, 可用来换取二维码图片
}

// 上报地理位置事件
type EventLocation struct {
	MsgHeader
	EventType string  `xml:"Event"     json:"Event"`     // LOCATION
	Latitude  float64 `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64 `xml:"Precision" json:"Precision"` // 地理位置精度(整数? 但是微信推送过来是浮点数形式)
}

// 模板消息发送反馈
type EventTemplateFinish struct {
	MsgHeader
	Msgid  string `xml:"MsgId" json:"MsgId"`   // 消息id
	Status string `xml:"Status" json:"Status"` // 发送状态为发送失败（[failed: system failed 非用户拒绝] [failed:user block 发送状态为用户拒绝接收]）
}
