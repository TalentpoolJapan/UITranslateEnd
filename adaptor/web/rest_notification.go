package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uitranslate/app/notification"
)

var (
	TopicAppServ     notification.TopicAppServ
	TriggerAppServ   notification.TriggerAppServ
	SubscribeAppServ notification.SubscribeAppServ
)

func RegisterNotificationHandler(engine *gin.Engine) {
	//engine.GET("/admin/notification/topic/page")
	engine.GET("/admin/notification/topic/info/list", ListTopic)
	engine.POST("/admin/notification/topic/info", AddTopic)
	engine.PUT("/admin/notification/topic/info", UpdateTopic)

	engine.GET("/admin/notification/topic/template/list", ListTopicTemplateByTopicId)
	engine.POST("/admin/notification/topic/template", AddTopicTemplate)
	engine.PUT("/admin/notification/topic/template", UpdateTopicTemplate)

	engine.GET("/admin/notification/trigger/list", ListTriggers)
	engine.POST("/admin/notification/trigger", AddTrigger)
	engine.PUT("/admin/notification/trigger", UpdateTrigger)

	// subscribe
	engine.POST("/api/notification/subscribe", SubscribeTopic)
}

func SubscribeTopic(context *gin.Context) {
	var cmd notification.SubscribeTopicCmd
	if err := context.BindJSON(&cmd); err != nil {
		context.JSON(http.StatusBadRequest, NewApiRestResult(RestResult{Code: -1, Message: err.Error()}))
		return
	}
	if err := SubscribeAppServ.SubscribeTopic(cmd); err != nil {
		context.JSON(http.StatusInternalServerError, NewApiRestResult(RestResult{Code: -1, Message: err.Error()}))
		return
	}
	context.JSON(http.StatusOK, NewApiRestResult(RestResult{Code: 0, Message: "Subscribe successfully"}))
}

func ListTopic(c *gin.Context) {
	topicInfoResps, bizErr := TopicAppServ.ListTopicInfo()
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "", Data: topicInfoResps})
}

func AddTopic(c *gin.Context) {
	var cmd notification.TopicInfoAddCmd
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: err.Error()})
		return
	}
	if err := TopicAppServ.AddTopicInfo(cmd); err != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Topic added successfully"})
}

func UpdateTopic(c *gin.Context) {
	var cmd notification.TopicInfoUpdateCmd
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: err.Error()})
		return
	}
	if err := TopicAppServ.UpdateTopicInfo(cmd); err != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Topic updated successfully"})
}

func ListTopicTemplateByTopicId(c *gin.Context) {
	var qry notification.TopicTemplateByTopicIdQuery
	if err := c.ShouldBindQuery(&qry); err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: err.Error()})
		return
	}
	topicTemplateResps, bizErr := TopicAppServ.ListTopicTemplateByTopicId(qry)
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "", Data: topicTemplateResps})
}

func AddTopicTemplate(c *gin.Context) {
	var cmd notification.TopicTemplateAddCmd
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: err.Error()})
		return
	}
	if err := TopicAppServ.AddTopicTemplate(cmd); err != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Topic template added successfully"})
}

func UpdateTopicTemplate(c *gin.Context) {
	var cmd notification.TopicTemplateUpdateCmd
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: err.Error()})
		return
	}
	if err := TopicAppServ.UpdateTopicTemplate(cmd); err != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Topic template updated successfully"})
}

func ListTriggers(c *gin.Context) {
	triggerResps, bizErr := TriggerAppServ.ListTrigger()
	if bizErr != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: bizErr.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "", Data: triggerResps})
}

func AddTrigger(c *gin.Context) {
	var cmd notification.TriggerAddCmd
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: err.Error()})
		return
	}
	if err := TriggerAppServ.AddTrigger(&cmd); err != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Trigger added successfully"})
}

func UpdateTrigger(c *gin.Context) {
	var cmd notification.TriggerUpdateCmd
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, RestResult{Code: -1, Message: err.Error()})
		return
	}
	if err := TriggerAppServ.UpdateTrigger(&cmd); err != nil {
		c.JSON(http.StatusInternalServerError, RestResult{Code: -1, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, RestResult{Code: 0, Message: "Trigger update successfully"})
}
