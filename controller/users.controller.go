package controller

import (
	"comunty/ms-auth/conf"
	"comunty/ms-auth/model"
	"comunty/ms-auth/util"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterUserController(req *model.RegisterRequest, ctx *gin.Context) (interface{}, error) {
	user := req.User
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	db := util.Mongo{
		DB:         conf.GlobalConf.Credentials.Persist.MongoDB.Database,
		Collection: "users",
	}
	res, err := db.InsertOne(ctx.Request.Context(), user)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func GetUsersController(ctx *gin.Context) (interface{}, error) {
	var filter interface{}
	db := util.Mongo{
		DB:         conf.GlobalConf.Credentials.Persist.MongoDB.Database,
		Collection: "users",
	}
	opts := options.FindOptions{}
	res, err := db.Find(filter, &opts, ctx.Request.Context())
	if err != nil {
		return nil, err
	}
	return res, nil
}
