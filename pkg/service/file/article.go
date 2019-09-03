package file

import (
	"beegolearn/pkg/common"
	"encoding/json"
	"fmt"
	logger "go.uber.org/zap"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

// Article 文章内容
type Article struct {
	Title        string // 文章的标题
	FileName     string
	UpdatedTime  time.Time // 最后更新时间
	Tags         []string  `json:"tags,string"` // 标签
	Body         string
	CategoryPath string
	Description  string
}

type Articles []Article

var (
	splitFlag = "---"
)

func InitArticle() error {
	temp := struct {
		Article []Article
	}{}

	// 根据目录获取所有文章
	for c := range GetCategories() {
		fmt.Println(c)
	}
	art := append(Articles{}, temp.Article...)
	for _, article := range art {
		stat, err := os.Stat(path.Join(common.GetPathRoot(), common.BLOG, article.CategoryPath, article.FileName))
		if err != nil {
			continue
		}
		article.UpdatedTime = stat.ModTime()
	}
	return nil
}

// 获取指定分类下的文章
func (category *Category) getAritclesByCategory() Articles {

	number := 0

	artDir := path.Join(common.GetPathRoot(), common.BLOG, category.Path)
	files, err := ioutil.ReadDir(artDir)
	if err != nil {
		logger.L().Error("read dir error", logger.Error(err))
		return nil
	}
	if len(files) == 0 {
		logger.L().Error("no articles in this dir")
		return nil
	}
	for f := range files {
		file := files[f]
		if file.IsDir() {
			continue
		}
		bytes, err := ioutil.ReadFile(path.Join(artDir, file.Name()))
		if err != nil {
			logger.L().Error("read file error", logger.Error(err))
			continue
		}
		strs := strings.Split(string(bytes), splitFlag)
		if len(strs) < 2 {
			continue
		}
		var art Article
		if err := json.Unmarshal([]byte(strs[0]), &art); err != nil {
			logger.L().Error("unmarshal head error", logger.Error(err))
			continue
		}
		art.CategoryPath = category.Path
		art.UpdatedTime = file.ModTime()
		art.Body = strings.Join(strs[1:],splitFlag)
		category.Articles = append(category.Articles, art)
		number++
	}

	logger.L().Info("get category's article success", logger.String("category", category.Name), logger.Any("articles", category.Articles))
	category.Number = number
	return category.Articles
}

//下面三个方法是实现了 sort 接口，实现 Articles 的排序
// Len
func (a Articles) Len() int {
	return len(a)
}

// Swap 实现的 sort 接口
func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less 进行排序比较， 第一排序是 UpdatedAt ，
// 第二排序是 CreatedAt， 第三排序是 Title
func (a Articles) Less(i, j int) bool {
	if a[i].UpdatedTime.After(a[j].UpdatedTime) {
		return true
	}

	return false
}
