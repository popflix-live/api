package tmdb

type Backdrop struct {
	FilePath string `json:"file_path"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type ImagesResponse struct {
	Backdrops []Backdrop `json:"backdrops"`
}