package main

import (
	"log"
	"net/http"
	// 注意这里的路径
	echo "github.com/labstack/echo/v4"

	"study-echo/intf"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	// 得到一个 echo.Echo 的实例
	e := echo.New()
	// 注册路由
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", intf.SaveUser)
	e.GET("/users/:id", intf.GetUser)
	//允许同一路由定义两种不同的访问方式,get和post,根据URL提交方式的不同进行区分ç
	e.POST("/show", intf.Show)
	e.GET("/show", intf.Show)
	e.POST("/save", intf.Save)
	log.Println("_____-----______")
	e.POST("/save", intf.Save)
	e.PUT("/users/:id", intf.UpdateUser)
	e.DELETE("/users/:id", intf.DeleteUser)
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
		// or
		// return c.XML(http.StatusCreated, u)
	})
	// 开启 HTTP Server
	e.Logger.Fatal(e.Start(":2020"))
}
