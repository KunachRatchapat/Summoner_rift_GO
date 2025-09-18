package exception

type AddingCoin struct{}

func (e *AddingCoin) Error() string{
	return "Failed to Add Coin !!"
}