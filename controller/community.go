package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ----- 社区相关 -----

// CommunityHandler 社区信息接口
// @Summary 社区信息接口
// @Description 获取所有社区的接口
// @Tags 社区相关接口
// @Produce application/json
// @Param Authorization header string false "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区(community_id, community_name)以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
// @Summary 社区分类详情的接口
// @Description 获取指定id的社区的详细信息的接口
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer JWT"
// @Param id path string false "id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /community/{id} [get]
func CommunityDetailHandler(c *gin.Context) {
	// 1.获取社区ID
	idStr := c.Param("id") // 获取路径参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get community detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2.根据ID获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
