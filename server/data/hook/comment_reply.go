package hook

import (
	"context"

	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/commentreply"
	"github.com/BeanWei/tingyu/data/ent/hook"
	"github.com/BeanWei/tingyu/g"
)

func init() {
	ent.DB().CommentReply.Use(
		// 回复创建或删除后更新对应评论的回复总数
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.CommentReplyFunc(func(ctx context.Context, m *ent.CommentReplyMutation) (ent.Value, error) {
					commentId, exists := m.CommentID()
					val, err := next.Mutate(ctx, m)
					if err != nil {
						return nil, err
					}
					if exists && commentId > 0 {
						g.Pool().Submit(func() {
							total := ent.DB().CommentReply.Query().Where(
								commentreply.DeletedAtEQ(0),
								commentreply.CommentIDEQ(commentId),
							).CountX(context.Background())
							ent.DB().Comment.UpdateOneID(commentId).SetReplyCount(total).ExecX(context.Background())
						})
					}
					return val, err
				})
			},
			ent.OpCreate|ent.OpDeleteOne,
		),
	)
}
