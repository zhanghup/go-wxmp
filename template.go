package wxmp

func (this *context) Template() ITemplate {
	return &template{
		context: this,
	}
}

type ITemplate interface {
	IndustryCreate(obj NewIndustry) error
	IndustryGet() (*TemplateIndustry, error)
	Add(templateId string) (string, error)
	List() ([]TemplateModel, error)
	Delete(templateId string) error
	Send(tmp Template) (int64, error)
}

type template struct {
	context *context
}

func (this *template) error(err interface{}) error {
	return this.context.error(err, "模板消息")
}

type NewIndustry struct {
	IndustryId1 string `json:"industry_id1"`
	IndustryId2 string `json:"industry_id2"`
}

// 设置所属行业 - 设置行业可在微信公众平台后台完成，每月可修改行业1次，帐号仅可使用所属行业中相关的模板，为方便第三方开发者，提供通过接口调用的方式来修改账号所属行业
func (this *template) IndustryCreate(obj NewIndustry) error {

	result := Error{}
	err := this.context.post("/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN", obj, &result)
	if err != nil {
		return this.error(err)
	}
	return this.error(result)
}

// 获取帐号设置的行业信息 - 可登录微信公众平台，在公众号后台中查看行业信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号所设置的行业信息
func (this *template) IndustryGet() (*TemplateIndustry, error) {
	result := struct {
		Error
		TemplateIndustry
	}{}
	err := this.context.get("/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN", nil, &result)
	if err != nil {
		return nil, this.error(err)
	}
	return &result.TemplateIndustry, this.error(result.Error)
}

// 获得模板ID - 从行业模板库选择模板到帐号后台，获得模板ID的过程可在微信公众平台后台完成。为方便第三方开发者，提供通过接口调用的方式来获取模板ID
func (this *template) Add(templateId string) (string, error) {
	result := struct {
		Error
		TemplateId string `json:"template_id"`
	}{}

	err := this.context.post("/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN", map[string]interface{}{"template_id_short": templateId}, &result)
	if err != nil {
		return "", err
	}
	return result.TemplateId, this.error(result.Error)
}

// 获取模板列表 - 获取已添加至帐号下所有模板列表，可在微信公众平台后台中查看模板列表信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号下所有模板信息
func (this *template) List() ([]TemplateModel, error) {
	result := struct {
		TemplateList []TemplateModel `json:"template_list"`
		Error
	}{}
	err := this.context.get("/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN", nil, &result)
	if err != nil {
		return nil, err
	}
	return result.TemplateList, this.error(result.Error)
}

// 删除模板 - 删除模板可在微信公众平台后台完成，为方便第三方开发者，提供通过接口调用的方式来删除某帐号下的模板
func (this *template) Delete(templateId string) error {
	result := Error{}
	err := this.context.post("/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN", map[string]interface{}{"template_id": templateId}, &result)
	if err != nil {
		return err
	}
	return this.error(result)
}

// 发送模板消息 - 返回的是消息id
func (this *template) Send(tmp Template) (int64, error) {
	result := struct {
		Msgid int64 `json:"msgid"`
		Error
	}{}
	err := this.context.post("/cgi-bin/message/template/send?access_token=ACCESS_TOKEN", tmp, &result)
	if err != nil {
		return 0, err
	}
	return result.Msgid, this.error(result.Error)
}
