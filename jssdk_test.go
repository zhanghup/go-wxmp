package wxmp

import (
	"fmt"
	"testing"
)

func TestJssdk_AuthUrl(t *testing.T) {
	fmt.Println(c.JSSDK().AuthUrl("https://www.baidu.com","snsapi_userinfos","test"))
}

func TestJssdk_Auth(t *testing.T) {
	code := "021vML301bFYNX1d2D301ZQI301vML3f"
	fmt.Println(c.JSSDK().Auth(code))
}

func TestJssdk_AuthUserInfo(t *testing.T) {
	code := "001K6VsN0Hftxa2RnLvN0NLIsN0K6Vsx"
	fmt.Println(c.JSSDK().AuthUserInfo(code))
}