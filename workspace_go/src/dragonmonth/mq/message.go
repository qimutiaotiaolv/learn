package mq

/***
*转换webp所需要的参数
*
 **/
type ImageConverConfig struct {
	WidthPix  uint
	HeightPix uint
	Quality   float32
	Lossless  bool
}

func NewDefaultImageConverConnfig() *ImageConverConfig {
	config := &ImageConverConfig{
		WidthPix:  1242,
		HeightPix: 2208,
		Quality:   70,
		Lossless:  false,
	}
	return config
}

/***
*当原始图片(png)存储到临时文件后，会向topic_image_waitConver2Webp发送此message
*OriginalImagePath: 图片临时目录
*Config: 转换图片时的参数
 **/
type MessageImageWaitConver2Webp struct {
	OriginalImagePath string
	SaveWebpTempPath  string
	Config            *ImageConverConfig
}

func NewMessageImageConver2Webp(path string, savePath string, config *ImageConverConfig) *MessageImageWaitConver2Webp {
	message := &MessageImageWaitConver2Webp{
		OriginalImagePath: path,
		SaveWebpTempPath:  savePath,
		Config:            config,
	}
	if message.Config == nil {
		message.Config = NewDefaultImageConverConnfig()
	}
	return message
}

/*
*  数据库保存
*
 */
