package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 投票

// PostVoteHandler 投票信息接口
// @Summary 投票信息接口
// @Description 投票信息的接口
// @Tags 投票相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer JWT"
// @Param object body models.ParamVoteData false "获取参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /vote [post]
func PostVoteHandler(c *gin.Context) {
	// 1.获取参数及参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := RemoveTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	// 获取当前请求的用户ID
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 2.业务处理
	if err = logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}
