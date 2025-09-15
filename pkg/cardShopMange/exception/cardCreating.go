package exception

type CardCreating struct{}

func (e *CardCreating) Error() string{
	return "Card is Create Failed !"
}