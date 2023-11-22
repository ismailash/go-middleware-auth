package modelutil

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status Status `json:"status"`
	Data   any    `json:"data"`
}

type PagedResponse struct {
	Status Status `json:"status"`
	Data   []any  `json:"data"`
	Paging any    `json:"paging"`
}
