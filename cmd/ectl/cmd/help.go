package cmd

import (
	"errors"
	"fmt"
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
	if _, err := os.Stat(outputFilePath); err == nil {
		return errors.New("文件已存在 " + outputFilePath)
	}
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
