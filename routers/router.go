package routers

import (
	"github.com/zhhigh/gomodule/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

}
