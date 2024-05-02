package adaptor

import (
	"github.com/gin-gonic/gin"
	"strconv"
	app2 "uitranslate/app"
	"uitranslate/app/dto"
)

var (
	CategoryAppServ app2.ICategoryApplicationService = app2.NewCategoryApplicationService()
)

func RegisterHandler(engine *gin.Engine) {
	// category
	engine.GET("/admin/category/page", PageCategory)
	engine.GET("/admin/category", AllCategory)
	engine.POST("/admin/category", AddCategory)
	engine.PUT("/admin/category", UpdateCategory)

	// i18n query
	engine.GET("/api/category/:name", CategoryApiData)
}

func PageCategory(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 20
	}
	parentId, err := strconv.Atoi(c.Query("parent_id"))
	if err != nil {
		parentId = 0
		// todo
	}

	var req = dto.CategoryPageReq{
		Page:     int64(page),
		PageSize: int64(pageSize),
		ParentId: int64(parentId),
		Name:     c.Query("name"),
	}

	pageCategory, bizErr := CategoryAppServ.PageCategory(req)
	if bizErr != nil {
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "", "data": pageCategory})
}

func AllCategory(c *gin.Context) {
	parentId, err := strconv.Atoi(c.Param("parent_id"))
	if err != nil || parentId == 0 {
		parentId = 1
	}

	categories, bizErr := CategoryAppServ.AllCategory(int64(parentId))
	if bizErr != nil {
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "", "data": categories})
}

func AddCategory(c *gin.Context) {
	var req dto.AddCategoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	bizErr := CategoryAppServ.AddCategory(req)
	if bizErr != nil {
		return
	}
}

func UpdateCategory(c *gin.Context) {
	var req dto.UpdateCategoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	bizErr := CategoryAppServ.UpdateCategory(req)
	if bizErr != nil {
		return
	}
}

//====api

func CategoryApiData(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		return
	}
	data, err := CategoryAppServ.CategoryApiData(name)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
}
