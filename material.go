package wxmp

import (
	"bytes"
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
	GetTempMaterialVideo(mediaId string) (string, error)
	GetTempMaterial(mediaId string) (io.Reader, error)
}

type material struct {
	context *context
}

func (this *material) error(err interface{}) error {
	return this.context.error(err, "素材管理")
}

type NewTempMaterialRes struct {
	Error
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

// 新增临时素材
func (this *material) NewTempMaterial(ty MaterialType, name string, f io.Reader) (*NewTempMaterialRes, error) {

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile("media", name)
	if err != nil {
		return nil, this.error(err)
	}

	_, err = io.Copy(formFile, f)
	if err != nil {
		return nil, this.error(err)
	}

	err = writer.Close()
	if err != nil {
		return nil, this.error(err)
	}

	res := new(NewTempMaterialRes)
	err = this.context.postIO("/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type="+string(ty), writer.FormDataContentType(), body, res)
	if err != nil {
		return nil, this.error(err)
	}
	return res, this.error(res.Error)
}

// 获取临时素材（即下载临时的多媒体文件） - 视频文件
func (this *material) GetTempMaterialVideo(mediaId string) (string, error) {
	data := struct {
		Error
		VideoUrl string `json:"video_url"`
	}{}

	err := this.context.get("/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id="+mediaId, nil, &data)
	if err != nil {
		return "", this.error(err)
	}
	return data.VideoUrl, this.error(err)
}

// 获取临时素材（即下载临时的多媒体文件） - 视频文件
func (this *material) GetTempMaterialVioce(mediaId string) (string, error) {
	data := struct {
		Error
		VideoUrl string `json:"video_url"`
	}{}

	err := this.context.get("/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id="+mediaId, nil, &data)
	if err != nil {
		return "", this.error(err)
	}
	return data.VideoUrl, this.error(err)
}

// 获取临时素材（即下载临时的多媒体文件）
func (this *material) GetTempMaterial(mediaId string) (io.Reader, error) {
	res, err := this.context.getIO("/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id="+mediaId, nil)
	if err != nil {
		return nil, this.error(err)
	}
	return res, nil
}

// 新增永久素材
func (this *material) NewMaterial(ty MaterialType, name string, f io.Reader) (*NewTempMaterialRes, error) {

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile("media", name)
	if err != nil {
		return nil, this.error(err)
	}

	_, err = io.Copy(formFile, f)
	if err != nil {
		return nil, this.error(err)
	}

	err = writer.Close()
	if err != nil {
		return nil, this.error(err)
	}

	res := new(NewTempMaterialRes)
	err = this.context.postIO("/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type="+string(ty), writer.FormDataContentType(), body, res)
	if err != nil {
		return nil, this.error(err)
	}
	return res, this.error(res.Error)
}
