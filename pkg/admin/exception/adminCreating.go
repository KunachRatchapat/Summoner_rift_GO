package exception

import "fmt"

type AdminCreating struct {
	AdminID string
}

func (e *AdminCreating) Error() string { 
	return fmt.Sprintf("Creating card: %s Failed !!", e.AdminID)
}