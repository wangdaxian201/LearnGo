package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name" form-data:"name"`
	Age  int    `json:"age" form-data:"age"`
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func JsonDemo(c *gin.Context) {
	user := &User{}
	if errUser := c.ShouldBindJSON(user); errUser != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": `response error`,
		})
	}

	fmt.Printf("name:%v, age: %v \n", user.Name, user.Age)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": `the body should be User`,
	})
}

func main() {
	fmt.Println("learn gin....")
	r := gin.Default()
	r.GET("/ping", Ping)
	r.POST("/JsonDemo", JsonDemo)
	// router group

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
