package Models

import (
	"github.com/jinzhu/gorm"
)

func (Articles) TableName() string {
	return "Articles"
}

type ArticlesService struct {
	DB *gorm.DB
}

func NewArticlesService(db *gorm.DB) *ArticlesService {
	return &ArticlesService{DB: db}
}

// Index 查询文章列表 perPage 每页显示数量 page 页码 keywords 搜索关键字可能为空
func (c *ArticlesService) Index(keywords string, perPage int64, page int64, OrderBy string, OrderDir string) (*[]Articles, int64, error) {
	var articles []Articles
	db := c.DB
	if keywords != "" {
		db = db.Where("title like ?", "%"+keywords+"%")
	}
	// 查询总数量
	var count int64
	result := db.Model(&Articles{}).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	result2 := db.Order(OrderBy + " " + OrderDir).Limit(perPage).Offset((page - 1) * perPage).Find(&articles)

	return &articles, count, result2.Error
}

// Insert
func (c *ArticlesService) Insert(articles *Articles) (int64, error) {
	result := c.DB.Create(&articles)
	return articles.Id, result.Error
}

// FindOne
func (c *ArticlesService) FindOne(id int64) (*Articles, error) {
	var articles Articles
	result := c.DB.Where("id = ?", id).First(&articles)
	return &articles, result.Error
}

// Update
func (c *ArticlesService) Update(articles *Articles) error {
	result := c.DB.Save(&articles)
	return result.Error
}

// Delete
func (c *ArticlesService) Delete(id int64) error {
	result := c.DB.Where("id = ?", id).Delete(&Articles{})
	return result.Error
}
