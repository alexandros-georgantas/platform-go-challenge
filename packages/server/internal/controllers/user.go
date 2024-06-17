package controllers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type UserController interface {
// 	SignUp() gin.HandlerFunc
// 	Login() gin.HandlerFunc
// }

// type usersController struct {
// 	userService services.UserService
// }

// func (uc *usersController) Register() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := c.BindJSON(&user)
// 		if err != nil {
// 			c.PureJSON(http.StatusBadRequest, gin.H{
// 				"msg": "your request is wrong",
// 			})

// 		}

// 		userId, err := uc.userService.Register(&user)
// 		if err != nil {

// 			c.PureJSON(http.StatusInternalServerError, gin.H{
// 				"msg": "oops something is wrong with our server",
// 			})

// 		}
// 		c.PureJSON(http.StatusCreated, gin.H{
// 			"msg": "done",
// 		})

// 	}
// }

// func (uc *usersController) Login() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var login contracts.Login
// 		err := c.BindJSON(&login)
// 		if err != nil {
// 			api.BadRequest(c)
// 			return
// 		}

// 		tokenId, err := uc.userService.Login(login.Email, login.Password)
// 		if err != nil {
// 			api.InternalServerError(c)
// 			return
// 		}

// 		api.OK(c, gin.H{"tokenId": tokenId})
// 	}
// }
