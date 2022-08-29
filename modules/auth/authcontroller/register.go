package authcontroller

import (
	"abug-store/common"
	"abug-store/components/appctx"
	"abug-store/components/hasher"
	"abug-store/modules/auth/authmodel"
	"abug-store/modules/auth/authservice"
	"abug-store/modules/auth/authstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetDatabaseConnection()
		var data authmodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		md5 := hasher.NewMd5Hash()
		storage := authstorage.NewSQLStorage(db)
		service := authservice.NewRegisterService(storage, md5)
		
		if err := service.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
