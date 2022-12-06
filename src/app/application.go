package app

import (
	"github.com/BFDavidGamboa/bookstore_oauth-api/src/clients/cassandra"
	"github.com/BFDavidGamboa/bookstore_oauth-api/src/domain/access_token"
	"github.com/BFDavidGamboa/bookstore_oauth-api/src/http"
	"github.com/BFDavidGamboa/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	dbRepo := db.NewRepository()
	accTkn := access_token.NewService(dbRepo)
	atHandler := http.NewAccessTokenHandler(accTkn)
	//  atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
