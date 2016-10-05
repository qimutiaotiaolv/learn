package controllers

import (
	"dragonmonth/mq"
	// "dragonmonth/shared"
	// "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
)

type SaveUploadFileController struct {
	beego.Controller
}

func (this *SaveUploadFileController) Post() {
	_, _, err := this.GetFile("uploadname")
	// defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		fileName := fmt.Sprintf("%stmp-%s.png", "/Users/yangyanxiang/Desktop/石榴FM/png/", bson.NewObjectId().Hex())
		err = this.SaveToFile("uploadname", fileName)
		if err != nil {
			fmt.Println("保存文件失败 ", err)
			this.Ctx.WriteString("保存成功")
			return
		}
		savePath := fmt.Sprintf("%stmp-%s.webp", "/Users/yangyanxiang/Desktop/石榴FM/webp/", bson.NewObjectId().Hex())
		message := mq.NewMessageImageConver2Webp(fileName, savePath, mq.NewDefaultImageConverConnfig())
		mq.Producer_MessageSender_ImageWaitConver2Webp(message)

	}
	this.Ctx.WriteString("保存成功")
}
