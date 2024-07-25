package repo_test

import (
	"testing"
	"time"
	"uitranslate/domain/notification/topic"
	"uitranslate/infrastructure/notification/repo"
	"xorm.io/xorm/core"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"xorm.io/xorm"
)

func setupMockDB() (*xorm.Engine, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	coreDB := core.FromDB(db)

	engine, err := xorm.NewEngineWithDB("mysql", "mock_dsn", coreDB)
	if err != nil {
		return nil, nil, err
	}
	return engine, mock, nil
}

func TestGetTopicInfoById(t *testing.T) {
	engine, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer engine.Close()

	notificationRepo := repo.NewNotificationRepository(engine)
	topicId := int64(1)
	var topicInfoPO = repo.TopicInfoPO{
		Id:              topicId,
		Title:           "Test Title",
		Description:     "Test Description",
		Status:          1,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
		SubscribeTarget: "Test Target",
		TriggerId:       1,
	}
	rows := sqlmock.NewRows([]string{"id", "title", "description", "status", "create_time", "update_time", "subscribe_target", "trigger_id"}).
		AddRow(topicInfoPO.Id, topicInfoPO.Title, topicInfoPO.Description, topicInfoPO.Status, topicInfoPO.CreateTime, topicInfoPO.UpdateTime, topicInfoPO.SubscribeTarget, topicInfoPO.TriggerId)
	mock.ExpectQuery("^SELECT \\* FROM `nf_topic_info` WHERE `id`=?$").WithArgs(topicId).WillReturnRows(rows)

	result, err := notificationRepo.GetTopicInfoById(topicId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, topicInfoPO.Id, result.ID)
}

func TestListTopicInfo(t *testing.T) {
	engine, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer engine.Close()

	notificationRepo := repo.NewNotificationRepository(engine)
	topicInfoPOs := []repo.TopicInfoPO{
		{
			Id:              1,
			Title:           "Test Title 1",
			Description:     "Test Description 1",
			Status:          1,
			CreateTime:      time.Now(),
			UpdateTime:      time.Now(),
			SubscribeTarget: "Test Target 1",
			TriggerId:       1,
		},
		{
			Id:              2,
			Title:           "Test Title 2",
			Description:     "Test Description 2",
			Status:          2,
			CreateTime:      time.Now(),
			UpdateTime:      time.Now(),
			SubscribeTarget: "Test Target 2",
			TriggerId:       2,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "status", "create_time", "update_time", "subscribe_target", "trigger_id"})
	for _, topicInfoPO := range topicInfoPOs {
		rows.AddRow(topicInfoPO.Id, topicInfoPO.Title, topicInfoPO.Description, topicInfoPO.Status, topicInfoPO.CreateTime, topicInfoPO.UpdateTime, topicInfoPO.SubscribeTarget, topicInfoPO.TriggerId)
	}
	mock.ExpectQuery("^SELECT \\* FROM `nf_topic_info`$").WillReturnRows(rows)

	results, err := notificationRepo.ListTopicInfo()
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, len(topicInfoPOs), len(results))
}

func TestSaveTopicInfo(t *testing.T) {
	engine, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer engine.Close()

	notificationRepo := repo.NewNotificationRepository(engine)
	topicInfo := topic.TopicInfo{
		ID:              1,
		Title:           "Test Title",
		Description:     "Test Description",
		Status:          topic.Status(1),
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
		SubscribeTarget: "Test Target",
		TriggerId:       1,
	}

	mock.ExpectExec("^INSERT INTO `nf_topic_info`").WillReturnResult(sqlmock.NewResult(1, 1))

	err = notificationRepo.SaveTopicInfo(topicInfo)
	assert.NoError(t, err)
}

func TestUpdateTopicInfo(t *testing.T) {
	engine, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer engine.Close()

	notificationRepo := repo.NewNotificationRepository(engine)
	topicInfo := topic.TopicInfo{
		ID:              1,
		Title:           "Updated Title",
		Description:     "Updated Description",
		Status:          topic.Status(1),
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
		SubscribeTarget: "Updated Target",
		TriggerId:       1,
	}

	mock.ExpectExec("^UPDATE `nf_topic_info` SET").WillReturnResult(sqlmock.NewResult(1, 1))

	err = notificationRepo.UpdateTopicInfo(topicInfo)
	assert.NoError(t, err)
}
