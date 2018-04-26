package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type AddController struct {
	beego.Controller
}

func (this *AddController) Get() {
	this.Data["content"] = "value"
	this.Layout = "admin/layout.html"
	this.TplName = "admin/add.tpl"
}

//
//func (this *AddController) Post() {
//	pkgname := this.GetString("admin/layout.html")
//	content := this.GetString("content")
//	pk := models.GetCruPkg(pkgname)
//	if pk.Id {
//
//	}
//
//}
