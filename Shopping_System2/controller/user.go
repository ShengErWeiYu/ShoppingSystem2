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
)

func DisRegister(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.DisRegister(claims); err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, gin.H{
			"message": "注销账户成功，再见！",
		})
	}
}
func ReviseSex(c *gin.Context) {
	ReviseSexUser := new(model.ReviseSexUser)
	if err := c.ShouldBind(ReviseSexUser); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ReviseSexUser),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ReviseSex(ReviseSexUser, claims); err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, gin.H{
			"message": "修改性别成功!",
		})
	}
}
func RevisePassword(c *gin.Context) {
	RevisePasswordUser := new(model.RevisePasswordUser)
	if err := c.ShouldBind(RevisePasswordUser); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, RevisePasswordUser),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.RevisePassword(RevisePasswordUser, claims); err != nil {
		if errors.Is(err, mysql.ErrorAnswer) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改密码成功!",
		})
	}
}
func ReviseUsername(c *gin.Context) {
	ReviseNameUser := new(model.ReviseNameUser)
	if err := c.ShouldBind(ReviseNameUser); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ReviseNameUser),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ReviseUsername(ReviseNameUser, claims); err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, gin.H{
			"message": "修改用户名成功！",
		})
	}
}
func GetSecurityQuestion(c *gin.Context) {
	GetSecurityQuestion := new(model.GetSecurityQuestion)
	if err := c.ShouldBind(GetSecurityQuestion); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, GetSecurityQuestion),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if SecurityQuestion, err := services.GetSecurityQuestionFromRedis(GetSecurityQuestion); err != nil {
		log.Println("缓存中未找到相关数据！")
	} else {
		c.JSON(200, gin.H{
			"message":          "此条数据来自缓存！",
			"SecurityQuestion": SecurityQuestion,
		})
		return
	}
	if SecurityQuestion, err1, err2 := services.GetSecurityQuestion(GetSecurityQuestion); err1 != nil && err2 == nil {
		if errors.Is(err1, mysql.ErrorUserNotExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err1.Error(),
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"SecurityQuestion": SecurityQuestion,
		})
	}
}
func ForgetPassword(c *gin.Context) {
	FgtPswUser := new(model.FgtPswUser)
	if err := c.ShouldBind(FgtPswUser); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, FgtPswUser),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if err := services.ForgetPassword(FgtPswUser); err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
			return
		}
		if errors.Is(err, mysql.ErrorAnswer) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"message": "重置密码成功！",
		})
	}

}

func Register(c *gin.Context) {
	RegisterUser := new(model.User)
	if err := c.ShouldBind(RegisterUser); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, RegisterUser),
		})
		return
	} else {
		http.RespSuccess(c, nil)

	}
	if err := services.Register(RegisterUser); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"message": "注册成功！欢迎" + RegisterUser.Username + "加入我们!",
		})
	}
}

func Login(c *gin.Context) {
	LoginUser := new(model.LoginUser)
	if err := c.ShouldBind(LoginUser); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, LoginUser),
		})
		return
	} else {
		http.RespSuccess(c, nil)
		if err := services.Login(LoginUser); err != nil {
			if errors.Is(err, mysql.ErrorUserNotExist) {
				c.JSON(200, gin.H{
					"ErrorMessage": err.Error(),
				})
			}
			if errors.Is(err, mysql.ErrorPassword) {
				c.JSON(200, gin.H{
					"ErrorMessage": err.Error(),
				})
			}
		} else {
			tokenString, _ := untils.GenToken(LoginUser.Uid)
			c.JSON(200, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": gin.H{"token": tokenString},
			})
			username := services.GetUsernameByUid(LoginUser.Uid)
			c.JSON(200, gin.H{
				"message": "登录成功，欢迎" + username,
			})
		}
	}
}
