package beatstore

import (
	"errors"
	"time"
)

type Beat struct {
	Id        int       `json:"-"          db:"id"`
	ArtistId  int       `json:"-"          db:"artist_id"`
	Name      string    `json:"name"       db:"name"       binding:"required"`
	Bpm       string    `json:"bpm"        db:"bpm"        binding:"required"`
	Key       string    `json:"key"        db:"key"        binding:"required"`
	PhotoPath string    `json:"photo_path" db:"photo_path" binding:"required"`
	MP3Path   string    `json:"mp3_path"   db:"mp3_path"   binding:"required"`
	WavPath   string    `json:"wav_path"   db:"wav_path"`
	Likes     int       `json:"likes"      db:"likes"`
	Genre     string    `json:"genre"      db:"genre"`
	Mood      string    `json:"mood"       db:"mood"`
	Tags      []Tag     `json:"tags"       db:"tags" `
	Price     Price     `json:"price"      db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Tag struct {
	Id      int    `json:"id"       db:"id"`
	BeatId  int    `json:"beat_id"  db:"beat_id"  binding:"required"`
	TagName string `json:"tag_name" db:"tag_name" binding:"required"`
}

type Price struct {
	Id        int    `json:"id"        db:"id"`
	BeatId    int    `json:"beat_id"   db:"beat_id"  binding:"required"`
	Standart  string `json:"standart"  db:"standart" binding:"required"`
	Premium   string `json:"premium"   db:"premium"`
	Unlimited string `json:"unlimited" db:"unlimited"`
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
	Tags      *[]Tag  `json:"tags"`
	Price     *Price  `json:"price"`
}

func (b *BeatUpdateInput) Validate() error {
	if b.Name == nil && b.Bpm == nil && b.Key == nil && b.PhotoPath == nil &&
		b.MP3Path == nil && b.WavPath == nil && b.Tags == nil && b.Genre == nil &&
		b.Mood == nil && b.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
