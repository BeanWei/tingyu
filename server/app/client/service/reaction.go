package service

import (
	"context"

	"github.com/BeanWei/tingyu/app/client/dto"
	"github.com/BeanWei/tingyu/g"
	"github.com/jmoiron/sqlx"
)

func GetReactionsForOneSubject(ctx context.Context, userId, subjectId int64) ([]*dto.Reaction, error) {
	reactions := make([]*dto.Reaction, 0)
	if err := g.RDB().SelectContext(ctx, &reactions, `
		select t1.subject_id, t1.react_code as code, count(t1.react_code) as count,
		case when (
			select count(t2.id) from user_reactions as t2 where t2.user_id = $1 and t2.subject_id = t1.subject_id and t2.react_code = t1.react_code
		) > 0 then true else false end as active
		from user_reactions as t1
		where t1.subject_id = $2
		group by t1.subject_id, t1.react_code
	`, userId, subjectId); err != nil {
		return nil, err
	}
	return reactions, nil
}

func GetReactionsForManySubject(ctx context.Context, userId int64, subjectIds []int64) (map[int64][]*dto.Reaction, error) {
	if len(subjectIds) == 0 {
		return make(map[int64][]*dto.Reaction), nil
	}

	reactions := make([]*dto.Reaction, 0)
	query, args, err := sqlx.In(`
		select t1.subject_id, t1.react_code as code, count(t1.react_code) as count,
		case when (
			select count(t2.id) from user_reactions as t2 where t2.user_id = ? and t2.subject_id = t1.subject_id and t2.react_code = t1.react_code
		) > 0 then true else false end as active
		from user_reactions as t1
		where t1.subject_id in (?)
		group by t1.subject_id, t1.react_code
	`, userId, subjectIds)
	if err != nil {
		return nil, err
	}
	if err := g.RDB().SelectContext(ctx, &reactions, g.RDB().Rebind(query), args...); err != nil {
		return nil, err
	}
	res := make(map[int64][]*dto.Reaction)
	for _, reaction := range reactions {
		if len(res[reaction.SubjectId]) == 0 {
			res[reaction.SubjectId] = make([]*dto.Reaction, 0)
		}
		res[reaction.SubjectId] = append(res[reaction.SubjectId], reaction)
	}
	return res, nil
}
