package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
)

type PlaylistService struct {
	repos repository.Playlist
}

func NewPlaylistService(repos repository.Playlist) *PlaylistService {
	return &PlaylistService{
		repos: repos,
	}
}

func (s *PlaylistService) Create(accountId int, input model.Playlist) (int, error) {
	return s.repos.Create(accountId, input)
}

func (s *PlaylistService) Get(playlistId int) (model.Playlist, error) {
	return s.repos.Get(playlistId)
}

func (s *PlaylistService) GetAll() ([]model.Playlist, error) {
	return s.repos.GetAll()
}

func (s *PlaylistService) GetAccountsPlaylist(accountId, playlistId int) (model.Playlist, error) {
	return s.repos.GetAccountsPlaylist(accountId, playlistId)
}

func (s *PlaylistService) GetAllAccountsPlaylists(accountId int) ([]model.Playlist, error) {
	return s.repos.GetAllAccountsPlaylists(accountId)
}

func (s *PlaylistService) Update(playlistId int, input model.PlaylistUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.Update(playlistId, input)
}

func (s *PlaylistService) Delete(playlistId int) error {
	return s.repos.Delete(playlistId)
}

func (s *PlaylistService) AddBeat(playlistId, beatId int) error {
	return s.repos.AddBeat(playlistId, beatId)
}

func (s *PlaylistService) GetBeat(playlistId, beatId int) (model.Beat, error) {
	return s.repos.GetBeat(playlistId, beatId)
}

func (s *PlaylistService) GetAllBeats(playlistId int) ([]model.Beat, error) {
	return s.repos.GetAllBeats(playlistId)
}

func (s *PlaylistService) GetBeatFromAccountsPlaylists(accountId, playlistId, beatId int) (model.Beat, error) {
	return s.repos.GetBeatFromAccountsPlaylists(accountId, playlistId, beatId)
}

func (s *PlaylistService) GetAllBeatsFromAccountsPlaylists(accountId, playlistId int) ([]model.Beat, error) {
	return s.repos.GetAllBeatsFromAccountsPlaylists(accountId, playlistId)
}

func (s *PlaylistService) DeleteBeat(playlistId, beatId int) error {
	return s.repos.DeleteBeat(playlistId, beatId)
}
