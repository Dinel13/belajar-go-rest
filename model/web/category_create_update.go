package web

type CategoryCreateUpdate struct {
	Id   int    `validate : "required" json:"id"`
	Name string `json :"name", validate : "required, max=200, min=1"  json:"name"`
}
