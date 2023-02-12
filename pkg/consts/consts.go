package consts

import "time"

const (
	MySQLDefaultDSN     = "zhgg:zhgg@tcp(127.0.0.1:33069)/zhgg-dy?charset=utf8mb4&parseTime=true&loc=Local"
	UserServiceAddr     = ":9000" //User服务地址
	RelationAddr        = ":10000"
	PublishServiceAddr  = ":10010"
	CommentServiceAddr  = ":10020"
	FavoriteServiceAddr = ":10030"
	ETCDAddress         = "127.0.0.1:2379"

	OssEndPoint        = "oss-cn-hangzhou.aliyuncs.com"
	OssAccessKeyId     = "oss"
	OssAccessKeySecret = "oss"
	OssBucketName      = "douyin-zhgg"
	UserTableName      = "user"
	VideoTableName     = "video"
	CommentTableName   = "comment"
	FavoriteTableName  = "favorite"
	RelationTableName  = "relation"

	LimitVideoNum = 30
	FrameNum      = 3

	// jwt
	SecretKey           = "secret key"
	IdentityKey         = "id"
	TokenExpireDuration = time.Hour * 24

	// service name
	ApiServiceName      = "api"
	UserServiceName     = "user"
	RelationServiceName = "relation"
	PublishServiceName  = "publish"
	FeedServiceName     = "feed"

	//时间字段格式
	TimeFormat = "2023-02-06 16:25:05"

	TCP = "tcp"
)
