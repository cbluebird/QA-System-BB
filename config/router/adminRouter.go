package router

import (
	"QA-system/controller/adminController"
	midwares "QA-system/midware"
	"github.com/gin-gonic/gin"
)

func adminRouterInit(r *gin.RouterGroup) {
	admin := r.Group("/admin", midwares.CheckLogin)
	{
		admin.POST("/add", adminController.CreateList)
		admin.GET("/single/get", adminController.GetAdminListByID)
		admin.GET("/detail/get", adminController.GetAns)
		admin.POST("/draft/status", adminController.UpdateStatus)
		admin.GET("/list/get", adminController.GetList)
		admin.POST("/draft/update", adminController.UpdateList)
		admin.POST("/delete", adminController.Delete)
	}
}
