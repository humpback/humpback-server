package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.NewString()
}

// GetRandomString 获取随机字符串
func GetRandomString() string {
	return strings.ReplaceAll(NewUUID(), "-", "")
}

// GenerateRandomNumber 生成6位随机数
func GenerateRandomNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}
