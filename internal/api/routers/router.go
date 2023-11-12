package router

import (
	handler "github.com/anthdm/gobank/internal/api/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	/*
		r := gin.Default()

		r.Use(gin.Logger())
		if gin.Mode() == gin.ReleaseMode {
			r.Use(middleware.Security())
			r.Use(middleware.Xss())
		}
		r.Use(middleware.Cors())
		r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 60)) // 60 requests per minute

		docs.SwaggerInfo.BasePath = "/api/v1"
		v1 := r.Group("/api/v1")
		{
			v1.GET("/", books.Healthcheck)
			v1.GET("/books", middleware.APIKeyAuth(), books.FindBooks)
			v1.POST("/books", middleware.APIKeyAuth(), middleware.JWTAuth(), books.CreateBook)
			v1.GET("/books/:id", middleware.APIKeyAuth(), books.FindBook)
			v1.PUT("/books/:id", middleware.APIKeyAuth(), books.UpdateBook)
			v1.DELETE("/books/:id", middleware.APIKeyAuth(), books.DeleteBook)

			v1.POST("/login", middleware.APIKeyAuth(), auth.LoginHandler)
			v1.POST("/register", middleware.APIKeyAuth(), auth.RegisterHandler)
		}
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	*/

	router := mux.NewRouter()
	router.HandleFunc("/login", handler.MakeHTTPHandleFunc(handler.Login)).Methods("POST")
	router.HandleFunc("/account", handler.MakeHTTPHandleFunc(handler.CreateAccount)).Methods("POST")
	router.HandleFunc("/account", handler.MakeHTTPHandleFunc(handler.GetAccount)).Methods("GET")
	router.HandleFunc("/account/{id}", handler.WithJWTAuth(handler.MakeHTTPHandleFunc(handler.GetAccountByID)))
	router.HandleFunc("/transfer", handler.MakeHTTPHandleFunc(handler.Transfer))

	return router
}
