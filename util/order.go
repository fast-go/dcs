package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOrderNum() string {
	// 生成当前时间的年月日
	today := time.Now().Format("20060102")

	// 生成4位随机数
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(9000) + 1000

	// 组合成订单号
	orderNo := fmt.Sprintf("%s%d", today, randomNum)

	return orderNo
}
