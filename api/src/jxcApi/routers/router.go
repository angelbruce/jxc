// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"jxcApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),

		beego.NSNamespace("/classify",
			beego.NSInclude(
				&controllers.ClassifyController{},
			),
		),

		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),

		beego.NSNamespace("/shelf",
			beego.NSInclude(
				&controllers.ShelfController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/user_link",
			beego.NSInclude(
				&controllers.UserLinkController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
