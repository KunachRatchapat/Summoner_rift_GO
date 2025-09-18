package service

import (
	entities "github.com/tehdev/summoner-rift-api/entities"
	_adminModel "github.com/tehdev/summoner-rift-api/pkg/admin/model"
	_adminRepository "github.com/tehdev/summoner-rift-api/pkg/admin/repository"
	_playerModel "github.com/tehdev/summoner-rift-api/pkg/player/model"
	_playerRepository "github.com/tehdev/summoner-rift-api/pkg/player/repository"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository: playerRepository,
		adminRepository:  adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {
	if !s.IsThisGuyRealPlayer(playerCreatingReq.ID) {
		playerEntity := &entities.Player{
			ID:     playerCreatingReq.ID,
			Email:  playerCreatingReq.Email,
			Name:   playerCreatingReq.Name,
			Avatar: playerCreatingReq.Avatar,
		}

		if _, err := s.playerRepository.Creating(playerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingInfoReq *_adminModel.AdminCreatingReq) error {
	if !s.IsThisGuyRealAdmin(adminCreatingInfoReq.ID) {
		adminEntity := &entities.Admin{
			ID:     adminCreatingInfoReq.ID,
			Email:  adminCreatingInfoReq.Email,
			Name:   adminCreatingInfoReq.Name,
			Avatar: adminCreatingInfoReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}
	}

	return nil

}

func (s *googleOAuth2Service) IsThisGuyRealPlayer(palyerId string) bool {
	player, err := s.playerRepository.FindByID(palyerId)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) IsThisGuyRealAdmin(adminId string) bool {
	admin, err := s.adminRepository.FindByID(adminId)
	if err != nil {
		return false
	}

	return admin != nil
}