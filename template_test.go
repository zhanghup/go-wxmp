package wxmp

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestTemplate_CreateIndustry(t *testing.T) {
	fmt.Println(c.Template().IndustryCreate(NewIndustry{
		IndustryId1: "1",
		IndustryId2: "2",
	}))
}

func TestTemplate_IndustryGet(t *testing.T) {
	fmt.Println(c.Template().IndustryGet())
}

func TestTemplate_TemplateAdd(t *testing.T) {
	fmt.Println(c.Template().Add("TM00015"))
}

func TestTemplate_List(t *testing.T) {
	a, err := c.Template().List()
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
func TestTemplate_Send(t *testing.T) {
	a, err := c.Template().Send(Template{
		Touser:     "oUVhLxMBWN2uzWL5vv6ZkZbeApy8",
		TemplateId: "kEdg_G8Hz3ZlowZ1HZ4GXeQZwRAlQ9t8k8l4PfeJUe0",
		Url:        tools.Ptr.String("https://www.baidu.com"),
		Data: TemplateData{
			"first": TemplateDataItem{
				Value: "helloworld",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
