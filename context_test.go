package wxmp

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"net/url"
	"testing"
)
var c *context


func TestToken(t *testing.T) {
	err := c.token()
	if err != nil{
		panic(err)
	}
	fmt.Print(c.cache.Get("access_token"))
}

func TestUrlEncode(t *testing.T) {
	//url编码
	str := "https://www.baidu.com"
	fmt.Printf("url.QueryEscape:%s", url.QueryEscape(str))
	fmt.Println()
	s, _ := url.QueryUnescape("https://www.baidu.com")
	fmt.Printf("url.QueryUnescape:%s", s)
	fmt.Println()
}


func init() {
	c = &context{
		appid: "wx5fb1dbe2ad657632",
		appsecret:"2265142a415434c47d334ff294798b69",
		cache:tools.NewCache(),
	}
}
