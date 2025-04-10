package article

type ArticleLimitOffset struct {
	Limit  int `json:"limit" validate:"required,min=1"`
	Offset int `json:"offset" validate:"gte=0"`
}

type ArticleId struct {
	Id int `json:"id" validate:"required,min=0"`
}

type ArticleRequestV2 struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

type ArticleData struct {
	Title    string
	Content  string
	Category string
	Status   string
}
type DataCreated struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Content  string
	Category string
	Status   string
}

type Response struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
