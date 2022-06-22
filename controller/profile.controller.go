package controller

import (
	"comunty/ms-auth/conf"
	"comunty/ms-auth/model"
	"comunty/ms-auth/util"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertProfileController(req *model.RegisterRequest, reqHeaders http.Header, ctx *gin.Context) error {
	user := req.User
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	mongo := util.Mongo{
		DB:         conf.GlobalConf.Credentials.Persist.MongoDB.Database,
		Collection: "profile",
	}
	if _, err := mongo.InsertOne(ctx.Request.Context(), user); err != nil {
		return err
	}
	return nil
}
