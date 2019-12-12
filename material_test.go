package wxmp

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"os"
	"testing"
)

func TestMaterial_NewTempMaterial(t *testing.T) {
	f, err := os.Open(`C:\Users\Administrator\Desktop\资源\图片\下载.jpg`)
	if err != nil {
		panic(err)
	}
	res, err := c.Material().NewTempMaterial(MATERIAL_IMAGE, "test.jpg", f)
	if err != nil {
		panic(err)
	}
	tools.Str().JSONStringPrintln(res)
}

func TestMaterial_GetTempMaterial(t *testing.T) {
	data, err := c.Material().GetTempMaterial("n8R4i9hftwthFd6FPnHuc1KJNuR5SyJFMiYEmZm77WKylRKZVi2ym4CG0u9jVfaI")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

