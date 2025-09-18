package service

import (
	_adminRepository "github.com/tehdev/summoner-rift-api/pkg/admin/repository"
	_playerRepository "github.com/tehdev/summoner-rift-api/pkg/player/repository"
)

type googleOAuth2Service struct{
	_playerRepository  _playerRepository.PlayerRepository
	adminRepository   _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(
	playerRepository  _playerRepository.PlayerRepository,
	adminRepository   _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository,
		adminRepository,
	}
}

