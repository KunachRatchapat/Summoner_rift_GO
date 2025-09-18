package service

import(
	_adminModel 	"github.com/tehdev/summoner-rift-api/pkg/admin/model"
	_playerModel 	"github.com/tehdev/summoner-rift-api/pkg/player/model"
)

type OAuth2Service interface{
	PlayerAccountCreating(playerCreatingRe *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
	IsThisGuyRealPlayer(playerID string) bool
	IsThisGuyRealAdmin(adminID string) bool

}