package api

import (
	"errors"
	"net/http"

	"dousheng/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User db.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: -1, StatusMsg: "用户名或密码不能为空"},
		})
		return
	}
	token := username + password
	user := db.User{}
	if db.DB == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: -1, StatusMsg: "服务器异常"},
		})
		return
	}
	result := db.DB.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newUser := db.User{
			Name:     username,
			Password: password,
			Token:    token,
		}
		db.DB.Create(&newUser)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "注册成功"},
			UserId:   int64(newUser.ID),
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "该用户已存在"},
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user := db.User{}
	result := db.DB.Where("token = ?", token).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   int64(user.ID),
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	user := db.User{}
	result := db.DB.Where("token = ?", token).Find(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	}
}
