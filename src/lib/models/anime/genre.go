package anime

type Genre struct {
	MalID int    `json:"mal_id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type GenreResponse struct {
	Data []Genre `json:"data"`
}