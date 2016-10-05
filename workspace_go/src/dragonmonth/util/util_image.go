package util

import (
	"bytes"
	// "dragonmonth/mq"
	"github.com/astaxie/beego"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
)

type ImageConverConfig struct {
	WidthPix  uint
	HeightPix uint
	Quality   float32
	Lossless  bool
	FromPath  string
	ToPath    string
}

func NewImageConverConfig(width, height uint, quality float32, lossless bool, fromPath, toPath string) *ImageConverConfig {
	c := &ImageConverConfig{
		WidthPix:  width,
		HeightPix: height,
		Quality:   quality,
		Lossless:  lossless,
		FromPath:  fromPath,
		ToPath:    toPath,
	}
	return c
}

func Encode2Webp(config *ImageConverConfig, saveSuccessCallback func(*ImageConverConfig)) {
	file, err := os.Open(config.FromPath)
	if err != nil {
		beego.BeeLogger.Error("Encode2Webp::图片打开失败:%s", err.Error())
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		beego.BeeLogger.Error("Encode2Webp::图片编码错误:%s", err.Error())
	}
	file.Close()
	scalImage := resize.Resize(config.WidthPix, config.HeightPix, img, resize.Lanczos3)
	var buf bytes.Buffer
	if err := webp.Encode(&buf, scalImage, &webp.Options{Lossless: config.Lossless, Quality: config.Quality}); err != nil {
		beego.BeeLogger.Error("Encode2Webp::转化webp错误:%s", err.Error())
		return
	}
	if err = ioutil.WriteFile(config.ToPath, buf.Bytes(), 0666); err != nil {
		beego.BeeLogger.Error("Encode2Webp::webp保存失败%s", err.Error())
		return
	}
	saveSuccessCallback(config)
}
