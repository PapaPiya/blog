package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	cfg  Config
	conf = "../conf/conf.toml"
)

// InitCfg 初始化配置文件
func InitConf() error {
	_, err := toml.DecodeFile(conf, &cfg)
	return err
}

// Addr 获取服务器需要的监听地址
func Addr() string {
	return fmt.Sprintf("%s:%d", cfg.Application.Host, cfg.Application.Port)
}

func GetCfgName() string {
	return cfg.Application.Name
}

func GetCfgURL() string {
	return cfg.Application.URL
}

func GetCfgHost() string {
	return cfg.Application.Host
}

func GetCfgPort() uint {
	return cfg.Application.Port
}

func GetCfgDebug() bool {
	return cfg.Application.Debug
}

func GetCfgMarkdownDir() string {
	return cfg.Application.MarkdownDir
}

func GetCfgICP() string {
	return cfg.Application.ICP
}

func GetCfgStatics() string {
	return cfg.Application.Statics
}

func GetCfgSecret() string {
	return cfg.Application.Secret
}

func GetCfgMode() string {
	return cfg.Log.Mode
}

func GetCfgDir() string {
	return cfg.Log.Dir
}

func GetCfgFormat() string {
	return cfg.Log.Format
}

func GetCfgAccess() bool {
	return cfg.Log.Access
}

func GetPathRoot() string {
	return cfg.Path.RootPath
}