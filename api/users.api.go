package api

import (
	"comunty/ms-auth/controller"
	"comunty/ms-auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterHandler(ctx *gin.Context) {
	var request model.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err := controller.RegisterUserController(&request, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func GetUserHandler(ctx *gin.Context) {
	var users []model.User
	var res interface{}
	var err error
	if res, err = controller.GetUsersController(ctx); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	cur := res.(*mongo.Cursor)
	defer cur.Close(ctx)
	if err = cur.Decode(&users); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, &users)
}
