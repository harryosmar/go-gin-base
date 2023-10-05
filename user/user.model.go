package user

type Provinces struct {
	Id   int64  `json:"id" uri:"id"`
	Name string `json:"name"`
}

type Paging struct {
	Page  int `uri:"page" `
	Limit int `uri:"limit"`
}
