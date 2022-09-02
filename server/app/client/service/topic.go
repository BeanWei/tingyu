package service

import (
	"context"

	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/topic"
	"github.com/BeanWei/tingyu/data/ent/user"
)

func TopicIsFollowed(ctx context.Context, topicId, userId int64) bool {
	return ent.DB().User.Query().Where(
		user.IDEQ(userId),
		user.HasTopicsWith(topic.IDEQ(topicId)),
	).ExistX(ctx)
}
