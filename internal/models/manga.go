package models

type Manga struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Genre []string `json:"genre"`
	Cover string   `json:"cover"`
	Price int      `json:"price"`
}
