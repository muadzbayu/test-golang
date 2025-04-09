package article

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateData(db *gorm.DB, data ArticleData) (int, error)
	GetDataLimitOffset(db *gorm.DB, limit int, offset int) ([]Response, error)
	GetDataById(db *gorm.DB, id int) (ArticleData, error)
	EditData(db *gorm.DB, id int, data ArticleData) (int, error)
	DeleteData(db *gorm.DB, id int) (int, error)
}

type articleRepository struct {
	Log *logrus.Logger
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

func (r articleRepository) CreateData(db *gorm.DB, data ArticleData) (int, error) {
	var id int

	sql := `
		INSERT INTO posts(title, content, category, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	if err := db.Raw(sql, data.Title, data.Content, data.Category, data.Status).Scan(&id).Error; err != nil {
		return id, err
	}

	return id, nil
}

func (r articleRepository) GetDataLimitOffset(db *gorm.DB, limit int, offset int) ([]Response, error) {

	var results []Response

	sql := `
		select
			title, content, category, status
		from
			posts
		limit $1
		offset $2
	`

	if err := db.Raw(sql, limit, offset).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
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
