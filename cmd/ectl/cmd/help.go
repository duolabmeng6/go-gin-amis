package cmd

import (
	"fmt"
	"os"
	"text/template"
)

type H map[string]interface{}

// RenderTemplateToFile 渲染模板并输出到文件
func RenderTemplateToFile(tmplFilePath string, data H, outputFilePath string) error {
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
