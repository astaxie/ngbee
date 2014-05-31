package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/ngbee/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:all", &controllers.MainController{})
}
