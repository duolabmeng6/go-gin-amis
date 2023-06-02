package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"os"
	"text/template"
)

// 遍历 字段名称 字段类型 字段备注
type Column struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Type     string `json:"type"`
	Sortable bool   `json:"sortable,omitempty"`
}

type H map[string]interface{}

// RenderTemplateToFile 渲染模板并输出到文件
func RenderTemplateToFile(tmplFilePath string, data H, outputFilePath string) error {
	//检查文件 outputFilePath 文件是否存在 如果存在 则不要覆盖
	//if _, err := os.Stat(outputFilePath); err == nil {
	//	return errors.New("文件已存在 " + outputFilePath)
	//}
	// 解析模板文件
	tmpl, err := template.ParseFiles(tmplFilePath)
	if err != nil {
		return fmt.Errorf("解析模板文件时发生错误: %v", err)
	}
	// 创建输出文件
	file, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("创建输出文件时发生错误: %v", err)
	}
	defer file.Close()

	// 渲染模板并输出结果
	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("渲染模板时发生错误: %v", err)
	}

	return nil
}

// E常见英文翻译函数
// 比如 users 用户 articles 文章
func E常见英文翻译函数(englishName string) string {
	//读入当前目录下文件 翻译内容.json
	//读取文件内容
	文件内容 := ecore.E读入文件("/Users/chensuilong/Desktop/goproject/go-gin-amis/cmd/ectl/cmd/翻译内容.json")
	var data [][]string
	err := json.Unmarshal(文件内容, &data)
	if err != nil {
		fmt.Println(err)
		return englishName
	}
	var chineseName string

	// 遍历数据
	for _, pair := range data {
		if len(pair) == 2 && pair[0] == englishName {
			chineseName = pair[1]
			return chineseName
		}
	}
	return englishName
}

// E常见字段对应的输入控件
// 比如 bigint input-number , varchat input-text,
func E常见字段对应的输入控件(typeName string) string {
	kv := map[string]string{}
	kv["bigint"] = "input-number"
	kv["varchar"] = "input-text"
	kv["text"] = "textarea"
	kv["datetime"] = "input-datetime"
	kv["timestamp"] = "input-datetime"
	kv["bool"] = "switch"
	// 如果没有的话返回 input-text
	if v, ok := kv[typeName]; ok {
		return v
	}
	return "input-text"
}
