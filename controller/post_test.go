package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"community_id":1,
		"title":"test",
		"content":"just a test"
		}`
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// 判断响应内容
	// 方法1 ： 判断响应内容是否包含指定字符串
	assert.Contains(t, w.Body.String(), "需要登录")

	// 方法2 : 将响应的内容反序列化到ResponseData, 然后判断字段与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}

func TestGetPostDetailHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post/4717746902274048"
	r.GET(url, GetPostDetailHandler)

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// 判断响应内容
	// 方法1 ： 判断响应内容是否包含指定字符串
	assert.Contains(t, w.Body.String(), "请求参数错误")

	// 方法2 : 将响应的内容反序列化到ResponseData, 然后判断字段与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeInvalidParam)
}
