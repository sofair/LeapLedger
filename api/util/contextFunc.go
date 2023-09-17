package util

import (
	"KeepAccount/global"
	userModel "KeepAccount/model/user"
	"KeepAccount/util"
	"github.com/gin-gonic/gin"
)

type _contextFunc interface {
	SetUserId(id uint, ctx *gin.Context)
	SetClaims(claims *util.CustomClaims, ctx *gin.Context)
	GetUserId(ctx *gin.Context) uint
	GetUser(ctx *gin.Context) (*userModel.User, error)
	GetToken(ctx *gin.Context) string
	GetClaims(ctx *gin.Context) util.CustomClaims
	GetClient(ctx *gin.Context) global.Client
}

type contextFunc struct {
}

var ContextFunc = new(contextFunc)

const (
	_UserId = "_user_id_"
	_Claims = "_claims_"
)

func (cf *contextFunc) SetUserId(id uint, ctx *gin.Context) {
	ctx.Set(_UserId, id)
}

func (cf *contextFunc) SetClaims(claims *util.CustomClaims, ctx *gin.Context) {
	ctx.Set(_Claims, claims)
}

func (cf *contextFunc) GetUserId(ctx *gin.Context) uint {
	return ctx.MustGet(_UserId).(uint)
}

func (cf *contextFunc) GetUser(ctx *gin.Context) (*userModel.User, error) {
	user := new(userModel.User)
	err := user.SelectById(cf.GetUserId(ctx))
	return user, err
}

func (cf *contextFunc) GetToken(ctx *gin.Context) string {
	return ctx.Request.Header.Get("authorization")
}

func (cf *contextFunc) GetClaims(ctx *gin.Context) util.CustomClaims {
	return ctx.MustGet(_Claims).(util.CustomClaims)
}

func (cf *contextFunc) GetClient(ctx *gin.Context) global.Client {
	userAgent := ctx.GetHeader("User-Agent")
	for clientType, client := range userModel.GetClients() {
		if client.CheckUserAgent(userAgent) {
			return clientType
		}
	}
	panic("Not found client")
}
