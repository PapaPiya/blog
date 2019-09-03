package app

import (
	"beegolearn/pkg/common"
	"beegolearn/pkg/service/router"
	"beegolearn/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// Init 初始化
func Init() *gin.Engine {
	// 初始化日志配置
	setDefaultConfig()

	// 初始化引擎
	engine := gin.New()
	// 初始化模板相关的内容
	loadTemplate(engine)

	// 配置路由
	router.New(engine)

	return engine
}

func setDefaultConfig() {
	if common.GetCfgDebug() {
		gin.SetMode(gin.DebugMode)

		return
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	gin.DefaultWriter = ioutil.Discard
}

func loadTemplate(e *gin.Engine) {
	// 初始化静态文件路径
	e.StaticFS("/public", http.Dir(common.STATIC))
	// 初始化 favicon 图标
	//e.StaticFile("/favicon.ico", root+"/resources/static/favicon.ico")

	// 初始化模板中的自定义函数
	e.FuncMap = funcMap()
	// 初始化模板
	e.LoadHTMLGlob(common.TEMPLATE)

}

func funcMap() map[string]interface{} {
	return map[string]interface{}{
		// markdown 转 html
		"markdowntohtml": util.MarkdownToHTML,
		"staticpath": func(path string) string {
			return fmt.Sprintf("%s:%d/%s", common.GetCfgURL(), common.GetCfgPort(), strings.Trim(path, "/"))
		},
		"tolower": func(str string) string {
			return strings.ToLower(str)
		},
	}
}

// Run 启动服务
func Run(engine *gin.Engine) error {
	return engine.Run(common.Addr())
}
