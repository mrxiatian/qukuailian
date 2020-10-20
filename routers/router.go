package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login.html", &controllers.MainController{},"get:Login;post:LoginCheck")
	beego.Router("/rege.html", &controllers.MainController{},"post:Out")
}
