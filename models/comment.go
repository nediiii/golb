package models

import "github.com/jinzhu/gorm"

// Comment represent the comment
type Comment struct {
	gorm.Model
	Nickname string     // 评论者昵称!
	Email    string     // 评论者邮箱! 仅对用户可见,保护评论者隐私
	Target   string     // 评论对象! [post|comment]
	Content  string     // 评论内容!
	Domain   string     // 评论者网址
	IP       string     // 评论者IP
	Agent    string     // 评论者客户端
	Status   string     // 评论状态  [公开(默认)|私密(仅作者和评论双方可见)]
	PostID   uint       // 被评论的文章ID!
	ParentID uint       // 父级评论id
	Reply    []*Comment `gorm:"foreignkey:ParentID"` // 评论的回复
}

// IsNode IsNode
func (v *Comment) IsNode() {}
