package article

import (
	"fmt"
	"math"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateData(db *gorm.DB, data ArticleData) (int, error)
	GetDataLimitOffset(db *gorm.DB, limit int, offset int, isPreview bool) ([]Response, int, error)
	GetDataById(db *gorm.DB, id int) (ArticleData, error)
	EditData(db *gorm.DB, id int, data ArticleData) (int, error)
	DeleteData(db *gorm.DB, id int) (int, error)
	TrashData(db *gorm.DB, id int) (int, error)
}

type articleRepository struct {
	Log *logrus.Logger
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

func (r articleRepository) CreateData(db *gorm.DB, data ArticleData) (int, error) {
	query := `INSERT INTO posts(title, content, category, status)
          VALUES (?, ?, ?, ?)`

	result := db.Exec(query, data.Title, data.Content, data.Category, data.Status)
	if result.Error != nil {
		return 0, result.Error
	}

	return 1, nil
}

func (r articleRepository) GetDataLimitOffset(db *gorm.DB, limit int, offset int, isPreview bool) ([]Response, int, error) {
	var results []Response
	var totalItems int64
	countSql := "SELECT COUNT(*) FROM posts"

	query := db.Table("posts")

	if isPreview {
		query = query.Where("status = ?", "publish")
		countSql = "SELECT COUNT(*) FROM posts WHERE status = 'publish'"
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&results).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Raw(countSql).Scan(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(totalItems) / float64(limit)))

	return results, totalPage, nil
}

func (r articleRepository) GetDataById(db *gorm.DB, id int) (ArticleData, error) {

	// Query to get all data
	var result ArticleData

	sql := `
		select
			*
		from
			posts
		where 
			id = ?
	`
	if err := db.Raw(sql, id).Scan(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (r articleRepository) EditData(db *gorm.DB, id int, data ArticleData) (int, error) {

	result := db.Exec(`
		UPDATE posts
		SET title = ?, content = ?, category = ?, status = ?
		WHERE id = ?`,
		data.Title, data.Content, data.Category, data.Status, id,
	)

	if result.Error != nil {
		return 0, result.Error
	}

	// Optional: check how many rows were affected
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("no record updated, id %d not found", id)
	}

	return 1, nil
}

func (r articleRepository) DeleteData(db *gorm.DB, id int) (int, error) {
	result := db.Exec(`
		DELETE FROM posts
		WHERE id = ?`, id,
	)

	if result.Error != nil {
		return 0, result.Error
	}

	// Optional: check if any row was actually deleted
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("no record deleted, id %d not found", id)
	}

	return 1, nil
}

func (r articleRepository) TrashData(db *gorm.DB, id int) (int, error) {

	result := db.Exec(`
		UPDATE posts
		SET status = 'trash'
		WHERE id = ?`, id,
	)

	if result.Error != nil {
		return 0, result.Error
	}

	// Optional: check how many rows were affected
	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("no record updated, id %d not found", id)
	}

	return 1, nil
}
