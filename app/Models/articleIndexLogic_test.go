package Models

import (
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"testing"
)

type ArticleItems struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type ArticleIndexResponse struct {
	Items []ArticleItems `json:"items"`
	Total int64          `json:"total"`
}

type ArticleIndexRequest struct {
	Page     int64  `form:"page"`
	PerPage  int64  `form:"perPage"`
	Keywords string `form:"keywords,optional"`
	OrderBy  string `form:"orderBy,optional,default=id"`
	OrderDir string `form:"orderDir,optional,default=desc"`
}

func TestIndex(t *testing.T) {
	//ArticleModel := NewArticlesModel(sqlx.NewMysql("root@tcp(127.0.0.1:3310)/gotest?charset=utf8&parseTime=true&loc=Local"))
	//resp, err := ArticleModel.Index(context.Background(), "", 10, 1, "id", "desc")
	////优雅的输出结构体
	//t.Logf("%+v", resp)
	//t.Logf("%+v", err)

	//db, _ := gorm.Open(mysql.Open(config.E数据库连接字符串), &gorm.Config{})
	//stmt := db.Session(&gorm.Session{DryRun: true})
	//stmt.Table("articles").Select("id", "title", "content", "created_at", "updated_at").Find(&ArticleItems{})
	//fmt.Println("sql", stmt.Statement.SQL.String())
	//SELECT id,title,content,created_at,updated_at FROM articles ORDER BY id desc, title desc LIMIT 0,10

	//t.Log(sql)
}
