package exception

type CardCounting struct{}

func (e *CardCounting) Error() string{
	return  "Failed to count Cards"
}