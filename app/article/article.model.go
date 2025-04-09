package article

type ArticleReponse struct {
	Terminal_code          string `json:"terminal_code"`
	Partner_name           string `json:"partner_name"`
	Merchant_name          string `json:"merchant_name"`
	Merchant_terminal_name string `json:"merchant_terminal_name"`
	Issuer                 string `json:"issuer"`
}

type ArticleLimitOffset struct {
	Limit  int `json:"limit" validate:"required,min=1"`
	Offset int `json:"offset" validate:"required,min=0"`
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

type SamSetting struct {
	MID     string `json:"mid"`
	TID     string `json:"tid"`
	SAMID   uint   `json:"sam_id"`
	SAMPIN  string `json:"sam_pin"`
	SAMUID  string `json:"sam_uid"`
	SAMSlot int    `json:"sam_slot"`
}

type Issuer struct {
	IssuerID   uint         `json:"issuer_id"`
	IssuerName string       `json:"issuer_name"`
	SamSetting []SamSetting `json:"sam_setting"`
}

type Response struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type GetVersionModel struct {
	Version_app string `json:"version_app"`
	Url_app     string `json:"url_app"`
}
