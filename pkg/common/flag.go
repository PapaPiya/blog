package common

// 此处是处理默认参数

import (
	"flag"
	logger "go.uber.org/zap"
	"os"
	"path/filepath"
)

var configDir string

// 初始化配置文件
func init() {
	// 初始化所有命令行传入的参数
	flags()
}

func flags() {
	// 配置文件所在目录
	flag.StringVar(&configDir, "config", "", "Config file directory")
	flag.Parse()
}

// 获取项目的根目录
func getTempPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.L().Error("get temp path",logger.Error(err))
	}
	return path
}
