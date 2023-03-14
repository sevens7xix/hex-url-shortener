package domain

import "time"

type Data struct {
	ID       uint64    `json:"id"`
	Original string    `json:"original"`
	Short    string    `json:"short"`
	Created  time.Time `json:"created"`
}

func NewData(id uint64, original, short string, created time.Time) Data {
	return Data{
		ID:       id,
		Original: original,
		Short:    short,
		Created:  created,
	}
}
