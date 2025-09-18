package exception

import "fmt"

type CardNotFound struct {
	cardID uint64
}

func (e *CardNotFound) Error() string {
	return fmt.Sprintf("CardID: %d was not found",e.cardID)
}