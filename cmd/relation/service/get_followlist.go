package service

import (
	"ByteTech-7815/douyin-zhgg/dal/db"
	"ByteTech-7815/douyin-zhgg/dal/pack"
	"ByteTech-7815/douyin-zhgg/kitex_gen/relation"
	"ByteTech-7815/douyin-zhgg/kitex_gen/user"
	"ByteTech-7815/douyin-zhgg/pkg/errno"
	"context"
)

type GetFollowListService struct {
	ctx context.Context
}

// NewFollowListService new GetFollowListService
func NewFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{ctx: ctx}
}

func (s *GetFollowListService) GetFollowList(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	// TODO: 这里我删掉了JWT鉴权，不清楚是不是应该去掉 感觉在API就可以

	userIds := []int64{req.UserId}
	users, err := db.QueryUserById(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	u := users[0]
	userList, err := db.GetFollowingUsers(s.ctx, int64(u.ID))
	if err != nil {
		return nil, err
	}
	return pack.FollowList(userList), nil
}
