package api

import (
	"fmt"
	"log"
	"myapi/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Routes struct {
	Path    string
	Handler []func(*gin.RouterGroup, *gorm.DB)
}

var router = []Routes{
	{"/teach", []func(*gin.RouterGroup, *gorm.DB){getTeach, createTeach}},
	{"/major", []func(*gin.RouterGroup, *gorm.DB){getMajor}},
	{"/join-topic", []func(*gin.RouterGroup, *gorm.DB){getJoinTopic}},
}

func RunServer(db *gorm.DB) {
	r := gin.Default()

	authMiddleware, err := getAuthMiddleware(db)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	v1 := r.Group("/api/v1")
	registerHandler(v1, db)
	loginHandler(v1, authMiddleware.LoginHandler)
	v1.Use(authMiddleware.MiddlewareFunc())
	for _, route := range router {
		for _, controller := range route.Handler {
			controller(v1.Group(route.Path), db)
		}
	}

	// use ginSwagger middleware to serve the API docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + utils.PORT)
}

var validate *validator.Validate

func CheckInputError(input interface{}) error {
	validate = validator.New()
	err := validate.Struct(input)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	fields := []string{}
	for _, err := range err.(validator.ValidationErrors) {
		// Dang An: err.Namespace() can differ when a custom TagNameFunc is registered or not (In this case, we use the name from json tag)
		fields = append(fields, err.Namespace()+" "+err.Type().Name())
	}
	return fmt.Errorf("Error:Field validation for: %v", strings.Join(fields, ", "))
}
