package controllers

import (
	"encoding/json"
	"io"
	// "reflect"

	// "repeatro/internal/models"
	"repeatro/internal/schemes"
	"repeatro/internal/services"

	// "repeatro/internal/schemes"
	"repeatro/internal/security"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserService services.UserServiceInterface
	security    *security.Security
}

func CreateNewUserController(userService *services.UserService, security *security.Security) *UserController {
	return &UserController{
		UserService: userService,
		security:    security,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	var userRegister schemes.AuthUser
	if err = json.Unmarshal(body, &userRegister); err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	token, err := uc.UserService.Register(userRegister)

	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(200, token)
}

func (uc *UserController) Login(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	//get email and password
	var userLogin schemes.AuthUser
	if err = json.Unmarshal(body, &userLogin); err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	token, err := uc.UserService.Login(userLogin)

	if err != nil {
		ctx.AbortWithError(500, err)
		return 
	}

	ctx.JSON(200, token)
	
	// authorization := ctx.Request.Header.Get("Authorization")
	// token := authorization[7:]
	// //decode it with secutiry methods 
	// claims := uc.security.DecodeToken(token)

	// userIf 
	//check in db that existst
	//accept
	//reject
	// ctx.JSON(200, a[6:])
}




