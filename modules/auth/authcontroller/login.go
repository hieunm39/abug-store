package authcontroller

import (
	"abug-store/common"
	"abug-store/components/appctx"
	"abug-store/components/hasher"
	"abug-store/components/tokenprovider/jwt"
	"abug-store/modules/auth/authmodel"
	"abug-store/modules/auth/authservice"
	"abug-store/modules/auth/authstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var userLogin authmodel.UserLogin

		if err := c.ShouldBind(&userLogin); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		
		md5 := hasher.NewMd5Hash()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())
		db := appCtx.GetDatabaseConnection()
		store := authstorage.NewSQLStorage(db)
		service := authservice.NewLoginService(tokenProvider, store, md5, 60*60*24*7)

		account, err := service.Login(c.Request.Context(), &userLogin)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}