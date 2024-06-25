package middleware

// DBMiddleware 创建一个中间件，将 *gorm.DB 实例添加到上下文中
//func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Set("db", db)
//		c.Next()
//	}
//}
