package cmd

import (
	"GoTour/internal/word"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var str string
var mode int8

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线转为大写驼峰",
	"4: 下划线转为大写驼峰",
	"5: 驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word查看帮助文档")
		}

		log.Printf("输出结果：%s", content)
	},
}

const (
	ModeUpper = 1 + iota
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "输入的单词")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "输入转换模式")
}
