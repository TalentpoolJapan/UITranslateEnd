package adaptor

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"uitranslate/cms/app"
	"uitranslate/cms/app/dto"
)

var (
	CategoryAppServ app.ICategoryApplicationService = app.NewCategoryApplicationService()
)

func RegisterHandler(engine *gin.Engine) {
	// category
	engine.GET("/admin/category/page", AllCategory)
	engine.POST("/admin/category", AddCategory)
	engine.PUT("/admin/category", UpdateCategory)

	// i18n query
	engine.GET("/api/category/:name", CategoryApiData)
}

func AllCategory(c *gin.Context) {
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
