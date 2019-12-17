package wxmp

// 私有
type templateIndustry struct {
	PrimaryIndustry   string `json:"primary_industry"`
	SecondaryIndustry string `json:"secondary_industry"`
	Code              string `json:"code"`
}

type TemplateIndustry struct {
	PrimaryIndustry   TemplateIndustryItem `json:"primary_industry"`
	SecondaryIndustry TemplateIndustryItem `json:"secondary_industry"`
}

type TemplateIndustryItem struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

// 返回的模板列表
type TemplateModel struct {
	TemplateId        string `json:"template_id"`
	Title             string `json:"title"`
	PrimaryIndustry   string `json:"primary_industry"`
	SecondaryIndustry string `json:"secondary_industry"`
	Content           string `json:"content"`
	Example           string `json:"example"`
}

// 发送的模板消息
type TemplateMiniprogram struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}
type TemplateData map[string]TemplateDataItem
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}
type Template struct {
	Touser      string               `json:"touser"` // openid
	TemplateId  string               `json:"template_id"`
	Url         *string              `json:"url"`         // 可空
	Miniprogram *TemplateMiniprogram `json:"miniprogram"` // 可空
	Data        TemplateData         `json:"data"`
}

var TemplateIndustries []templateIndustry

func init() {
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "IT科技", SecondaryIndustry: "互联网/电子商务", Code: "1"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "IT科技", SecondaryIndustry: "IT软件与服务", Code: "2"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "IT科技", SecondaryIndustry: "IT硬件与设备", Code: "3"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "IT科技", SecondaryIndustry: "电子技术", Code: "4"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "IT科技", SecondaryIndustry: "通信与运营商", Code: "5"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "IT科技", SecondaryIndustry: "网络游戏", Code: "6"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "金融业", SecondaryIndustry: "银行", Code: "7"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "金融业", SecondaryIndustry: "基金理财信托", Code: "8"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "金融业", SecondaryIndustry: "保险", Code: "9"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "餐饮", SecondaryIndustry: "餐饮", Code: "10"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "酒店旅游", SecondaryIndustry: "酒店", Code: "11"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "酒店旅游", SecondaryIndustry: "旅游", Code: "12"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "运输与仓储", SecondaryIndustry: "快递", Code: "13"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "运输与仓储", SecondaryIndustry: "物流", Code: "14"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "运输与仓储", SecondaryIndustry: "仓储", Code: "15"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "教育", SecondaryIndustry: "培训", Code: "16"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "教育", SecondaryIndustry: "院校", Code: "17"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "政府与公共事业", SecondaryIndustry: "学术科研", Code: "18"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "政府与公共事业", SecondaryIndustry: "交警", Code: "19"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "政府与公共事业", SecondaryIndustry: "博物馆", Code: "20"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "政府与公共事业", SecondaryIndustry: "公共事业非盈利机构", Code: "21"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "医药护理", SecondaryIndustry: "医药医疗", Code: "22"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "医药护理", SecondaryIndustry: "护理美容", Code: "23"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "医药护理", SecondaryIndustry: "保健与卫生", Code: "24"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "交通工具", SecondaryIndustry: "汽车相关", Code: "25"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "交通工具", SecondaryIndustry: "摩托车相关", Code: "26"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "交通工具", SecondaryIndustry: "火车相关", Code: "27"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "交通工具", SecondaryIndustry: "飞机相关", Code: "28"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "房地产", SecondaryIndustry: "建筑", Code: "29"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "房地产", SecondaryIndustry: "物业", Code: "30"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "消费品", SecondaryIndustry: "消费品", Code: "31"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "商业服务", SecondaryIndustry: "法律", Code: "32"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "商业服务", SecondaryIndustry: "会展", Code: "33"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "商业服务", SecondaryIndustry: "中介服务", Code: "34"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "商业服务", SecondaryIndustry: "认证", Code: "35"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "商业服务", SecondaryIndustry: "审计", Code: "36"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "文体娱乐", SecondaryIndustry: "传媒", Code: "37"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "文体娱乐", SecondaryIndustry: "体育", Code: "38"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "文体娱乐", SecondaryIndustry: "娱乐休闲", Code: "39"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "印刷", SecondaryIndustry: "印刷", Code: "40"})
	TemplateIndustries = append(TemplateIndustries, templateIndustry{PrimaryIndustry: "其它", SecondaryIndustry: "其它", Code: "41"})
}
