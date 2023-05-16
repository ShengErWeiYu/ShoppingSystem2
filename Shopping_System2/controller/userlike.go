package controller

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
	"Shopping_System/services"
	"Shopping_System/untils"
	"Shopping_System/untils/http"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func SearchGoodInLike(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if likelist, err := services.SearchGoodInLike(claims); err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, gin.H{
			"您的收藏如下": likelist,
		})
	}
}
func LikeAGood(c *gin.Context) {
	LikeGood := new(model.UserLike)
	if err := c.ShouldBind(LikeGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, LikeGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.LikeAGood(LikeGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message1": "已将Gid为" + strconv.FormatInt(LikeGood.Gid, 10) + "的商品添加到您的收藏中",
			"message2": "如若想查看商品详情，请移步用户商品查询页:http://127.0.0.1:8080/user/goods/search",
		})
	}
}
