package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"uitranslate/app/category"
	"uitranslate/app/category/dto"
)

var (
	CategoryAppServ = category.Impl
)

func RegisterCategoryHandler(engine *gin.Engine) {
	// category
	engine.GET("/admin/category/page", PageCategory)
	engine.GET("/admin/category/list", AllCategory)
	engine.POST("/admin/category", AddCategory)
	engine.PUT("/admin/category", UpdateCategory)

	engine.GET("/api/category/list/:name", CategoryListApiDataByName)
	engine.GET("/api/category/:id", CategoryApiDataById)
	engine.GET("/api/category/list", ApiAllCategory)
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
	}

	var req = dto.CategoryPageReq{
		Page:     int64(page),
		PageSize: int64(pageSize),
		ParentId: int64(parentId),
		Name:     c.Query("name"),
	}

	pageCategory, bizErr := CategoryAppServ.PageCategory(req)
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "", Data: pageCategory})
}

func AllCategory(c *gin.Context) {
	parentId, err := strconv.Atoi(c.Query("parent_id"))
	if err != nil {
		parentId = 0
	}

	categories, bizErr := CategoryAppServ.AllCategoryByParentId(int64(parentId))
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "", Data: categories})
}

func AddCategory(c *gin.Context) {
	var req dto.AddCategoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: "Invalid request"})
		return
	}
	bizErr := CategoryAppServ.AddCategory(req)
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Category added successfully"})
}

//func BatchAddCategory(c *gin.Context) {
//	var reqs []dto.AddCategoryReq
//	c.ShouldBindJSON(&reqs)
//	//if err != nil {
//	//	c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: "Invalid request"})
//	//	return
//	//}
//	for _, req := range reqs {
//		bizErr := CategoryAppServ.AddCategory(req)
//		if bizErr != nil {
//			c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
//			return
//		}
//	}
//
//	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Category added successfully"})
//}

func UpdateCategory(c *gin.Context) {
	var req dto.UpdateCategoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: "Invalid request"})
		return
	}
	bizErr := CategoryAppServ.UpdateCategory(req)
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Category updated successfully"})
}

//====api

func CategoryApiDataById(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, NewApiRestResult(RestResult{Code: -1, Message: "Invalid category Id"}))
		return
	}
	data, err := CategoryAppServ.CategoryApiDataById(int64(categoryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewApiRestResult(RestResult{Code: -1, Message: err.Error()}))
		return
	}
	c.JSON(http.StatusOK, NewApiRestResult(RestResult{Code: 0, Message: "", Data: data}))
}

func CategoryListApiDataByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, NewApiRestResult(RestResult{Code: -1, Message: "Name is required"}))
		return
	}
	language := c.GetHeader("Talentpool-Language")
	if language == "" {
		language = "english"
	}
	data, err := CategoryAppServ.ListCategoryByParentName(name, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewApiRestResult(RestResult{Code: -1, Message: err.Error()}))
		return
	}
	c.JSON(http.StatusOK, NewApiRestResult(RestResult{Code: 0, Message: "", Data: data}))
}

func ApiAllCategory(c *gin.Context) {
	parentId, err := strconv.Atoi(c.Query("parent_id"))
	if err != nil {
		parentId = 0
	}

	categories, bizErr := CategoryAppServ.AllCategoryByParentId(int64(parentId))
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, NewApiRestResult(RestResult{Code: -1, Message: bizErr.Error()}))
		return
	}
	c.JSON(http.StatusOK, NewApiRestResult(RestResult{Code: 0, Message: "", Data: categories}))
}
