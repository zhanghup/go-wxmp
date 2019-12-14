package wxmp

// 文本消息
type MsgTextRes struct {
	MsgHeader
	Content string `xml:"Content" json:"Content"` // 回复的消息内容(换行: 在content中能够换行, 微信客户端支持换行显示)
}

// 图片消息
type MsgImageRes struct {
	MsgHeader
	Image struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image" json:"Image"`
}

// 语音消息
type MsgVoiceRes struct {
	MsgHeader
	Voice struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice" json:"Voice"`
}

// 视频消息
type MsgVideoRes struct {
	MsgHeader
	Video struct {
		MediaId     string `xml:"MediaId"               json:"MediaId"`               // 通过素材管理接口上传多媒体文件得到 MediaId
		Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 视频消息的标题, 可以为空
		Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 视频消息的描述, 可以为空
	} `xml:"Video" json:"Video"`
}

// 音乐消息
type MsgMusicRes struct {
	MsgHeader
	Music struct {
		Title        string `xml:"Title,omitempty"        json:"Title,omitempty"`       // 音乐标题
		Description  string `xml:"Description,omitempty"  json:"Description,omitempty"` // 音乐描述
		MusicURL     string `xml:"MusicUrl"               json:"MusicUrl"`              // 音乐链接
		HQMusicURL   string `xml:"HQMusicUrl"             json:"HQMusicUrl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `xml:"ThumbMediaId"           json:"ThumbMediaId"`          // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `xml:"Music" json:"Music"`
}

// 图文消息里的 Article
type MsgArticleRes struct {
	Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 图文消息描述
	PicURL      string `xml:"PicUrl,omitempty"      json:"PicUrl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         string `xml:"Url,omitempty"         json:"Url,omitempty"`         // 点击图文消息跳转链接
}

// 图文消息
type MsgNewsRes struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MsgHeader
	ArticleCount int             `xml:"ArticleCount"            json:"ArticleCount"`       // 图文消息个数, 限制为10条以内
	Articles     []MsgArticleRes `xml:"Articles>item,omitempty" json:"Articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

// 将消息转发到多客服, 参见多客服模块
type MsgTransferToCustomerService struct {
	MsgHeader
	TransInfo *MsgTransInfo `xml:"TransInfo,omitempty" json:"TransInfo,omitempty"`
}

type MsgTransInfo struct {
	KfAccount string `xml:"KfAccount" json:"KfAccount"`
}