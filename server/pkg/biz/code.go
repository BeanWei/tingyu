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
	CodeUsernameExisted        = 20006
	CodeUserNotFound           = 20007
)

var (
	codeI18nText = map[int][2]string{
		CodeServerError:            {"内部服务器错误", "server error"},
		CodeParamBindError:         {"参数信息错误", "param bind error"},
		CodeNotAuthorized:          {"未登录或非法访问", "not authorized"},
		CodeForbidden:              {"你没有权限访问该资源", "forbidden"},
		CodeUnauthorizedAuthFailed: {"账户不存在或密码错误", "auth failed"},
		CodeTooManyLoginError:      {"登录失败次数过多，请稍后再试", "too many login error"},
		CodeInvalidPhone:           {"无效的手机号", "invalid phone"},
		CodeInvalidEmail:           {"无效的邮箱", "invalid email"},
		CodeInvalidPassword:        {"密码长度6~16且必须包含数字和大小写字母", "invalid password"},
		CodeUsernameExisted:        {"该手机号或邮箱已经注册", "username existed"},
		CodeUserNotFound:           {"用户不存在", "user not found"},
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
