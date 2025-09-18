package exception

type PlayerCoinShow struct{}

func (e *PlayerCoinShow) Error() string{
	return "Failed to Show Coin Player !!"
}

