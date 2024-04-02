package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/waltherx/honda-backend/api/controller"
	"github.com/waltherx/honda-backend/bootstrap"
	"github.com/waltherx/honda-backend/domain"
	"github.com/waltherx/honda-backend/mongo"
	"github.com/waltherx/honda-backend/repository"
	"github.com/waltherx/honda-backend/usecase"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
