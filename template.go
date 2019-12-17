package wxmp

import (
	"fmt"
)

func (this *context) Template() ITemplate {
	return &template{
		context: this,
	}
}

type ITemplate interface {
	CreateIndustry(industry ...string) error
}

type template struct {
	context *context
}

func (this *template) error(err interface{}) error {
	return this.context.error(err, "模板消息")
}

func (this *template) CreateIndustry(industry ...string) error {
	data := map[string]interface{}{}
	for i, o := range industry {
		data[fmt.Sprintf("industry_id%d", i+1)] = o
	}
	result := Error{}
	err := this.context.post("/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN", data, &result)
	if err != nil {
		return this.error(err)
	}
	return this.error(result)
}
