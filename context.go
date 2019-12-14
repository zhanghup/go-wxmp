package wxmp

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhanghup/go-tools"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type IContext interface {
	JSSDK() IJssdk
	Material() Imaterial
	Menu() Imenu
	Message() Imessage
}

type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type context struct {
	appid     string
	appsecret string
	stoken    string // 这个token是配置在服务器上的token，用于数据校验

	cache tools.IMap
}

func NewContext(appid, appsecret, token string) IContext {
	c := new(context)
	c.cache = tools.NewCache()
	c.appid = appid
	c.appsecret = appsecret
	c.stoken = token
	return c
}

func (this *context) error(err interface{}) string {
	if err == nil {
		return ""
	}
	var s = ""
	switch err.(type) {
	case string:
		s = err.(string)
	case error:
		s = err.(error).Error()
	case Error:
		e := err.(Error)
		if e.Errcode == 0 {
			return ""
		}
		s = fmt.Sprintf("%d: %s", e.Errcode, e.Errmsg)
	}
	return s
}

func (this *context) token() error {
	if this.cache.Contain("access_token") {
		return nil
	}

	res, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", this.appid, this.appsecret))
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	tok := struct {
		Error
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	err = json.Unmarshal(data, &tok)
	if err != nil {
		return err
	}
	if tok.Errcode != 0 {
		return errors.New(fmt.Sprintf("%d: %s", tok.Errcode, tok.Errmsg))
	}
	this.cache.Set2("access_token", tok.AccessToken, time.Now().Unix()+int64(tok.ExpiresIn))
	return nil
}

// 可做容灾
func (this *context) url() string {
	return "https://api.weixin.qq.com"
}

func (this *context) get(url string, param map[string]interface{}, result interface{}) error {
	err := this.token()
	if err != nil {
		return err
	}
	token := this.cache.Get("access_token").(string)
	url = strings.Replace(this.url()+url, "ACCESS_TOKEN", token, 1)

	return tools.Http().GetI(url, param, result)
}

func (this *context) getIO(url string, param map[string]interface{}) (io.Reader, error) {
	err := this.token()
	if err != nil {
		return nil, err
	}
	token := this.cache.Get("access_token").(string)
	url = strings.Replace(this.url()+url, "ACCESS_TOKEN", token, 1)

	res, err := tools.Http().GetF(url, param)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (this *context) post(url string, param, result interface{}) error {
	err := this.token()
	if err != nil {
		return err
	}
	token := this.cache.Get("access_token").(string)
	url = strings.Replace(url, "ACCESS_TOKEN", token, 1)

	return tools.Http().PostI(this.url()+url, param, result)
}

func (this *context) postIO(url string, contentType string, param io.Reader, result interface{}) error {
	err := this.token()
	if err != nil {
		return err
	}
	token := this.cache.Get("access_token").(string)
	url = strings.Replace(url, "ACCESS_TOKEN", token, 1)

	res, err := http.Post(this.url()+url, contentType, param)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func (this *context) sign(timestamp, nonce, sign string) bool {
	strs := sort.StringSlice{this.stoken, timestamp, nonce}
	strs.Sort()

	h := sha1.New()

	bufw := bufio.NewWriterSize(h, 128) // sha1.BlockSize 的整数倍
	bufw.WriteString(strs[0])
	bufw.WriteString(strs[1])
	bufw.WriteString(strs[2])
	bufw.Flush()

	hashsum := h.Sum(nil)
	s := hex.EncodeToString(hashsum)
	return s == sign

}
