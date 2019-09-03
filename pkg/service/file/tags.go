package file

// Tags 标签
type Tags []Tag

// Tag 标签
type Tag struct {
	Title  string // 标签名称
	Number int    // 标签下文章的数量
	Active bool   // 是否是选中的
}

// Len 实现 Sort 的接口
func (tags Tags) Len() int {
	return len(tags)
}

// Swap 实现的 Sort 接口
func (tags Tags) Swap(i, j int) {
	tags[i], tags[j] = tags[j], tags[i]
}

// Less 实现的 Sort 接口， 按照标签数量排序
func (tags Tags) Less(i, j int) bool {
	return tags[i].Number > tags[j].Number
}
