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

func DeleteGoodInShoppingCart(c *gin.Context) {
	DeleteGood := new(model.DeleteAGoodFormShoppingCart)
	if err := c.ShouldBind(DeleteGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, DeleteGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.DeleteGoodInShoppingCart(DeleteGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "您已成功从购物车删除本商品",
		})
	}
}

func BuyAGoodInShoppingCart(c *gin.Context) {
	BuyA := new(model.BuyA)
	if err := c.ShouldBind(BuyA); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, BuyA),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.BuyAGoodInShoppingCart(BuyA, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "您已成功购买本商品",
		})
	}
}
func OperateSizeOfGood(c *gin.Context) {
	ReviseSize := new(model.ReviseSize)
	if err := c.ShouldBind(ReviseSize); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ReviseSize),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.OperateSizeOfGood(ReviseSize, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorSizeExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "已修改您购物车内选中的商品的规格",
		})
	}
}
func OperateNumberOfGood(c *gin.Context) {
	ReviseNumber := new(model.ReviseNumber)
	if err := c.ShouldBind(ReviseNumber); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ReviseNumber),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.OperateNumberOfGood(ReviseNumber, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "已修改您购物车内选中的商品的数量",
		})
	}
}
func BuyAllInShoppingCart(c *gin.Context) {
	BuyAll := new(model.BuyAll)
	if err := c.ShouldBind(BuyAll); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, BuyAll),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.BuyAllInShoppingCart(BuyAll, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "您购物车内的商品已全部购买",
		})
	}
}
func ClearShoppingCart(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ClearShoppingCart(claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "您的购物车已清空",
		})
	}
}
func LoveAGood(c *gin.Context) {
	LoveGood := new(model.ShoppingCart)
	if err := c.ShouldBind(LoveGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, LoveGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.LoveAGood(LoveGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorSizeExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message1": "已将Gid为" + strconv.FormatInt(int64(LoveGood.Gid), 10) + "的商品添加到您的收藏中",
			"message2": "如若想查看商品详情，请移步用户商品查询页:http://127.0.0.1:8080/user/shoppingcart/search",
		})
	}
}

func SearchGoodInShoppingCart(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if lovelist, err := services.SearchGoodInShoppingCart(claims); err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, gin.H{
			"您的购物车如下": lovelist,
		})
	}
}
