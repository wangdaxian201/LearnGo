package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
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
	// 强制日志颜色化
	gin.ForceConsoleColor()
	gin.SetMode("debug")
	router := gin.Default()
	// router group
	router.GET("/ping", Ping)
	v1 := router.Group("/v1")
	{
		v1.POST("/JsonDemo", JsonDemo)
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	// 不能热重载的run
	//err := r.Run(":8080")
	//if err != nil {
	//	_ = fmt.Errorf("server error: %v", err)
	//	return
	//} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
