package wxmp

/*
1. 自定义菜单最多包括3个一级菜单，每个一级菜单最多包含5个二级菜单。
2. 一级菜单最多4个汉字，二级菜单最多7个汉字，多出来的部分将会以“...”代替。
3. 创建自定义菜单后，菜单的刷新策略是，在用户进入公众号会话页或公众号profile页时，如果发现上一次拉取菜单的请求在5分钟以前，就会拉取一下菜单，如果菜单有更新，就会刷新客户端的菜单。测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。​
*/
func (this *context) Menu() Imenu {
	return &menu{
		context: this,
	}
}

type Imenu interface {
	Create(btns []Button) error
	Delete() error
	Get() ([]Button, error)
}

type menu struct {
	context *context
}

func (this *menu) error(err interface{}) error {
	return this.context.error(err, "自定义菜单")
}

type Button struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Key      string `json:"key"`
	Url      string `json:"url"`
	Value    string `json:"value"`
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`

	SubButton []Button `json:"sub_button"`
}

func (this *menu) Create(btns []Button) error {
	result := Error{}
	err := this.context.post("/cgi-bin/menu/create?access_token=ACCESS_TOKEN", map[string]interface{}{
		"button": btns,
	}, &result)
	if err != nil {
		return this.error(err)
	}
	return this.error(result)
}

func (this *menu) Delete() error {
	result := Error{}
	err := this.context.get("/cgi-bin/menu/delete?access_token=ACCESS_TOKEN", nil, &result)
	if err != nil {
		return this.error(err)
	}
	return this.error(result)
}

func (this *menu) Get() ([]Button, error) {
	data := struct {
		Error
		IsMenuOPen   int `json:"is_menu_o_pen"`
		SelfmenuInfo struct {
			Button []struct {
				Name      string `json:"name"`
				Type      string `json:"type"`
				Value     string `json:"value"`
				Key       string `json:"key"`
				Url       string `json:"url"`
				Appid     string `json:"appid"`
				Pagepath  string `json:"pagepath"`
				SubButton struct {
					List []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						Key      string `json:"key"`
						Value    string `json:"value"`
						Url      string `json:"url"`
						Appid    string `json:"appid"`
						Pagepath string `json:"pagepath"`
					} `json:"list"`
				} `json:"sub_button"`
			} `json:"button"`
		} `json:"selfmenu_info"`
	}{}
	err := this.context.get("/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN", nil, &data)
	if err != nil {
		return nil, this.error(err)
	}

	err = this.error(data.Error)
	if err != nil {
		return nil, err
	}

	result := make([]Button, 0)
	for _, o := range data.SelfmenuInfo.Button {
		btn := Button{
			Name:      o.Name,
			Type:      o.Type,
			Value:     o.Value,
			Key:       o.Key,
			Url:       o.Url,
			Appid:     o.Appid,
			Pagepath:  o.Pagepath,
			SubButton: make([]Button, 0),
		}
		if o.SubButton.List != nil && len(o.SubButton.List) > 0 {
			for _, oo := range o.SubButton.List {
				btn.SubButton = append(btn.SubButton, Button{
					Name:     oo.Name,
					Type:     oo.Type,
					Value:    oo.Value,
					Key:      oo.Key,
					Url:      oo.Url,
					Appid:    oo.Appid,
					Pagepath: oo.Pagepath,
				})
			}
		}
		result = append(result, btn)
	}
	return result, nil

}
