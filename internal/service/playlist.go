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

func (s *PlaylistService) GetAllAccountsPlaylists(accountId int) ([]beatstore.Playlist, error) {
	return s.repos.GetAllAccountsPlaylists(accountId)
}

func (s *PlaylistService) Update(playlistId int, input beatstore.PlaylistUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.Update(playlistId, input)
}
