package dto

type Results struct {
	Next    string   `json:"next"`
	Results []Planet `json:"results"`
}

type Planet struct {
	Name    string   `json:"name"`
	Climate string   `json:"climate"`
	Terrain string   `json:"terrain"`
	Films   []string `json:"films"`
}
