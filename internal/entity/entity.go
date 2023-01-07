package entity

import "strconv"

type Recipient struct {
	ID int
}

func (user Recipient) Recipient() string {
	return strconv.Itoa(user.ID)
}
