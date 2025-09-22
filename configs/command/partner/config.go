package config

import (
	"os"
	"strconv"
)

var (
	PORT int
)

func init() {
	// 環境変数PORTを取得（必須）
	portStr := os.Getenv("PORT")
	if portStr == "" {
		panic("環境変数PORTが設定されていません")
	}

	var err error
	PORT, err = strconv.Atoi(portStr)
	if err != nil {
		panic("環境変数PORTが無効な数値です: " + portStr)
	}
}
