package beatstore

type Beat struct {
	Id    int     `json:"id" db:"id"`
	Bpm   int     `json:"bpm" db:"bpm"`
	Key   string  `json:"key" db:"key"`
	Path  string  `json:"path" db:"path"`
	Tags  []Tag   `json:"tags" db:"tags"`
	Price float32 `json:"price" db:"price"`
}

type Tag struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
