package beatstore

type Beat struct {
	Id    int     `json:"id"`
	Bpm   int     `json:"bpm"`
	Key   string  `json:"key"`
	Path  string  `json:"path"`
	Tags  string  `json:"tags"`
	Price float32 `json:"price"`
}
