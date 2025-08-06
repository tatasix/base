package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// GenerateSignature 生成API请求签名
// method: HTTP方法(GET, POST等)
// path: 请求路径(/v1/smss/...)
// queryParams: URL查询参数
// body: 请求体JSON转换后的map
// appId: 应用ID
// appSecret: 应用密钥 (客户端保存，不传输)
// 返回: 签名, 时间戳, 随机字符串(nonce)
func GenerateSignature(method, path string, queryParams url.Values, body map[string]interface{}, appId, appSecret string) (signature, timestamp, nonce string) {
	// 生成时间戳和随机字符串
	timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	nonce = generateNonce(16)

	// 构建签名字符串
	signStr := buildSignString(method, path, queryParams, body, timestamp, nonce, appId)

	// 生成签名
	signature = generateHmacSha256(signStr, appSecret)

	return signature, timestamp, nonce
}

// generateNonce 生成指定长度的随机字符串
func generateNonce(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// buildSignString 构建待签名字符串 (不包含appSecret)
func buildSignString(method, path string, query url.Values, body map[string]interface{}, timestamp, nonce, appId string) string {
	// 构建参数表
	params := make(map[string]string)

	// 添加URL查询参数
	for k, v := range query {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}

	// 添加JSON体参数（一级）
	for k, v := range body {
		if str, ok := v.(string); ok {
			params[k] = str
		} else if num, ok := v.(float64); ok {
			params[k] = strconv.FormatFloat(num, 'f', -1, 64)
		} else if b, ok := v.(bool); ok {
			params[k] = strconv.FormatBool(b)
		}
	}

	// 添加时间戳和随机字符串
	params["timestamp"] = timestamp
	params["nonce"] = nonce
	params["appId"] = appId // 添加appId作为签名的一部分

	// 按照键名排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建键值对字符串
	pairs := make([]string, 0, len(keys))
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, params[k]))
	}

	// 拼接最终字符串: METHOD + PATH + QUERY_STRING
	// 注意: 不再添加appSecret到签名字符串中
	signStr := method + path + strings.Join(pairs, "&")
	return signStr
}

// generateHmacSha256 生成HMAC-SHA256签名
func generateHmacSha256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
