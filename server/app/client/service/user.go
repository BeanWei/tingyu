package service

import (
	"context"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/BeanWei/tingyu/app/client/types"
	"github.com/BeanWei/tingyu/data/ent"
	"github.com/BeanWei/tingyu/data/ent/user"
	"github.com/BeanWei/tingyu/data/enums"
	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/biz"
	"github.com/BeanWei/tingyu/pkg/cryptor"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/validator"
)

const (
	LOGIN_ERR_KEY       = "user-login-err"
	MAX_LOGIN_ERR_TIMES = 10
)

// UserLoginOrSignIn 登录或注册
func UserLoginOrSignIn(ctx context.Context, req *types.UserLoginReq) (*ent.User, *errors.Error) {
	usr, err := ent.DB().User.Query().Where(user.UsernameEQ(req.Username)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// 如果账号不存在则注册
			if err := ValidateUsername(ctx, req.Username); err != nil {
				return nil, err
			}
			if err := ValidPassword(req.Password); err != nil {
				return nil, err
			}
			salt := random.RandString(10)
			usr := ent.DB().User.Create().
				SetStatus(enums.UserStatusActive).
				SetUsername(req.Username).
				SetNickname(req.Username).
				SetPassword(cryptor.HashUserPwd(req.Password, salt)).
				SetSalt(salt).
				SaveX(ctx)
			return usr, nil
		}
		return nil, errors.New(err, errors.ErrorTypePublic, nil)
	}
	lek := fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, usr.ID)
	// 登录错误次数
	if errTimes, _ := g.Redis().Get(ctx, lek).Int(); errTimes >= MAX_LOGIN_ERR_TIMES {
		return nil, biz.NewError(biz.CodeTooManyLoginError, err)
	}
	if cryptor.HashUserPwd(req.Password, usr.Salt) != usr.Password {
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
	return biz.NewError(biz.CodeInvalidPassword, errors.NewPublic("invalid password"))
}
