package domain

type Data struct {
	Body string `json:"text" form:"text" query:"text"`
}
