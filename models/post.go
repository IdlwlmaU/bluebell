package models

import "time"

// 内存对齐

// Post 帖子信息结构体
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情结构体
type ApiPostDetail struct {
	AuthorName       string                    `json:"author_name"`
	VoteNum          int64                     `json:"vote_num"`
	*Post                                      // 嵌入帖子结构体
	*CommunityDetail `json:"community_detail"` // 嵌入社区信息结构体
}