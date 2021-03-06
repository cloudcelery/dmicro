package router

import (
	//"github.com/afex/breaker-go/breaker"
	"dmicro/web/dd/internal/controller"
	"github.com/gin-gonic/gin"
)

// 这里就是 http 请求路径对应的 handler 处理方法了。
func RegisterPassportRouter(g *gin.RouterGroup, c *controller.PassportController) {
	POST(g, "/passport/Login", c.Login, "用户登入")
	POST(g, "/passport/Sms", c.Sms, "获取验证码")
	POST(g, "/passport/SmsLogin", c.SmsLogin, "短信验证码登录")
	POST(g, "/passport/OauthLogin", c.OauthLogin, "第三方账号登录")
	POST(g, "/passport/SetPwd", c.SetPwd, "设置密码")
}
