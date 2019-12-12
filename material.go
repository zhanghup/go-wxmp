package go_wxmp

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type MaterialType string

const (
	MATERIAL_IMAGE MaterialType = "image"
	MATERIAL_VOICE MaterialType = "voice"
	MATERIAL_VIDEO MaterialType = "video"
	MATERIAL_THUMP MaterialType = "thump"
)

func (this *context) Material() Imaterial {
	return &material{
		context: this,
	}
}

type Imaterial interface {
}

type material struct {
	context *context
}

func (this *material) error(err interface{}, fn string) error {
	s := this.context.error(err)
	if len(s) == 0 {
		return nil
	}
	return fmt.Errorf("微信公众号 - 素材管理 - %s - %s", fn, s)
}

func (this *material) NewTempMaterial(ty MaterialType, f io.Reader) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", f.Name())
	if err != nil {
		return this.error(err, "NewTempMaterial")
	}

	multipart.
	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
}
