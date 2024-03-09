package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	MYSQL_HOST    = "tcp(127.0.0.1:3306)"
	MYSQL_SECRECT = "yYVim5WbqzkWziNY"
	//EEpLWKlYixYtYGSx
	//MYSQL_HOST    = "tcp(192.168.1.165:3306)"
	//MYSQL_SECRECT = "a"
	MYSQL_DB           = "talentpool"
	DEEPL_FREE_API_KEY = "ed8fb40e-858f-7167-44c8-65ec333131c2:fx"

	DB, _ = xorm.NewEngine("mysql", fmt.Sprintf("root:%s@%s/%s?charset=utf8", MYSQL_SECRECT, MYSQL_HOST, MYSQL_DB))
)

// 设置跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,TalentPool-Language,token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// 后台权限验证中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO 权限验证
		c.Next()
	}
}

func main() {
	DB.ShowSQL(true)
	r := gin.Default()
	r.Use(Cors())
	authorized := r.Group("/uitranslate")
	authorized.Use(AuthRequired())
	{
		// 获取翻译分类列表
		// 分类列表应该添加该分类的说明
		// 下拉分类ID从5开始 往上的属于零时用不放在生产环境 目前是5-14
		// 界面语言ID从100开始
		// 1 下拉选项翻译固定，有自己的结构需要单独处理
		// 2 界面语言预留100个分类这里面的结构体是一样的
		// 3 Jobcategory工作分类需要单独拉出来做界面因为结构体不一样
		authorized.GET("/get/selectcategorylist", GetTranslateSelectCategory)
		authorized.GET("/get/uicategorylist", GetTranslateUICategory)
		// 修改分类列表
		authorized.POST("/update/selectcategorylist", UpdateTranslateCategory)
		authorized.POST("/update/uicategorylist", UpdateTranslateCategory)
		// 获取分类下面的翻译内容
		authorized.GET("/get/selectcategorylist/classid/:id", GetTranslateSelectByClassid)
		authorized.POST("/add/selectcategorybyclassid", AddTranslateSelectByClassid)
		authorized.POST("/update/selectcategorybyclassid", UpdateTranslateSelectByClassid)
		authorized.POST("/delete/selectcategorybyclassid", DeleteTranslateSelectByClassid)

		authorized.GET("/get/uicategorylist/classid/:id", GetTranslateUIByClassid)
		authorized.POST("/add/uicategorybyclassid", AddTranslateUIByClassid)
		authorized.POST("/update/uicategorybyclassid", UpdateTranslateUIByClassid)
		authorized.POST("/delete/uicategorybyclassid", DeleteTranslateUIByClassid)

		authorized.GET("/get/jobcategoryclass", GetJobCategoryClass)
		authorized.GET("/get/jobcategorysubclass/:id", GetJobCategorySubClass)
		authorized.POST("/update/jobcategoryclass", EditJobCategoryClass)
		authorized.POST("/update/jobcategorysubclass", EditJobCategorySubClass)
		authorized.POST("/add/jobcategoryclass", AddJobCategoryClass)
		authorized.POST("/add/jobcategorysubclass", AddJobCategorySubClass)
		authorized.POST("/delete/jobcategoryclass", DeleteJobCategoryClass)
		authorized.POST("/delete/jobcategorysubclass", DeleteJobCategorySubClass)

		authorized.POST("/en2ja", TranslateEnglishToJapanese)
		authorized.POST("/ja2en", TranslateJapaneseToEnglish)

	}

	r.Run(":8332")

}

// 获取分类
type TranslateClass struct {
	Id        int    `json:"id" binding:"required"`
	Classname string `json:"classname" binding:"required"`
	Tag       string `json:"tag" binding:"required"`
}

func GetTranslateSelectCategory(c *gin.Context) {
	var _TranslateClass []TranslateClass
	err := DB.Where("id>=? and id<?", 5, 14).Find(&_TranslateClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error(), "data": _TranslateClass})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "", "data": _TranslateClass})
}
func GetTranslateUICategory(c *gin.Context) {
	var _TranslateClass []TranslateClass
	err := DB.Where("id>=? and id<?", 100, 110).Find(&_TranslateClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error(), "data": _TranslateClass})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "", "data": _TranslateClass})
}

// 更新分类
func UpdateTranslateCategory(c *gin.Context) {
	var _TranslateClass TranslateClass
	err := c.ShouldBindJSON(&_TranslateClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	_, err = DB.Table("translate_class").Where("id=?", _TranslateClass.Id).Update(_TranslateClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 0, "msg": ""})
}

type GetTranslateClassid struct {
	ID int `uri:"id" binding:"required"`
}

// 获取分类下面的翻译内容

// 分类是一个单独的独立结构
type JobCategory struct {
	Id       int    `json:"id" binding:"required"`
	NameEn   string `json:"english" binding:"required"`
	NameJa   string `json:"japanese" binding:"required"`
	Parentid int    `json:"parentid"`
}

// ////以下是独立结构表
type Industry struct {
	Id         int    `json:"id"`
	IndustryEn string `json:"english" binding:"required"`
	IndustryJa string `json:"japanese" binding:"required"`
}
type CompanyType struct {
	Id            int    `json:"id"`
	CompanyTypeEn string `json:"english" binding:"required"`
	CompanyTypeJa string `json:"japanese" binding:"required"`
}
type Country struct {
	Id        int    `json:"id"`
	CountryEn string `json:"english" binding:"required"`
	CountryJa string `json:"japanese" binding:"required"`
}
type Education struct {
	Id     int    `json:"id"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type JobType struct {
	Id     int    `json:"id"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type Languages struct {
	Id     int    `json:"id"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type WorkStyle struct {
	Id     int    `json:"id"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type Salary struct {
	Id     int    `json:"id"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type WorkExp struct {
	Id     int    `json:"id"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type JapanCity struct {
	Id     int    `json:"id"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}

// /////////////////这部分对应网页界面语言////////////////////////
type UITranslate struct {
	Id       int    `json:"id"`
	Transkey string `json:"transkey"`
	English  string `json:"english" binding:"required"`
	Japanese string `json:"japanese" binding:"required"`
	Classid  int    `json:"classid" binding:"required"`
}

/////////////////////////////////////////////

type CommonSelect struct {
	Id      int    `json:"id"`
	NameEn  string `json:"english" binding:"required"`
	NameJa  string `json:"japanese" binding:"required"`
	Classid int    `json:"classid" binding:"required"`
}

type CommonSelectDelete struct {
	Id      int `json:"id"`
	Classid int `json:"classid" binding:"required"`
}

/////////////////////////////////////////////

// /////////////////JobCategory///////////////////////
// 获取工作一级分类
func GetJobCategoryClass(c *gin.Context) {
	var _JobCategory []JobCategory
	err := DB.Table("job_category").Where("parentid=?", 0).Find(&_JobCategory)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "", "data": _JobCategory})
}

// 获取工作二级分类
type JobCategorySubClass struct {
	Parentid int `uri:"id" bindind:"required"`
}

func GetJobCategorySubClass(c *gin.Context) {
	var (
		_JobCategorySubClass JobCategorySubClass
		_JobCategory         []JobCategory
	)
	err := c.ShouldBindUri(&_JobCategorySubClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	err = DB.Table("job_category").Where("parentid=?", _JobCategorySubClass.Parentid).Find(&_JobCategory)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "", "data": _JobCategory})
}

// 添加工作一级分类
type InsertJobCategoryClass struct {
	Parentid int    `json:"parentid"`
	NameEn   string `json:"english" binding:"required"`
	NameJa   string `json:"japanese" binding:"required"`
}

func AddJobCategoryClass(c *gin.Context) {
	var _InsertJobCategoryClass InsertJobCategoryClass
	err := c.ShouldBindJSON(&_InsertJobCategoryClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	_, err = DB.Table("job_category").Insert(&_InsertJobCategoryClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

// 添加工作二级分类
type InsertJobCategorySubClass struct {
	NameEn   string `json:"english" binding:"required"`
	NameJa   string `json:"japanese" binding:"required"`
	Parentid int    `json:"parentid" binding:"required"`
}

func AddJobCategorySubClass(c *gin.Context) {
	var _InsertJobCategorySubClass InsertJobCategorySubClass
	err := c.ShouldBindJSON(&_InsertJobCategorySubClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	_, err = DB.Table("job_category").Insert(&_InsertJobCategorySubClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

// 修改工作一级分类
type UpdateJobCategoryClass struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}

func EditJobCategoryClass(c *gin.Context) {
	var _UpdateJobCategoryClass UpdateJobCategoryClass
	err := c.ShouldBindJSON(&_UpdateJobCategoryClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	_, err = DB.Table("job_category").Where("id=?", _UpdateJobCategoryClass.Id).And("parentid=0").Update(&_UpdateJobCategoryClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

// 修改二级分类
type UpdateJobCategorySubClass struct {
	Id       int    `json:"id" binding:"required"`
	NameEn   string `json:"english" binding:"required"`
	NameJa   string `json:"japanese" binding:"required"`
	Parentid int    `json:"parentid" binding:"required"`
}

func EditJobCategorySubClass(c *gin.Context) {
	var _UpdateJobCategorySubClass UpdateJobCategorySubClass
	err := c.ShouldBindJSON(&_UpdateJobCategorySubClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	_, err = DB.Table("job_category").Where("id=?", _UpdateJobCategorySubClass.Id).And("parentid=?", _UpdateJobCategorySubClass.Parentid).And("parentid!=0").Update(&_UpdateJobCategorySubClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

// 删除一级分类
type RemoveJobCategoryClass struct {
	Id int `json:"id" binding:"required"`
}

func DeleteJobCategoryClass(c *gin.Context) {
	var (
		_RemoveJobCategoryClass RemoveJobCategoryClass
		_JobCategory            []JobCategory
	)
	err := c.ShouldBindJSON(&_RemoveJobCategoryClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	err = DB.Table("job_category").Where("parentid=?", _RemoveJobCategoryClass.Id).Limit(1).Find(&_JobCategory)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	if len(_JobCategory) > 1 {
		c.JSON(200, gin.H{"status": 1, "msg": "delete subclass first"})
	}
	_, err = DB.Where("id=?", _RemoveJobCategoryClass.Id).And("parentid=0").Delete(&JobCategory{})
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

// 删除二级分类
type RemoveJobCategorySubClass struct {
	Id int `json:"id" binding:"required"`
}

func DeleteJobCategorySubClass(c *gin.Context) {
	var _RemoveJobCategorySubClass RemoveJobCategorySubClass
	err := c.ShouldBindJSON(&_RemoveJobCategorySubClass)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	_, err = DB.Where("id=?", _RemoveJobCategorySubClass.Id).And("parentid!=0").Delete(&JobCategory{})
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////
func GetTranslateSelectByClassid(c *gin.Context) {
	var _GetTranslateClassid GetTranslateClassid
	err := c.ShouldBindUri(&_GetTranslateClassid)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	classid := _GetTranslateClassid.ID
	if classid == 5 {
		data, err := GetJobType()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 6 {
		data, err := GetLanguages()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 7 {
		data, err := GetWorkStyle()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 8 {
		data, err := GetEducation()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 9 {
		data, err := GetSalary()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 10 {
		data, err := GetIndustry()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 11 {
		data, err := GetCompanyType()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 12 {
		data, err := GetCountry()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	if classid == 13 {
		data, err := GetJapanCity()
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	c.JSON(200, gin.H{"status": 1, "msg": "ID not in range"})
}

func AddTranslateSelectByClassid(c *gin.Context) {
	var _CommonSelect CommonSelect
	err := c.ShouldBindJSON(&_CommonSelect)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	classid := _CommonSelect.Classid
	if classid < 5 || classid > 13 {
		c.JSON(200, gin.H{"status": 1, "msg": "ID not in range"})
		return
	}
	if classid == 5 {
		err := AddJobType(JobType{NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 6 {
		err := AddLanguages(Languages{NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 7 {
		err := AddWorkStyle(WorkStyle{NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 8 {
		err := AddEducation(Education{NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 9 {
		err := AddSalary(Salary{NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 10 {
		err := AddIndustry(Industry{IndustryEn: _CommonSelect.NameEn, IndustryJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 11 {
		err := AddCompanyType(CompanyType{CompanyTypeEn: _CommonSelect.NameEn, CompanyTypeJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 12 {
		err := AddCountry(Country{CountryEn: _CommonSelect.NameEn, CountryJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 13 {
		err := AddJapanCity(JapanCity{NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
}

func UpdateTranslateSelectByClassid(c *gin.Context) {
	var _CommonSelect CommonSelect
	err := c.ShouldBindJSON(&_CommonSelect)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	classid := _CommonSelect.Classid
	if classid < 5 || classid > 13 {
		c.JSON(200, gin.H{"status": 1, "msg": "ID not in range"})
		return
	}
	if classid == 5 {
		err := UpdateJobType(JobType{Id: _CommonSelect.Id, NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 6 {
		err := UpdateLanguages(Languages{Id: _CommonSelect.Id, NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 7 {
		err := UpdateWorkStyle(WorkStyle{Id: _CommonSelect.Id, NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 8 {
		err := UpdateEducation(Education{Id: _CommonSelect.Id, NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 9 {
		err := UpdateSalary(Salary{Id: _CommonSelect.Id, NameEn: _CommonSelect.NameEn, NameJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 10 {
		err := UpdateIndustry(Industry{Id: _CommonSelect.Id, IndustryEn: _CommonSelect.NameEn, IndustryJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 11 {
		err := UpdateCompanyType(CompanyType{Id: _CommonSelect.Id, CompanyTypeEn: _CommonSelect.NameEn, CompanyTypeJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 12 {
		err := UpdateCountry(Country{Id: _CommonSelect.Id, CountryEn: _CommonSelect.NameEn, CountryJa: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 13 {
		err := UpdateJapanCity(JapanCity{Id: _CommonSelect.Id, NameJa: _CommonSelect.NameEn, NameEn: _CommonSelect.NameJa})
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
}

func DeleteTranslateSelectByClassid(c *gin.Context) {
	var _CommonSelect CommonSelectDelete
	err := c.ShouldBindJSON(&_CommonSelect)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	classid := _CommonSelect.Classid
	if classid < 5 || classid > 13 {
		c.JSON(200, gin.H{"status": 1, "msg": "ID not in range"})
		return
	}
	if classid == 5 {
		err := DeleteJobType(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 6 {
		err := DeleteLanguages(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 7 {
		err := DeleteWorkStyle(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 8 {
		err := DeleteEducation(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 9 {
		err := DeleteSalary(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 10 {
		err := DeleteIndustry(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 11 {
		err := DeleteCompanyType(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 12 {
		err := DeleteCountry(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
	if classid == 13 {
		err := DeleteJapanCity(_CommonSelect.Id)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "ok"})
	}
}

func GetTranslateUIByClassid(c *gin.Context) {
	var _GetTranslateClassid GetTranslateClassid
	err := c.ShouldBindUri(&_GetTranslateClassid)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	classid := _GetTranslateClassid.ID
	if classid >= 100 {
		data, err := GetUITranslate(classid)
		if err != nil {
			c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
		return
	}
	c.JSON(200, gin.H{"status": 1, "msg": "ID not in range"})
}
func AddTranslateUIByClassid(c *gin.Context) {
	var _UITranslate UITranslate
	err := c.ShouldBindJSON(&_UITranslate)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	err = AddUITranslate(_UITranslate)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}
func UpdateTranslateUIByClassid(c *gin.Context) {
	var _UITranslate UITranslate
	err := c.ShouldBindJSON(&_UITranslate)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	err = UpdateUITranslate(_UITranslate)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

type UITranslateDelete struct {
	Id      int `json:"id"`
	Classid int `json:"classid" binding:"required"`
}

func DeleteTranslateUIByClassid(c *gin.Context) {
	var _UITranslate UITranslateDelete
	err := c.ShouldBindJSON(&_UITranslate)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	err = DeleteUITranslate(_UITranslate.Id)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

func GetJobType() (data []JobType, err error) {
	err = DB.Table("job_type").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddJobType(data JobType) (err error) {
	_, err = DB.Table("job_type").Insert(&data)
	return err
}
func UpdateJobType(data JobType) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("job_type").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteJobType(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&JobType{})
	return err
}
func GetLanguages() (data []Languages, err error) {
	err = DB.Table("languages").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddLanguages(data Languages) (err error) {
	_, err = DB.Table("languages").Insert(&data)
	return err
}
func UpdateLanguages(data Languages) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("languages").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteLanguages(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&Languages{})
	return err
}
func GetWorkStyle() (data []WorkStyle, err error) {
	err = DB.Table("work_style").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddWorkStyle(data WorkStyle) (err error) {
	_, err = DB.Table("work_style").Insert(&data)
	return err
}
func UpdateWorkStyle(data WorkStyle) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("work_style").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteWorkStyle(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&WorkStyle{})
	return err
}
func GetEducation() (data []Education, err error) {
	err = DB.Table("education").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddEducation(data Education) (err error) {
	_, err = DB.Table("education").Insert(&data)
	return err
}
func UpdateEducation(data Education) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("education").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteEducation(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&Education{})
	return err
}
func GetSalary() (data []Salary, err error) {
	err = DB.Table("salary").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddSalary(data Salary) (err error) {
	_, err = DB.Table("salary").Insert(&data)
	return err
}
func UpdateSalary(data Salary) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("salary").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteSalary(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&Salary{})
	return err
}
func GetIndustry() (data []Industry, err error) {
	err = DB.Table("industry").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddIndustry(data Industry) (err error) {
	_, err = DB.Table("industry").Insert(&data)
	return err
}
func UpdateIndustry(data Industry) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("industry").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteIndustry(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&Industry{})
	return err
}
func GetCompanyType() (data []CompanyType, err error) {
	err = DB.Table("company_type").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddCompanyType(data CompanyType) (err error) {
	_, err = DB.Table("company_type").Insert(&data)
	return err
}
func UpdateCompanyType(data CompanyType) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("company_type").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteCompanyType(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&CompanyType{})
	return err
}
func GetCountry() (data []Country, err error) {
	err = DB.Table("country").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddCountry(data Country) (err error) {
	_, err = DB.Table("country").Insert(&data)
	return err
}
func UpdateCountry(data Country) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("country").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteCountry(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&Country{})
	return err
}
func GetJapanCity() (data []JapanCity, err error) {
	err = DB.Table("japan_city").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddJapanCity(data JapanCity) (err error) {
	_, err = DB.Table("japan_city").Insert(&data)
	return err
}
func UpdateJapanCity(data JapanCity) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("japan_city").Where("id=?", data.Id).Update(&data)
	return err
}
func DeleteJapanCity(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&JapanCity{})
	return err
}

func GetUITranslate(classid int) (data []UITranslate, err error) {
	err = DB.Table("translate").Where("classid=?", classid).Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func AddUITranslate(data UITranslate) (err error) {
	_, err = DB.Table("translate").Insert(&data)
	return err
}
func UpdateUITranslate(data UITranslate) (err error) {
	if data.Id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Table("translate").Where("id=?", data.Id).Update(&data)
	return err
}

type Translate struct{}

func DeleteUITranslate(id int) (err error) {
	if id == 0 {
		return errors.New("id can not empty")
	}
	_, err = DB.Where("id=?", id).Delete(&Translate{})
	return err
}

// +----+--------------+------------------+
// | id | classname    | tag              |
// +----+--------------+------------------+
// |  5 | JobType      | jobType          | => job_type
// |  6 | Languages    | languages        | => languages
// |  7 | WorkStyle    | workStyle        | => work_style
// |  8 | Education    | education        | => education
// |  9 | Salary       | salary           | => salary
// | 10 | Industry     | industry         | => industry
// | 11 | Industry     | companyType      | => company_type
// | 12 | Country      | country          | => country
// | 13 | JapaneseCity | japanList        | => japan_city
// | 14 | Jobcategory  | allJobcategories | => job_category
// +----+--------------+------------------+

// //////////////DEEPL//////////////////
type DeepLRequest struct {
	Text        []string `json:"text"`
	TargetLang  string   `json:"target_lang"`
	Format      bool     `json:"preserve_formatting"`
	TagHandling string   `json:"tag_handling"`
}

type DeepLResponse struct {
	Translations []DeepLTranslations `json:"translations"`
}

type DeepLTranslations struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

func TranslateByDeepL(target string, text ...string) (DeepLResponse, error) {
	var source string
	for _, v := range text {
		source += v
	}
	// lang := DetectLanguage(source)

	// if lang == "" {
	// 	return DeepLResponse{}, msg.DETECT_SOURCE_LANGUAGE_ERR
	// }
	// var target string
	// if lang == "JA" {
	// 	target = "EN-US"
	// }
	// if lang == "EN-US" {
	// 	target = "JA"
	// }
	data, err := json.Marshal(&DeepLRequest{
		TargetLang:  target,
		Text:        text,
		Format:      true,
		TagHandling: "xml",
	})
	if err != nil {

	}
	resp, errs := DoTranslateDeepL(data)
	if errs != nil {
		return resp, errs
	}
	return resp, errs
}

func DoTranslateDeepL(data []byte) (DeepLResponse, error) {
	var _DeepLResponse DeepLResponse
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("POST", "https://api-free.deepl.com/v2/translate", strings.NewReader(string(data)))
	req.Header.Add("Authorization", "DeepL-Auth-Key "+DEEPL_FREE_API_KEY)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return _DeepLResponse, errors.New("network error")
	}
	defer resp.Body.Close()
	code := resp.StatusCode

	if code != 200 {
		if code == 429 {
			return _DeepLResponse, errors.New("too many request")
		} else if code == 456 {
			return _DeepLResponse, errors.New("quota exceeded")
		} else if code >= 500 {
			return _DeepLResponse, errors.New("temporary error")
		} else {
			return _DeepLResponse, errors.New("unknown error")
		}
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &_DeepLResponse)
	return _DeepLResponse, nil
}

type TranslateText struct {
	Text string `json:"text" binding:"required"`
}

func TranslateEnglishToJapanese(c *gin.Context) {
	var _TranslateText TranslateText
	err := c.ShouldBindJSON(&_TranslateText)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	data, err := TranslateByDeepL("JA", _TranslateText.Text)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 0, "msg": "ok", "data": data})
}
func TranslateJapaneseToEnglish(c *gin.Context) {
	var _TranslateText TranslateText
	err := c.ShouldBindJSON(&_TranslateText)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	data, err := TranslateByDeepL("EN-US", _TranslateText.Text)
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 0, "msg": "ok", "data": data})
}
