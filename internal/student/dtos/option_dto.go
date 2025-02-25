package dtos

type Option struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OptionDTO struct {
	Genders   []*Option `json:"genders"`
	Faculties []*Option `json:"faculties"`
	Courses   []*Option `json:"courses"`
	Programs  []*Option `json:"programs"`
	Statuses  []*Option `json:"statuses"`
}
