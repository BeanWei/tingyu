package enums

const (
	UserStatusActive   int8 = iota + 1 // 正常
	UserStatusDisabled                 // 禁用
)

const (
	PostVisibilityPublic     int8 = iota + 1 // 公开
	PostVisibilityPrivate                    // 仅自己可见
	PostVisibilityOnlyFriend                 // 仅好友可见
)

const (
	TopicStatusOnline   int8 = iota + 1 // 上线
	TopicStatusOffline                  // 下线
	TopicStatusAuditing                 // 审核中
)

const (
	PostStatusPass     int8 = iota + 1 //审核通过
	PostStatusFail                     // 审核未通过
	PostStatusAuditing                 // 审核中
)

const (
	CommentStatusPass     int8 = iota + 1 //审核通过
	CommentStatusFail                     // 审核未通过
	CommentStatusAuditing                 // 审核中
)

const (
	CommentReplyStatusPass     int8 = iota + 1 //审核通过
	CommentReplyStatusFail                     // 审核未通过
	CommentReplyStatusAuditing                 // 审核中
)
