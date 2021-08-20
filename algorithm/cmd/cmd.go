package cmd

import (
	"strconv"
	"strings"
)

const TitleLength = 4

/**
--tag 题目分类
--zh 中文名
--en 英文名
--no 编号
NewSolution --tag tree --zh '二叉树的遍历' --en 'traversal binary tree' --no 64
*/
// NewSolution todo 自动创建题目文件的脚本
func NewSolution(tag string, no int, name string) string {
	noStr := strconv.Itoa(no)
	for i := len(noStr); i < TitleLength; i++ {
		noStr = "0" + noStr
	}

	name = strings.ReplaceAll(name, " ", "-")

	name = noStr + "." + name
	return name
}

func getDirPath() {

}

func getDirName(no int, name string) string {
	noStr := strconv.Itoa(no)
	for i := len(noStr); i < TitleLength; i++ {
		noStr = "0" + noStr
	}

	name = strings.ReplaceAll(name, " ", "-")

	name = noStr + "." + name
	return name
}
