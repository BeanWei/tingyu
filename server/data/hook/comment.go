package hook

import (
	"context"

	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/comment"
	"github.com/BeanWei/tingyu/data/ent/hook"
	"github.com/BeanWei/tingyu/g"
)

func init() {
	ent.DB().Comment.Use(
		// 评论创建或删除后更新对应帖子的评论总数
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.CommentFunc(func(ctx context.Context, m *ent.CommentMutation) (ent.Value, error) {
					postId, exists := m.PostID()
					val, err := next.Mutate(ctx, m)
					if err != nil {
						return nil, err
					}
					if exists && postId > 0 {
						g.Pool().Submit(func() {
							total := ent.DB().Comment.Query().Where(
								comment.DeletedAtEQ(0),
								comment.PostIDEQ(postId),
							).CountX(context.Background())
							ent.DB().Post.UpdateOneID(postId).SetCommentCount(total).ExecX(context.Background())
						})
					}
					return val, err
				})
			},
			ent.OpCreate|ent.OpDeleteOne,
		),
	)
}
