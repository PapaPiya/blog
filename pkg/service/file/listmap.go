package file

import (
	logger "go.uber.org/zap"
	"sort"
	"strings"
)

// ListMap 将博客保存在 map 中，
type ListMap struct {
	Articles   Articles
	Categories Categories
	Tags       Tags
}
type Categories []Category

// Reload 重新加载文档
func (list *ListMap) Reload() {
	Model = InitList()
}

func (list *ListMap) GetAllArticles() Articles {
	return list.Articles
}
func (list *ListMap) GetAllTags() Tags {
	return list.Tags
}
func (list *ListMap) GetAllCategories() Categories {
	return list.Categories
}

func (list *ListMap) GetArticleContent(title string) string {
	for _, article := range list.Articles {
		if article.Title == title {
			str := string(article.Body)
			if str == "" {
				logger.L().Error("head empty or body empty")
			}
			return str
		}
	}

	return ""
}

// ArticleByPath 根据文章的 name 查询指定的文章
func (list *ListMap) ArticleByTitle(title string) *Article {
	if title == "" {
		logger.L().Error("title is empty")
		return nil
	}

	for _, article := range list.Articles {
		if article.Title == title {
			return &article
		}
	}
	logger.L().Error("article not exist", logger.String("Title", title))
	return nil
}

// ArticlesByCategory 根据分类获取博客列表
func (list *ListMap) ArticlesByCategory(category string) Articles {
	for _, cat := range list.Categories {
		if strings.ToLower(cat.Route) == strings.ToLower(category) {
			return cat.Articles
		}
	}
	return nil
}

// ArticlesByTag 根据标签获取博客列表
func (list *ListMap) ArticlesByTag(tag string) Articles {
	articles := make([]Article, 0)
	for _, article := range list.Articles {
		for _, tag := range article.Tags {
			if strings.ToLower(tag) == strings.ToLower(tag) {
				articles = append(articles, article)
			}
		}
	}

	return articles
}

func newListMap() *ListMap {

	categories, err := initCategories()
	if err != nil {
		logger.L().Error("init categories failure", logger.Error(err))
		return nil
	}

	list := ListMap{
		Categories: categories,
		Tags:       make(Tags, 0),
	}
	logger.L().Info("init categories success", logger.Any("Categories", list.Categories))
	list.initArticles()
	logger.L().Info("init articles success", logger.Any("Articles", list.Articles))
	list.initTags()
	logger.L().Info("init tags success", logger.Any("Tags", list.Tags))
	return &list
}

// 初始化文章
func (list *ListMap) initArticles() {
	articles := make(Articles, 0)

	for i := 0; i < len(list.Categories); i++ {
		a := (&list.Categories[i]).getAritclesByCategory()
		mergeArticles := make(Articles, len(articles)+len(a))
		copy(mergeArticles, articles)
		copy(mergeArticles[len(articles):], a)
		articles = mergeArticles
	}

	// 接收到所有文章，并按照 UpdateAt 倒序保存
	sort.Sort(&articles)
	list.Articles = articles
}

// 初始化 tags， 顺便做下排序
func (list *ListMap) initTags() {
	if list.Articles == nil {
		logger.L().Error("articles is empty")
		return
	}
	tags := make(map[string]int)
	for _, article := range list.Articles {
		for _, title := range article.Tags {
			if _, ok := tags[strings.ToLower(title)]; !ok {
				tags[strings.ToLower(title)] = 1
			} else {
				tags[strings.ToLower(title)]++
			}
		}
	}
	for k, v := range tags {
		tag := Tag{
			Title:  k,
			Number: v,
		}
		list.Tags = append(list.Tags, tag)
	}

}
