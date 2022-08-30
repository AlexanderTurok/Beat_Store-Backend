package beatstore

type Beat struct {
	Id    int     `json:"id" db:"id"`
	Bpm   string  `json:"bpm" db:"bpm"`
	Key   string  `json:"key" db:"key"`
	Path  string  `json:"path" db:"path"`
	Tag   string  `json:"tag" db:"tag"`
	Price float32 `json:"price" db:"price"`
}
