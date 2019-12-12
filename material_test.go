package go_wxmp

import (
	"github.com/zhanghup/go-tools"
	"os"
	"testing"
)

func TestMaterial_NewTempMaterial(t *testing.T) {
	f, err := os.Open("C:\\Users\\Administrator\\Desktop\\图片s\\img\\qita.jpg")
	if err != nil {
		panic(err)
	}
	res, err := c.Material().NewTempMaterial(MATERIAL_IMAGE, "test.jpg", f)
	if err != nil {
		panic(err)
	}
	tools.Str().JSONStringPrintln(res)
}
