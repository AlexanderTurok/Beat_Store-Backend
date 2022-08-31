package beatstore

import "errors"

type Beat struct {
	Id    int    `json:"id" db:"id"`
	Bpm   string `json:"bpm" db:"bpm"`
	Key   string `json:"key" db:"key"`
	Path  string `json:"path" db:"path" binding:"required"`
	Photo string `json:"photo" db:"photo"`
	Tag   string `json:"tag" db:"tag"`
	Price string `json:"price" db:"price" binding:"required"`
}

type BeatUpdateInput struct {
	Bpm   *string `json:"bpm"`
	Key   *string `json:"key"`
	Path  *string `json:"path"`
	Photo *string `json:"photo"`
	Tag   *string `json:"tag"`
	Price *string `json:"price"`
}

func (b *BeatUpdateInput) Validate() error {
	if b.Bpm == nil && b.Key == nil && b.Photo == nil &&
		b.Path == nil && b.Tag == nil && b.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
