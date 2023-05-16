package controller

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/services"
	"Shopping_System/untils"
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckOrder(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if orderlist, err := services.CheckOrder(claims); err != nil {
		if errors.Is(err, mysql.ErrorOrderExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"您的订单如下": orderlist,
		})
	}
}
