package midderware

import (
	file2 "beegolearn/pkg/service/file"
	"github.com/gin-gonic/gin"
	"path"
	"sort"
	"strings"
)

/*// Sites 站点的一些公共信息配置
func Sites(c *gin.Context) {
	c.Set("cfg", common.Cfg)

	c.Next()
}*/

// Navigation 导航栏用到的中间件
func Navigation(c *gin.Context) {
	categories := file2.Model.GetAllCategories()
	cates := make(file2.Categories, 0)
	pathValue := c.Param("name")
	for _, category := range categories {
		if  category.Number > 0 {
			// 如果当前参数路径和分类路径相同，就是激活状态
			if strings.ToLower(pathValue) == strings.ToLower(category.Route) ||
				c.Request.URL.Path == category.Route {
				category.Active = true
			}
			if category.Route != "/" {
				category.Route = path.Join("/categories", category.Route)
			}

			cates = append(cates, category)
		}
	}

	// 设置分类数据
	c.Set("categories", cates)
	c.Next()
}

// Tags 是右侧标签中用到的数据
func Tags(c *gin.Context) {

	tags := file2.Model.GetAllTags()
	name := c.Param("name")

	// 按照标签数量排序
	sort.Sort(&tags)

	for i := 0; i < len(tags); i++ {
		if strings.ToLower(name) == tags[i].Title {
			tags[i].Active = true
		} else {
			tags[i].Active = false
		}
	}

	c.Set("tags", tags)
	c.Next()
}
