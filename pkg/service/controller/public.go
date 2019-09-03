package controller

import (
	"beegolearn/pkg/common"
	"beegolearn/pkg/service/file"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	logger "go.uber.org/zap"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

// Home 主页控制器
func Home(c *gin.Context) {

	params := mergeH(c, gin.H{
		"title":    "博客列表",
		"keywords": "博客列表",
		"posts":    file.Model.GetAllArticles(),
	})

	c.HTML(200, "posts/index.html", params)
}

// About 关于控制器
func About(c *gin.Context) {
	// about 页面就直接展示的项目根目录下的 README.md
	path := filepath.Join(common.GetPathRoot(),common.BLOG,"README.md")
	about, err := ioutil.ReadFile(path)

	if err != nil {
		c.Redirect(307, "/errors")
		return
	}

	params := mergeH(c, gin.H{
		"about": string(about),
	})

	c.HTML(200, "layouts/about.html", params)
}

// MergeH 合并默认参数
func mergeH(c *gin.Context, h gin.H) gin.H {
	if c.Keys == nil {
		return h
	}

	if h == nil || len(h) == 0 {
		return c.Keys
	}

	mh := make(gin.H)

	for key, val := range c.Keys {
		mh[key] = val
	}

	for key, val := range h {
		mh[key] = val
	}

	return mh
}

// ToKeywords 用都好分割，拼接关键词
func ToKeywords(works ...string) string {
	return strings.Join(works, ",")
}

// Webhook github 钩子
func Webhook(c *gin.Context) {
	singn := c.GetHeader("X-Hub-Signature")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.L().Error("github hook", logger.Error(err))
		return
	}

	if !checkSecret(singn, body) {
		logger.L().Error("github hook check secret failure.")
		return
	}

	pullDocs()

	// 从新初始化博客列表的内容
	file.Model.Reload()
}

// 向 github 发起 pull 请求，更新文档
func pullDocs() {
	// 获取文档保存的路径
	docPath := filepath.Join(common.GetCfgMarkdownDir())
	// 执行 git pull
	cmd := exec.Command("git", "pull")
	// 切换到命令要执行的目录
	cmd.Dir = docPath

	// 执行，并返回结果
	res, err := cmd.Output()

	if err == nil {
		logger.L().Info("git pull success.", logger.ByteString("result", res))
	} else {
		logger.L().Info("git pull failure.", logger.Error(err))
	}

}

// 检测 github 传过来的 key
func checkSecret(singn string, body []byte) bool {
	if len(singn) != 45 || !strings.HasPrefix(singn, "sha1=") {
		return false
	}

	// github 中对应的加密串， 从配置文件去获取
	secret := []byte(common.GetCfgSecret())

	// 通过加密串和 body 计算签名
	mac := hmac.New(sha1.New, secret)
	mac.Write(body)
	mKey := mac.Sum(nil)

	// Hex 解码
	singnature := make([]byte, 20)
	hex.Decode(singnature, []byte(singn[5:]))

	// 比较签名是否一直
	if hmac.Equal(singnature, mKey) {
		return true
	}

	return false
}
