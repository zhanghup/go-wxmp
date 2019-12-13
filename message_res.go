package wxmp

// 文本消息
type MsgTextRes struct {
	MsgHeader
	Content string `xml:"Content" json:"Content"` // 回复的消息内容(换行: 在content中能够换行, 微信客户端支持换行显示)
}

// 图片消息
type Image struct {
	MsgHeader
	Image struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image" json:"Image"`
}

// 语音消息
type Voice struct {
	MsgHeader
	Voice struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice" json:"Voice"`
}
