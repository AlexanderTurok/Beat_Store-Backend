package beatstore

type Beat struct {
	Id    int     `json:"id"`
	Bpm   int     `json:"bpm"`
	Key   string  `json:"key"`
	Path  string  `json:"path"`
	Tags  []Tag   `json:"tags"`
	Price float32 `json:"price"`
}

type Tag struct {
	Name string `json:"name"`
}
