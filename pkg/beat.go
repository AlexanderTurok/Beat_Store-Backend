package beatstore

import (
	"errors"
	"time"
)

type Beat struct {
	Id        int       `json:"-" db:"id"`
	ArtistId  int       `json:"artist_id" db:"artist_db" binding:"required"`
	Name      string    `json:"name" db:"name" binding:"required"`
	Bpm       string    `json:"bpm" db:"bpm" binding:"required"`
	Key       string    `json:"key" db:"key" binding:"required"`
	PhotoPath string    `json:"photo_path" db:"photo_path" binding:"required"`
	MP3Path   string    `json:"mp3_path" db:"mp3_path" binding:"required"`
	WavPath   string    `json:"wav_path" db:"wav_path"`
	Like      int       `json:"like" db:"like"`
	Genre     string    `json:"genre" db:"genre"`
	Mood      string    `json:"mood" db:"mood"`
	Tags      []Tag     `json:"tags" db:"tags" binding:"required"`
	Price     Price     `json:"price" db:"price" binding:"required"`
	CreatedAt time.Time `json:"craeted_at" db:"created_at" binding:"required"`
}

type Tag struct {
	Id     int    `json:"-" db:"id"`
	BeatId int    `json:"beat_id" db:"beat_id"`
	Name   string `json:"name" db:"name"`
}

type Price struct {
	Id             int    `json:"-"`
	StandartPrice  string `json:"standart" db:"standart"`
	PremiumPrice   string `json:"premium" db:"premium"`
	UnlimitedPrice string `json:"unlimited" db:"unlimited"`
}

type BeatUpdateInput struct {
	Name      *string    `json:"name"`
	Bpm       *string    `json:"bpm"`
	Key       *string    `json:"key"`
	PhotoPath *string    `json:"photo_path"`
	MP3Path   *string    `json:"mp3_path"`
	WavPath   *string    `json:"wav_path"`
	Genre     *string    `json:"genre"`
	Mood      *string    `json:"mood"`
	Tags      *[]Tag     `json:"tags"`
	Price     *Price     `json:"price"`
	CreatedAt *time.Time `json:"craeted_at"`
}

func (b *BeatUpdateInput) Validate() error {
	if b.Name == nil && b.Bpm == nil && b.Key == nil && b.PhotoPath == nil &&
		b.MP3Path == nil && b.WavPath == nil && b.Tags == nil && b.Price == nil &&
		b.Genre == nil && b.Mood == nil && b.CreatedAt == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
