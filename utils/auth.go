package utils

import (
	"blog-system/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strings"
)

// 对密码进行加密
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// 验证加密后的密码
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func TranslateValidationErrors(err error) string {
	if validErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validErrors {
			filed, _ := reflect.TypeOf(models.User{}).FieldByName(e.Field())
			vmsg := filed.Tag.Get("vmsg")
			//解析自定义错误消息
			if vmsg != "" {
				message := strings.Split(vmsg, ",")
				for _, msg := range message {
					parts := strings.SplitN(msg, ":", 2)
					if len(parts) == 2 && parts[0] == e.Tag() {
						return parts[1]
					}
				}
			}
			//默认错误消息
			switch e.Tag() {
			case "required":
				return e.Field() + "不能为空"
			case "min":
				return e.Field() + "长度至少为" + e.Param()
			case "email":
				return "邮箱格式不正确"
			default:
				return e.Field() + "字段错不合法"
			}

		}
	}
	return err.Error()
}
