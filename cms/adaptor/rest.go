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
	engine.GET("/category/page", AllCategory)
	engine.POST("/category/add", AddCategory)
	engine.PUT("/category/update", UpdateCategory)
}

func AllCategory(c *gin.Context) {
	var err error
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return
	}
	var req = dto.CategoryPageReq{
		Page:     int64(page),
		PageSize: int64(pageSize),
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
