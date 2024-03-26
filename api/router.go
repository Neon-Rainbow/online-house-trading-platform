package api

import (
	"net/http"
	"online-house-trading-platform/api/root"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"online-house-trading-platform/api/auth"
	"online-house-trading-platform/api/houses"
	"online-house-trading-platform/api/user"
)

// SetupRouter 设置web服务器路由
// 该函数返回一个gin.Engine类型的指针,用于设置web服务器的路由
// 该函数具有一下功能:
// 1. 设置静态文件路径,路径为/web/static,用于加载静态文件
// 2. 设置模板文件路径,路径为/web/templates,用于加载模板文件
// 3. 设置/auth路由,用于处理用户的登录,注册和登出请求
// 4. 设置/user路由,用于处理用户的信息请求
// 5. 设置/houses路由,用于处理房屋信息请求
// 6. 设置/路由,用于处理首页请求
// 7. 设置404界面
// 该函数使用了数据库连接,用于处理用户的信息请求
func SetupRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	//加载静态文件
	router.Static("/static", "./web/static")

	//加载模板文件
	router.LoadHTMLGlob("./web/templates/**/*")

	//设置路由,地址为/auth
	auth.SetUpAuthAPI(router, db)

	//设置路由,地址为/user
	user.SetUpUserAPI(router, db)

	//设置路由,地址为/houses
	houses.SetUpHousesAPI(router, db)

	//设置路由,地址为/
	root.SetUpRootAPI(router, db)

	//404界面
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	return router
}
