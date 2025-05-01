package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func RouteLocalFrontend(e *gin.Engine) {
	dir := conf.LocalFrontend.Dir
	e.NoRoute(func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		fullPath := filepath.Join(dir, requestPath)

		// 尝试返回实际文件
		if _, err := filepath.Glob(fullPath); err == nil {
			c.File(fullPath)
		} else {
			// 如果文件不存在，返回 index.html（适用于 SPA 前端）
			c.File(filepath.Join(dir, "index.html"))
		}
	})

	logrus.Info("register local frontend in / dir: " + dir)
}
