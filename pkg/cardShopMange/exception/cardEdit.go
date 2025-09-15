package exception

import "fmt"

type CardEditing struct{
	CardID   uint64
}

func (e *CardEditing) Error() string{
	return  fmt.Sprintf("Editing Card id: %d failed ! ",e.CardID )
}

