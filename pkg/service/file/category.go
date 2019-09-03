package file

import (
	"beegolearn/pkg/common"
	"github.com/BurntSushi/toml"
)

// Model 博客内容的实例
var (
	Model    = InitList()

	categories []Category
)

// Category 文章分类
type Category struct {
	Name       string
	Number      int
	Path        string
	Route       string
	Description string
	Active      bool
	Articles    []Article
}

// List 博客列表
type List interface {
	GetAllArticles() Articles
	GetAllTags() Tags
	GetAllCategories() Categories
	GetArticleContent(string) string
	ArticleByTitle(string) *Article
	ArticlesByCategory(string) Articles
	ArticlesByTag(string) Articles
	Reload()
}

// Initialize 初始化数据
func InitList() List {
	// 暂时是保存到内存，后面计划支持其他存储方式
	l := newListMap()

	return l
}

// 解析分类配置文件
func initCategories() ([]Category,error) {

	temp := struct {
		Category []Category
	}{}

	if _, err := toml.DecodeFile(common.CATEGORY_PATH, &temp); err != nil {
		return nil, err
	}
	categories = temp.Category
	return temp.Category, nil
}

func GetCategories() []Category {
	return categories
}
