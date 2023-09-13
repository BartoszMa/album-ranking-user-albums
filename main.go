package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//router.Use(cors.New(cors.Config{
	//	//AllowOrigins: []string{os.Getenv(test)},
	//	AllowMethods: []string{"POST", "GET", "DELETE"},
	//}))

	errRouter := router.Run()
	//errRouter := router.Run("samplehost:9999")
	if errRouter != nil {
		panic(errRouter)
	}
}
