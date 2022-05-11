package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
	"time"
)

func init() {
	dbCfg := &settings.MySQLConfig{
		Host:         "127.0.0.1",
		Port:         3306,
		User:         "root",
		Password:     "123456",
		DbName:       "bluebell",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := &models.Post{
		ID:          123,
		AuthorID:    123,
		CommunityID: 1,
		Status:      1,
		Title:       "test",
		Content:     "just a test",
		CreateTime:  time.Time{},
	}
	err := CreatePost(post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
