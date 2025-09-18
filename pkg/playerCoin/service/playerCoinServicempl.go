package service

import (
	"github.com/tehdev/summoner-rift-api/entities"
	_playerCoinModel "github.com/tehdev/summoner-rift-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/tehdev/summoner-rift-api/pkg/playerCoin/repository"
)

type playerCoinServicempl struct{
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository 
}

func NewPlayerCoinServicempl(playerCoinRepository _playerCoinRepository.PlayerCoinRepository,) PlayerCoinService{
	return  &playerCoinServicempl{playerCoinRepository}
}

func (s *playerCoinServicempl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin,error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount: coinAddingReq.Amount,
	}

	playerCoinEntityResult, err := s.playerCoinRepository.CoinAdding(playerCoinEntity)
	if err != nil{
		return nil, err
	}

	return  playerCoinEntityResult.ToPlayerCoinModel(), nil

}

