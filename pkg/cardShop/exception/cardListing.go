package exception

type CardListing struct{}

//Overide คลาสแม่มันมา
func (e *CardListing) Error() string{
	return "Cards listing Faied"
}
