package util

import (
    "html/template"
)

// MarkdownToHTML 将 markdown 文档转换成 template.HTML
func MarkdownToHTML(input string) template.HTML {
	return HTML(input)
}
