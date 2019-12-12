package go_wxmp

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
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
	NewTempMaterial(ty MaterialType, name string, f io.Reader) (*NewTempMaterialRes, error)
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

type NewTempMaterialRes struct {
	Error
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

func (this *material) NewTempMaterial(ty MaterialType, name string, f io.Reader) (*NewTempMaterialRes, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	defer bodyWriter.Close()

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", name)
	if err != nil {
		return nil, this.error(err, "NewTempMaterial")
	}

	_, err = io.Copy(fileWriter, f)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()

	res := new(NewTempMaterialRes)
	err = this.context.postIO("/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type="+string(ty), contentType, bodyBuf, res)
	return res, err
}
