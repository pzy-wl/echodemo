package intf

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	return nil
}

//首字母小写的方法只能当前包访问,而大写的则可以根据路径进行访问
// e.POST("/save", save)
//func Save(c echo.Context) error {
//	// Get name and email
//	// 此处的姓名和邮箱都没有经过校验
//	name := c.FormValue("name")
//	email := c.FormValue("email")
//	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
//}
func Save(c echo.Context) error {
	log.Println("_____-----______")
	//暂时不能访问
	//curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:2020/save2
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}
	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
func UpdateUser(c echo.Context) error {
	return nil
}

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	println("123")
	return c.String(http.StatusOK, id)
}

func SaveUser(c echo.Context) error {
	return nil
}

//e.GET("/show", show)
func Show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
