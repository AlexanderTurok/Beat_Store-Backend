package model

import (
	"errors"
	"time"
)

type Beat struct {
	Id        int       `json:"id"              db:"id"`
	ArtistId  int       `json:"artist_id"       db:"artist_id"`
	Name      string    `json:"name"            db:"name"        binding:"required"`
	Bpm       string    `json:"bpm"             db:"bpm"         binding:"required"`
	Key       string    `json:"key"             db:"key"         binding:"required"`
	PhotoPath string    `json:"photo_path"      db:"photo_path"  binding:"required"`
	MP3Path   string    `json:"mp3_path"        db:"mp3_path"    binding:"required"`
	WavPath   string    `json:"wav_path"        db:"wav_path"`
	Genre     string    `json:"genre"           db:"genre"`
	Mood      string    `json:"mood"            db:"mood"`
	Tags      []Tag     `json:"tags"                             binding:"required"`
	Price     Price     `json:"price"`
	CreatedAt time.Time `json:"created_at"      db:"created_at"`
}

type Tag struct {
	Id   int    `json:"id"   db:"tag_id"`
	Name string `json:"name" db:"tag_name"   binding:"required"`
}

type Price struct {
	Standart   float32 `json:"standart"    db:"standart"    binding:"required"`
	Premium    float32 `json:"premium"     db:"premium"     binding:"required"`
	Ultimate   float32 `json:"ultimate"    db:"ultimate"    binding:"required"`
	StandartId string  `json:"standart_id" db:"standart_id" binding:"required"`
	PremiumId  string  `json:"premium_id"  db:"premium_id"  binding:"required"`
	UltimateId string  `json:"ultimate_id" db:"ultimate_id" binding:"required"`
}

type BeatUpdateInput struct {
	Name      *string `json:"name"`
	Bpm       *string `json:"bpm"`
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
	Standart   *float32 `json:"standart"`
	Premium    *float32 `json:"premium"`
	Ultimate   *float32 `json:"ultimate"`
	StandartId *string  `json:"standart_id"`
	PremiumId  *string  `json:"premium_id"`
	UltimateId *string  `json:"ultimate_id"`
}

func (b *BeatUpdateInput) Validate() error {
	if b.Name == nil && b.Bpm == nil && b.Key == nil && b.PhotoPath == nil &&
		b.MP3Path == nil && b.WavPath == nil && b.Genre == nil && b.Mood == nil {
		return errors.New("update struct has no values")
	}

	return nil
}
