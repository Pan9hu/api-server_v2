package service

import (
	"encoding/json"
	"errors"
	"github.com/Pan9Hu/api-server_v2/util"
)

var (
	success             = "10000" //Token鉴权通过返回码
	refreshCode         = "15000" //accessToken 过期刷新token的返回码，后续在前端调用 refresh api
	refreshTokenExpired = "20000" //refreshToken过期, 重新登录返回码
	tokenErrorCode      = "30000" //token出现无法预估的错误返回码（用户名不一致，损坏的token或非法token等）
	JWTUtil             util.JWTUtil
	tokenError          = errors.New("no login or unauthorized access")
	audienceError       = errors.New("aud wrong or unauthorized access")
)

type JWTService struct {
}

func (js *JWTService) TokenIsValid(xToken string) (string, error) {
	tokenMap := make(map[string]string)
	err := json.Unmarshal([]byte(xToken), &tokenMap)
	if err != nil {
		return tokenErrorCode, err
	}

	//Token格式错误, 尚未登录或为认证的访问。
	if len(tokenMap) != 3 {
		return tokenErrorCode, tokenError //Token格式错误, 尚未登录或为认证的访问。
	}

	// 对X-Auth-Token鉴权
	code, err := tokenHandle(tokenMap, "refresh")
	if err != nil {
		return code, err
	}

	code, err = tokenHandle(tokenMap, "access")
	if err != nil {
		return code, err
	}

	return code, err
}

func tokenHandle(tokenMap map[string]string, tokenType string) (string, error) {
	aud, err := JWTUtil.ParseToken(tokenMap[tokenType])
	if err != nil {
		if errors.Is(err, util.TokenExpired) {
			if tokenType == "refresh" {
				return refreshTokenExpired, err
			} else {
				return refreshCode, err
			}
		}
		return tokenErrorCode, err
	}
	if aud != tokenMap["username"] {
		return tokenErrorCode, audienceError
	}
	return success, nil
}
