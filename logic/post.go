package logic

import (
	"go.uber.org/zap"
	"know_or/dao/mysql"
	"know_or/dao/redis"
	"know_or/model"
	"know_or/pkg/snowflake"
	"strconv"
)

const (
	// USER_SUPPORT_NONE 代表用户没投票
	USER_SUPPORT_NONE = 0
)

// CreatePost 用来创建一条帖子
func CreatePost(p *model.Post) (err error) {
	// 查询发帖的社区id是否存在
	if _, err = mysql.GetCommunityDetailByID(p.CommunityID); err != nil {
		return err
	}
	// 生成帖子id
	p.ID = snowflake.GenID()
	// 插入数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	return redis.CreatePost(p.ID, p.CommunityID)
}

// GetPostByID 返回对应id的帖子详情
func GetPostByID(pid int64) (*model.ApiPostDetail, error) {
	// 查询帖子详情
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return nil, err
	}

	// 查询帖子投票数
	postId := strconv.Itoa(int(post.ID))
	supportData, err := redis.GetPostSupportData([]string{postId})
	if err != nil {
		return nil, err
	}

	// 查询作者
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return nil, err
	}

	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed",
			zap.Int64("author_id", post.CommunityID),
			zap.Error(err))
		return nil, err
	}

	// 返回数据
	data := &model.ApiPostDetail{
		AuthorName:      user.Username,
		CommunityDetail: community,
		Supports:        supportData[0],
		Post:            post,
	}
	return data, err
}

func GetPostByIDWithUid(pid, uid int64) (*model.ApiPostDetail, error) {
	data, err := GetPostByID(pid)
	if err != nil {
		return nil, err
	}

	// 查询 uid 的投票情况
	status := getPostSupportStatus(pid, uid)
	data.SupportStatus = status

	return data, nil
}

// GetPostList 获取帖子列表
//func GetPostList(page, pageSize int64) ([]*model.ApiPostDetail, error) {
//	posts, err := mysql.GetPostList(page, pageSize)
//	if err != nil {
//		return nil, err
//	}
//
//	data := make([]*model.ApiPostDetail, 0, len(posts))
//	for _, post := range posts {
//		// 查询作者
//		user, err := mysql.GetUserByID(post.AuthorID)
//		if err != nil {
//			zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
//			return nil, err
//		}
//
//		// 根据社区id查询社区详细信息
//		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
//		if err != nil {
//			zap.L().Error("mysql.GetPostByID(pid) failed",
//				zap.Int64("author_id", post.CommunityID),
//				zap.Error(err))
//			return nil, err
//		}
//		postDetails := &model.ApiPostDetail{
//			AuthorName:      user.Username,
//			Post:            post,
//			CommunityDetail: community,
//		}
//		data = append(data, postDetails)
//	}
//	return data, nil
//}

// GetPostList2 获取帖子列表（新版）
func GetPostList2(p *model.ParamPostList) ([]*model.ApiPostDetail, error) {
	// 去 redis 查询 id 列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return nil, nil
	}
	// 根据 id 取数据库查询帖子详情信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return nil, err
	}
	//zap.L().Debug("GetPostList2", zap.Any("posts", posts))

	// 查询每篇帖子的投票数
	supportData, err := redis.GetPostSupportData(ids)
	if err != nil {
		return nil, err
	}

	// 查询更详细的信息
	data := make([]*model.ApiPostDetail, 0, len(posts))
	for i, post := range posts {
		// 查询作者
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			return nil, err
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetPostByID(pid) failed",
				zap.Int64("author_id", post.CommunityID),
				zap.Error(err))
			return nil, err
		}
		postDetails := &model.ApiPostDetail{
			AuthorName:      user.Username,
			Supports:        supportData[i],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetails)
	}
	return data, nil
}

// GetCommunityPostList 按社区查询帖子，返回查询到的帖子
func GetCommunityPostList(p *model.ParamPostList) ([]*model.ApiPostDetail, error) {
	// 去 redis 查询 id 列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return nil, nil
	}
	// 根据 id 取数据库查询帖子详情信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return nil, err
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))

	// 提前查询好每篇帖子的投票数
	supportData, err := redis.GetPostSupportData(ids)
	if err != nil {
		return nil, err
	}

	// 查询更详细的信息
	data := make([]*model.ApiPostDetail, 0, len(posts))
	for i, post := range posts {
		// 查询作者
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			return nil, err
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetPostByID(pid) failed",
				zap.Int64("author_id", post.CommunityID),
				zap.Error(err))
			return nil, err
		}
		postDetails := &model.ApiPostDetail{
			AuthorName:      user.Username,
			Supports:        supportData[i],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetails)
	}
	return data, nil
}

// GetPostListNew 根据参数 CommunityID 判断是否需要查询指定社区的帖子
func GetPostListNew(p *model.ParamPostList) (data []*model.ApiPostDetail, err error) {
	if p.CommunityID <= 0 {
		// 查所有社区的
		data, err = GetPostList2(p)
	} else {
		// 查某个社区的
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
	return
}

// GetPostListWithUid 根据参数 CommunityID 判断是否需要查询指定社区的帖子
// 并查询当前用户对这些帖子的投票情况
func GetPostListWithUid(p *model.ParamPostList, uid int64) (data []*model.ApiPostDetail, err error) {
	data, err = GetPostListNew(p)
	if err != nil {
		zap.L().Error("GetPostListWithUser failed", zap.Error(err))
		return nil, err
	}

	// 对帖子加上用户的投票情况
	for i := range data {
		status := getPostSupportStatus(data[i].Post.ID, uid)
		data[i].SupportStatus = status
	}
	return
}

// getPostSupportStatus 获取用户对于某个帖子投票的情况
func getPostSupportStatus(pid int64, uid int64) int {
	res, err := redis.GetPostUserSupportStatus(pid, uid)
	if err != nil {
		return USER_SUPPORT_NONE
	}
	return int(res)
}
