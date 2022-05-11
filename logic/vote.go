package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

// 本项目使用简化的投票分数
// 投一票就加432分    86400 / 200  -->  200 张票可以续一天

/* 投票的几种情况:
direction=1 时，有两种情况：
	1.之前没有投票，现在投赞成票
	2.之前投反对票，现在改投赞成票

direction=0 时，有两种情况：
	1.之前投赞成票，现在取消投票
	2.之前投反对票，现在取消投票

direction=-1 时，有两种情况：
	1.之前没有投票，现在投反对票
	2.之前投赞成票，现在改投反对票

投票的限制:
每个帖子自发表之日起一个星期内允许用户投票，超过一个星期就不允许投票了
	1.到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2.到期之后删除 KeyPostVotedZSetPrefix
*/

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	//调用redis.VoteForPost()
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
