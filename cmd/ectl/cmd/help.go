package cmd

import (
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
	kv := map[string]string{
		"users":       "用户",
		"roles":       "角色",
		"permissions": "权限",
		"menus":       "菜单",
		"logs":        "日志",
		"settings":    "设置",
		"articles":    "文章",
		"created_at":  "创建时间",
		"updated_at":  "更新时间",
		"deleted_at":  "删除时间",
		"username":    "用户名",
		"password":    "密码",
		"integral":    "积分",
		"content":     "内容",
		"title":       "标题",
		"remarks":     "备注",
	}
	if v, ok := kv[englishName]; ok {
		return v
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
