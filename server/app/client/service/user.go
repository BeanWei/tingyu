package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/user"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/duke-git/lancet/v2/validator"
	"golang.org/x/crypto/pbkdf2"
)

const (
	LOGIN_ERR_KEY       = "userloginerr"
	MAX_LOGIN_ERR_TIMES = 10
)

// GetLoginUser 用户登录认证
func GetLoginUser(ctx context.Context, req *types.UserLoginReq) (*ent.User, *errors.Error) {
	usr, err := ent.DB().User.Query().Where(user.UsernameEQ(req.Username)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, biz.NewError(biz.CodeUnauthorizedAuthFailed, err)
	}
	lek := fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, usr.ID)
	// 登录错误次数
	if errTimes, _ := g.Redis().Get(ctx, lek).Int(); errTimes >= MAX_LOGIN_ERR_TIMES {
		return nil, biz.NewError(biz.CodeTooManyLoginError, err)
	}
	if HashUserPwd(req.Password, usr.Salt) != usr.Password {
		// 更新登录错误次数
		if times := g.Redis().Incr(ctx, lek).Val(); times == 1 {
			g.Redis().Expire(ctx, lek, time.Hour)
		}
		return nil, biz.NewError(biz.CodeUnauthorizedAuthFailed, err)
	}
	// 清空登录错误次数
	g.Redis().Del(ctx, lek)
	return usr, nil
}

// ValidateUsername 校验用户名是否合法
func ValidateUsername(ctx context.Context, username string) *errors.Error {
	if strings.Contains(username, "@") {
		if !validator.IsEmail(username) {
			return biz.NewError(biz.CodeInvalidEmail, fmt.Errorf("email %s is invalid", username))
		}
	} else if !validator.IsChineseMobile(username) {
		return biz.NewError(biz.CodeInvalidEmail, fmt.Errorf("phone %s is invalid", username))
	}
	if ent.DB().User.Query().Where(user.UsernameEqualFold(username)).ExistX(ctx) {
		return biz.NewError(biz.CodeUsernameExisted, fmt.Errorf("username %s is existed", username))
	}
	return nil
}

// ValidPassword 密码合法性验证: 必须包含字母大小写和数字且长度符合要求
func ValidPassword(password string) *errors.Error {
	if len(password) >= 6 || len(password) <= 16 {
		var num, lower, upper bool
		for _, r := range password {
			switch {
			case unicode.IsDigit(r):
				num = true
			case unicode.IsLower(r):
				lower = true
			case unicode.IsUpper(r):
				upper = true
			}
		}
		if num && lower && upper {
			return nil
		}
	}
	return biz.NewError(biz.CodeInvalidPassword, fmt.Errorf("password %s is invalid", password))
}

// HashUserPwd 用户明文密码加密
func HashUserPwd(password, salt string) string {
	passwd := pbkdf2.Key([]byte(password), []byte(salt), 10000, 50, sha256.New)
	return fmt.Sprintf("%x", passwd)
}
