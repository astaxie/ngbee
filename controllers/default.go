package controllers

import (
	"path"

	"github.com/astaxie/beego"
	"github.com/astaxie/ngbee/bootstrap"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var (
		defaultTmp    string
		defaultLayout string
		findRouter    bool
	)
	all := this.GetString(":all")
	this.Layout = path.Join(beego.AppConfig.String("theme"), bootstrap.URLRule.Layout)
	for _, r := range bootstrap.URLRule.Routers {
		if r.URL == all {
			this.TplNames = path.Join(beego.AppConfig.String("theme"), r.Template)
			if r.Layout != "" {
				this.Layout = path.Join(beego.AppConfig.String("theme"), r.Layout)
			}
			findRouter = true
		}
		if all == "/" {
			defaultTmp = r.Template
			defaultLayout = r.Layout
		}
	}

	if !findRouter {
		this.TplNames = path.Join(beego.AppConfig.String("theme"), defaultTmp)
		if defaultLayout != "" {
			this.Layout = path.Join(beego.AppConfig.String("theme"), defaultLayout)
		}
	}
}
