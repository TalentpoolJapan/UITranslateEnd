package main

import (
	"fmt"
	"net/http"

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
	MYSQL_DB = "talentpool"

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
		authorized.GET("/get/category", GetTranslateCategory)
		// 修改分类列表
		authorized.POST("/update/category", UpdateTranslateCategory)
		// 获取分类下面的翻译内容
		authorized.GET("/get/category/classid/:id", GetTranslateByClassid)

		authorized.GET("/get/jobcategoryclass", GetJobCategoryClass)
		authorized.GET("/get/jobcategorysubclass/:id", GetJobCategorySubClass)
		authorized.POST("/update/jobcategoryclass", EditJobCategoryClass)
		authorized.POST("/update/jobcategorysubclass", EditJobCategorySubClass)
		authorized.POST("/add/jobcategoryclass", AddJobCategoryClass)
		authorized.POST("/add/jobcategorysubclass", AddJobCategorySubClass)
		authorized.POST("/delete/jobcategoryclass", DeleteJobCategoryClass)
		authorized.POST("/delete/jobcategorysubclass", DeleteJobCategorySubClass)

	}

	r.Run(":8332")

}

// 获取分类
type TranslateClass struct {
	Id        int    `json:"id" binding:"required"`
	Classname string `json:"classname" binding:"required"`
	Tag       string `json:"tag" binding:"required"`
}

func GetTranslateCategory(c *gin.Context) {
	var _TranslateClass []TranslateClass
	err := DB.Where("id>=? and id<?", 5, 110).Find(&_TranslateClass)
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
	_, err = DB.Table("translate_class").ID(_TranslateClass.Id).Update(_TranslateClass)
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
	Id         int    `json:"id" binding:"required"`
	IndustryEn string `json:"english" binding:"required"`
	IndustryJa string `json:"japanese" binding:"required"`
}
type CompanyType struct {
	Id            int    `json:"id" binding:"required"`
	CompanyTypeEn string `json:"english" binding:"required"`
	CompanyTypeJa string `json:"japanese" binding:"required"`
}
type Country struct {
	Id        int    `json:"id" binding:"required"`
	CountryEn string `json:"english" binding:"required"`
	CountryJa string `json:"japanese" binding:"required"`
}
type Education struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type JobType struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type Languages struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type WorkStyle struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type Salary struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type WorkExp struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}
type JapanCity struct {
	Id     int    `json:"id" binding:"required"`
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
}

// /////////////////这部分对应网页界面语言////////////////////////
type UITranslate struct {
	Id       int    `json:"id" binding:"required"`
	Transkey string `json:"transkey" binding:"required"`
	English  string `json:"english" binding:"required"`
	Japanese string `json:"japanese" binding:"required"`
	Classid  int    `json:"classid" binding:"required"`
}

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
	Parentid int `uri:"parentid"`
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
	NameEn string `json:"english" binding:"required"`
	NameJa string `json:"japanese" binding:"required"`
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
	_, err = DB.Table("job_category").ID(_UpdateJobCategoryClass.Id).Where("parentid=0").Update(&_UpdateJobCategoryClass)
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
	_, err = DB.Table("job_category").ID(_UpdateJobCategorySubClass.Id).Where("parentid=?", _UpdateJobCategorySubClass.Parentid).And("parentid!=0").Update(&_UpdateJobCategorySubClass)
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
	_, err = DB.Table("job_category").ID(_RemoveJobCategoryClass.Id).Where("parentid=0").Delete(&JobCategory{})
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
	_, err = DB.Table("job_category").ID(_RemoveJobCategorySubClass.Id).Where("parentid!=0").Delete(&JobCategory{})
	if err != nil {
		c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 0, "msg": "ok"})
}

////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////
func GetTranslateByClassid(c *gin.Context) {
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
	//if classid == 14 {
	// data, err := GetJobCategory()
	// if err != nil {
	// 	c.JSON(200, gin.H{"status": 1, "msg": err.Error()})
	// 	return
	// }
	// c.JSON(200, gin.H{"status": 0, "msg": "", "data": data})
	// return
	//}
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

func GetJobType() (data []JobType, err error) {
	err = DB.Table("job_type").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetLanguages() (data []Languages, err error) {
	err = DB.Table("languages").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetWorkStyle() (data []WorkStyle, err error) {
	err = DB.Table("work_style").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetEducation() (data []Education, err error) {
	err = DB.Table("education").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetSalary() (data []Salary, err error) {
	err = DB.Table("salary").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetIndustry() (data []Industry, err error) {
	err = DB.Table("industry").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetCompanyType() (data []CompanyType, err error) {
	err = DB.Table("company_type").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetCountry() (data []Country, err error) {
	err = DB.Table("country").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func GetJapanCity() (data []JapanCity, err error) {
	err = DB.Table("japan_city").Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//	func GetJobCategory() (data []JobCategory, err error) {
//		err = DB.Table("job_category").Find(&data)
//		if err != nil {
//			return data, err
//		}
//		return data, nil
//	}
func GetUITranslate(id int) (data []UITranslate, err error) {
	err = DB.Table("translate").Where("classid=?", id).Find(&data)
	if err != nil {
		return data, err
	}
	return data, nil
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
