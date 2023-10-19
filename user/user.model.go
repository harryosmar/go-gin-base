package user

type Provinces struct {
	Id   int64  `json:"id" uri:"id"`
	Name string `json:"name"`
}

type Regencies struct {
	Id          int64  `json:"id" uri:"id"`
	ProvincesID int64  `json:"provinces_id"`
	Name        string `json:"name"`
}

type Districts struct {
	Id          int64  `json:"id" uri:"id"`
	RegenciesID int64  `json:"regencies_id"`
	Name        string `json:"name"`
}

type Villages struct {
	Id          int64  `json:"id" uri:"id"`
	DistrictsID int64  `json:"districts_id"`
	Name        string `json:"name"`
}

type Paging struct {
	Page  int `uri:"page"`
	Limit int `uri:"limit"`
}
