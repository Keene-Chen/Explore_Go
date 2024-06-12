package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	userService "restful_api/app/service"
	"restful_api/util"
	"strconv"

	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func init() {
	ginAutoRouter.Register(&User{})
}

type User struct {
}

// 分页
func (api *User) Pages(c *gin.Context) {
	condMap := make(map[string]interface{})
	//total :=userService.UserTotal(condMap)
	page := c.DefaultQuery("page", "1")
	pageIndex, err := strconv.Atoi(page)
	if err != nil {
		panic(err)
	}
	pageSize := 2
	users, total := userService.UserPage(condMap, pageIndex, pageSize)
	pages := util.Paginator(pageIndex, pageSize, total)
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"msg":   "ok",
		"data":  users,
		"pages": pages,
	})
}

// 列表 带传参 可根据条件查询
func (api *User) List(c *gin.Context) {
	userName := c.Query("name")
	condMap := make(map[string]interface{})

	if userName != "" {
		condMap["name"] = userName
	}
	list := userService.UserList(condMap)
	//结果序列化输出
	listJsons, _ := json.Marshal(list)

	fmt.Println("list==", list)
	fmt.Println("listJsons==", listJsons)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": string(listJsons),
	})
}

// 详细信息
func (api *User) UserInfo(c *gin.Context) {

}
