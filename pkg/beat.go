package beatstore

import (
	"errors"
	"time"
)

type Beat struct {
	Id             int       `json:"-"`
	ArtistId       int       `json:"artist_id"`
	Name           string    `json:"name"            binding:"required"`
	Bpm            string    `json:"bpm"             binding:"required"`
	Key            string    `json:"key"             binding:"required"`
	PhotoPath      string    `json:"photo_path"      binding:"required"`
	MP3Path        string    `json:"mp3_path"        binding:"required"`
	WavPath        string    `json:"wav_path"`
	Genre          string    `json:"genre"`
	Mood           string    `json:"mood"`
	Tags           []Tag     `json:"tags"`
	StandartPrice  string    `json:"standart_price"  binding:"required"`
	PremiumPrice   string    `json:"premium_price"   binding:"required"`
	UnlimitedPrice string    `json:"unlimited_price" binding:"required"`
	CreatedAt      time.Time `json:"created_at"      binding:"required"`
}

type Tag struct {
	Id      int    `json:"-"`
	BeatId  int    `json:"beat_id"  binding:"required"`
	TagName string `json:"tag_name" binding:"required"`
}

type BeatUpdateInput struct {
	Name           *string `json:"name"`
	Bpm            *string `json:"bpm"`
	Key            *string `json:"key"`
	PhotoPath      *string `json:"photo_path"`
	MP3Path        *string `json:"mp3_path"`
	WavPath        *string `json:"wav_path"`
	Genre          *string `json:"genre"`
	Mood           *string `json:"mood"`
	Tags           *[]Tag  `json:"tags"`
	StandartPrice  *string `json:"standart_price"`
	PremiumPrice   *string `json:"premium_price"`
	UnlimitedPrice *string `json:"unlimited_price"`
}

func (b *BeatUpdateInput) Validate() error {
	if b.Name == nil && b.Bpm == nil && b.Key == nil && b.PhotoPath == nil &&
		b.MP3Path == nil && b.WavPath == nil && b.Tags == nil && b.Genre == nil &&
		b.Mood == nil && b.StandartPrice == nil && b.PremiumPrice == nil && b.UnlimitedPrice == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
