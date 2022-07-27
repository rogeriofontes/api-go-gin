package controller

import (
	"net/http"

	"github.com/rogeriofontes/api-go-gin/api/service"
	"github.com/rogeriofontes/api-go-gin/models"
	"github.com/rogeriofontes/api-go-gin/util"
	"github.com/rogeriofontes/api-go-gin/util/token"

	"github.com/gin-gonic/gin"
)

//UserController struct
type UserController struct {
	service service.UserService
}

//NewUserController : NewUserController
func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

//CreateUser ->  calls CreateUser services for validated user
func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBind(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}

	hashPassword, _ := util.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	util.SuccessJSON(c, http.StatusOK, "Successfully Created user")
}

//LoginUser : Generates JWT Token for validated user
func (u *UserController) LoginUser(c *gin.Context) {
	var user models.UserLogin
	//var hmacSampleSecret []byte
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}
	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}
	/*token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})*/

	token, err := token.GenerateToken(dbUser.ID)

	//tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &util.ResponseLogin{
		Success: true,
		Message: "Token generated sucessfully",
		Token:   token,
	}
	c.JSON(http.StatusOK, response)
}
