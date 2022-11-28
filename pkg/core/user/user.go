package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	apiTypes "github.com/dnsjia/luban/api/types"
	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
	"github.com/dnsjia/luban/pkg/types"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.ERROR,
			"msg":  types.ParamsError,
			"data": "",
		})
		return
	}

	if err := options.DB.Model(&model.User{}).Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.ERROR,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.SUCCESS,
		"msg":  "注册成功",
		"data": "",
	})
}

func Login(c *gin.Context) {
	var (
		userRequest apiTypes.UserRequest
		user        model.User
	)
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.ERROR,
			"msg":  types.ParamsError,
			"data": "",
		})
		return
	}

	// Ldap方式登陆, true
	if userRequest.Ldap {
		user, err := LdapLogin(userRequest.Username, userRequest.Password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": types.ERROR,
				"msg":  err.Error(),
				"data": "",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": types.SUCCESS,
			"data": user,
			"msg":  types.UserLoginSuccess,
		})
		return

	}

	// 普通方式登陆
	if err := options.DB.Where("username = ?", userRequest.Username).First(&user).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": types.ERROR,
			"msg":  types.UserLoginFailed,
			"data": "",
		})
		return
	}

	// 判断用户密码是否一致
	if userRequest.Password != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"code": types.ERROR,
			"msg":  types.UserPasswordError,
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.SUCCESS,
		"data": user,
		"msg":  types.UserLoginSuccess,
	})
}

func Manage(c *gin.Context) {
	var users []model.User
	if err := options.DB.Model(&model.User{}).Find(&users).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": types.SUCCESS,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": types.SUCCESS,
		"data": users,
		"msg":  types.Success,
	})
}
