package common

const (
	// VERSION 版本信息
	VERSION  = "1.0.0"
	TEMPLATE = "../view/*/*"
	STATIC   = "../static/"
	BLOG     = "blogs"
	CATEGORY_PATH = "../conf/categories.toml"
)

type Config struct {
	Application struct {
		Name        string // 应用的名称
		URL         string // 应用的域名
		Host        string // 应用的监听地址，为空就是 0.0.0.0
		Port        uint   // 应用端口号
		Debug       bool   // 如果 false ， 对应的 gin 就是 release 模式;如果 ture ，对应的 gin 就是 debug 模式，
		MarkdownDir string // markdown 的默认目录
		ICP         string // icp 备案信息
		Statics     string // 百度统计 的 key
		Secret      string // Github 钩子中配置的 Secret
	}
	Log struct {
		Mode   string // 日志模式，对应这的是 LogClose ，LogFile， LogStdout
		Dir    string // 日志文件的保存目录
		Format string // 日志文件的格式
		Access bool   // 是否开启访问日志，一般开启，有时可以关闭，比如部署在 Nginx 后面，// Nginx 已经开启了，可以考虑把这个日志关闭，记录两份有点多了
	}
	Path struct {
		RootPath string
	}
}
