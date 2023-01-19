// Package zctoken 凭证处理包
package zctoken

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/zhaochuninhefei/gmgo/sm2"
	"gitee.com/zhaochuninhefei/gmgo/sm3"
	"strings"
)

//goland:noinspection GoSnakeCaseUsage
const (
	// ALG_DEFAULT 默认凭证算法
	ALG_DEFAULT = "SM2-with-SM3"
	// TYP_DEFAULT 默认凭证类型
	TYP_DEFAULT = "JWT"
)

// TokenHeader 凭证头部
type TokenHeader struct {
	// Alg 凭证算法
	Alg string `json:"alg"`
	// Typ 凭证类型
	Typ string `json:"typ"`
}

// CreateTokenHeader 创建凭证头部
//  @param alg 凭证算法
//  @param typ 凭证类型
//  @return *TokenHeader
func CreateTokenHeader(alg string, typ string) *TokenHeader {
	return &TokenHeader{
		Alg: alg,
		Typ: typ,
	}
}

// CreateTokenHeaderDefault 使用默认配置创建凭证头部
//  @return *TokenHeader
func CreateTokenHeaderDefault() *TokenHeader {
	return CreateTokenHeader(ALG_DEFAULT, TYP_DEFAULT)
}

// BuildTokenWithGM 使用SM2-with-SM3算法创建凭证
//  @param payloads 凭证有效负载
//  @param priKey 签名私钥(sm2)
//  @return string 凭证字符串
//  @return error
func BuildTokenWithGM(payloads map[string]string, priKey *sm2.PrivateKey) (string, error) {
	// 创建默认token头部
	tokenHeader := CreateTokenHeaderDefault()
	// 将token头转为json
	jsonTokenHeader, err := json.Marshal(&tokenHeader)
	if err != nil {
		return "", err
	}
	// 对token头做base64编码
	headerBase64 := base64.URLEncoding.EncodeToString(jsonTokenHeader)
	// 将token的有效负载转为json
	jsonPayloads, err := json.Marshal(payloads)
	if err != nil {
		return "", err
	}
	// 对token的有效负载做base64编码
	payloadsBase64 := base64.URLEncoding.EncodeToString(jsonPayloads)
	// 拼接token内容
	content := headerBase64 + "." + payloadsBase64
	// 对token内容做sm3摘要计算
	digest := sm3.Sm3Sum([]byte(content))
	// 对摘要做sm2签名
	sign, err := priKey.Sign(rand.Reader, digest, nil)
	if err != nil {
		return "", err
	}
	// 将签名转为hex字符串
	signStr := hex.EncodeToString(sign)
	// 拼接凭证
	token := fmt.Sprintf("%s.%s", content, signStr)
	return token, nil
}

func CheckTokenWithGM(token string, pubKey *sm2.PublicKey) (map[string]string, error) {
	tmpArr := strings.Split(token, ".")
	if len(tmpArr) != 3 {
		return nil, errors.New("-5:token格式错误")
	}
	headerBase64 := tmpArr[0]
	payloadsBase64 := tmpArr[1]
	signStr := tmpArr[2]

	// 检查token头
	jsonTokenHeader, err := base64.URLEncoding.DecodeString(headerBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头base64解码失败: %s", err)
	}
	var tokenHeader TokenHeader
	err = json.Unmarshal(jsonTokenHeader, &tokenHeader)
	if err != nil {
		return nil, fmt.Errorf("[-5]token头json反序列化失败: %s", err)
	}

	// 检查签名
	content := headerBase64 + "." + payloadsBase64
	digest := sm3.Sm3Sum([]byte(content))
	sign, err := hex.DecodeString(signStr)
	if err != nil {
		return nil, errors.New("-5:token签名验证失败")
	}
	if !pubKey.Verify(digest, sign) {
		return nil, errors.New("-5:token签名验证失败")
	}

	// 解析有效负载
	jsonPayloads, err := base64.URLEncoding.DecodeString(payloadsBase64)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载base64解码失败: %s", err)
	}
	var payloads map[string]string
	err = json.Unmarshal(jsonPayloads, &payloads)
	if err != nil {
		return nil, fmt.Errorf("[-5]token有效负载json反序列化失败: %s", err)
	}
	// 凭证过期检查
	exp := payloads["exp"]
	if exp != "" {

	}

	return nil, nil
}
