package redis

// redis key

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   // zest;帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  // zset;帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" // zet;记录用户及投票类型;参数是post_id
	KeyCommunitySetPrefix  = "community:"  // set;保存每个分区下帖子id
)

// getRedisKey 给redis函数加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
