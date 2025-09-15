package exception

import "fmt"

type CardArchving struct{
	CardID   uint64
}

func (e *CardArchving) Error() string{
	fmt.Sprintf("Archving card id: %d failed", e.CardID)
}