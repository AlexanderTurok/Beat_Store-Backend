package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type PlaylistService struct {
	repos repository.Playlist
}

func NewPlaylistService(repos repository.Playlist) *PlaylistService {
	return &PlaylistService{
		repos: repos,
	}
}

func (s *PlaylistService) Create(accountId int, input beatstore.Playlist) (int, error) {
	return s.repos.Create(accountId, input)
}
