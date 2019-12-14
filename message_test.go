package wxmp

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestMessageText(t *testing.T) {
	data := `
		<xml>
		  <ToUserName><![CDATA[toUser]]></ToUserName>
		  <FromUserName><![CDATA[fromUser]]></FromUserName>
		  <CreateTime>1348831860</CreateTime>
		  <MsgType><![CDATA[text]]></MsgType>
		  <Content><![CDATA[this is a test]]></Content>
		  <MsgId>1234567890123456</MsgId>
		</xml>
	`
	r := MsgText{}
	err := xml.Unmarshal([]byte(data), &r)
	if err != nil {
		panic(err)
	}
}

func TestMessage_HttpServer(t *testing.T) {
	c.Message().RegisterText(func(msg MsgText) interface{} {
		return MsgTextRes{
			MsgHeader: MsgHeader{
				ToUserName:   msg.FromUserName,
				FromUserName: msg.ToUserName,
				CreateTime:   time.Now().Unix(),
				MsgType:      MsgTypeText,
			},
			Content: "hahahah",
		}
	})
	http.HandleFunc("/test", c.Message().HttpServer())
	http.ListenAndServe(":40018", nil)
}

func TestName(t *testing.T) {
	a := map[string]interface{}{
		"a": 1,
		"b": 2,
	}
	d, err := xml.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
