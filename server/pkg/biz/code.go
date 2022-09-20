package biz

import "github.com/cloudwego/hertz/pkg/protocol/consts"

const (
	// 5xx 10000~10100
	CodeServerError = 10000
	// 400 11001~11999
	CodeParamBindError = 11001
	// 401 12001~12999
	CodeNotAuthorized = 12001
	// 403 13001~13999
	CodeForbidden = 13001
	// 404 14001~14999
	// 400 BizError 20001~99999
	CodeUnauthorizedAuthFailed = 20001
	CodeTooManyLoginError      = 20002
	CodeInvalidPhone           = 20003
	CodeInvalidEmail           = 20004
	CodeInvalidPassword        = 20005
	CodeInvalidNickname        = 20006
	CodeUsernameExisted        = 20007
	CodeNicknameExisted        = 20008
	CodeUserNotFound           = 20009
	CodeInvalidTopicTitle      = 20010
	CodeTopicIsFollowed        = 20011
	CodeTopicIsNotFollowed     = 20012
	CodeFileTypeUnSupport      = 20013
	CodeFileSizeTooLarge       = 20014
)

var (
	codeI18nText = map[int][2]string{
		CodeServerError:            {"内部服务器错误", "server error"},
		CodeParamBindError:         {"参数信息错误", "param bind error"},
		CodeNotAuthorized:          {"未登录或非法访问", "not authorized"},
		CodeForbidden:              {"你没有权限访问该资源", "forbidden"},
		CodeUnauthorizedAuthFailed: {"密码错误", "auth failed"},
		CodeTooManyLoginError:      {"登录失败次数过多，请稍后再试", "too many login error"},
		CodeInvalidPhone:           {"无效的手机号", "invalid phone"},
		CodeInvalidEmail:           {"无效的邮箱", "invalid email"},
		CodeInvalidPassword:        {"密码长度6~16且必须包含数字和大小写字母", "invalid password"},
		CodeInvalidNickname:        {"昵称长度3-12且不能包含特殊字符", "invalid nickname"},
		CodeUsernameExisted:        {"该手机号或邮箱已经注册", "username existed"},
		CodeNicknameExisted:        {"该昵称已经被使用", "nickname existed"},
		CodeUserNotFound:           {"用户不存在", "user not found"},
		CodeInvalidTopicTitle:      {"标题长度不能超过20", "invalid topic title"},
		CodeTopicIsFollowed:        {"你已经关注该话题", "topic is followed"},
		CodeTopicIsNotFollowed:     {"你还没关注该话题", "topic is not followed"},
		CodeFileTypeUnSupport:      {"不支持上传此类型文件", "not support upload file type"},
		CodeFileSizeTooLarge:       {"文件太大", "file size too large"},
	}
)

func CodeText(bizCode int) string {
	return codeI18nText[bizCode][0]
}

func Code2HttpCode(bizCode int) int {
	if bizCode == 0 {
		return consts.StatusOK
	}
	if bizCode > 20000 {
		return consts.StatusBadRequest
	}
	if 11000 < bizCode && bizCode < 12000 {
		return consts.StatusBadRequest
	}
	if 12000 < bizCode && bizCode < 13000 {
		return consts.StatusUnauthorized
	}
	if 13000 < bizCode && bizCode < 14000 {
		return consts.StatusForbidden
	}
	if 14000 < bizCode && bizCode < 15000 {
		return consts.StatusNotFound
	}
	return consts.StatusInternalServerError
}
