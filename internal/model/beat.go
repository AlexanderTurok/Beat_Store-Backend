package model

import (
	"errors"
	"time"
)

type Beat struct {
	Id        int64     `json:"id"              db:"id"`
	ProductId int64     `json:"product_id"      db:"product_id"`
	Name      string    `json:"name"            db:"name"        binding:"required"`
	Bpm       int64     `json:"bpm"             db:"bpm"         binding:"required"`
	Key       string    `json:"key"             db:"key"         binding:"required"`
	PhotoPath string    `json:"photo_path"      db:"photo_path"  binding:"required"`
	MP3Path   string    `json:"mp3_path"        db:"mp3_path"    binding:"required"`
	WavPath   string    `json:"wav_path"        db:"wav_path"`
	Genre     string    `json:"genre"           db:"genre"`
	Mood      string    `json:"mood"            db:"mood"`
	Tags      []Tag     `json:"tags"`
	Price     Price     `json:"price"`
	CreatedAt time.Time `json:"created_at"      db:"created_at"`
}

type Tag struct {
	Id   int64  `json:"id"   db:"tag_id"`
	Name string `json:"name" db:"tag_name" binding:"required"`
}

type Price struct {
	Standart int64 `json:"standart"    db:"standart"    binding:"required"`
	Premium  int64 `json:"premium"     db:"premium"     binding:"required"`
	Ultimate int64 `json:"ultimate"    db:"ultimate"    binding:"required"`
}

type BeatUpdateInput struct {
	Name      *string `json:"name"`
	Bpm       *int64  `json:"bpm"`
	Key       *string `json:"key"`
	PhotoPath *string `json:"photo_path"`
	MP3Path   *string `json:"mp3_path"`
	WavPath   *string `json:"wav_path"`
	Genre     *string `json:"genre"`
	Mood      *string `json:"mood"`
}

type TagUpdateInput struct {
	TagName *string `json:"name"`
}

type PriceUpdateInput struct {
	Standart *int64 `json:"standart"`
	Premium  *int64 `json:"premium"`
	Ultimate *int64 `json:"ultimate"`
}

func (b *BeatUpdateInput) Validate() error {
	if b.Name == nil && b.Bpm == nil && b.Key == nil && b.PhotoPath == nil &&
		b.MP3Path == nil && b.WavPath == nil && b.Genre == nil && b.Mood == nil {
		return errors.New("update struct has no values")
	}

	return nil
}
