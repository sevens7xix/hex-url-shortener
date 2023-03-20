package domain

import "time"

type Data struct {
	Original string    `json:"original"`
	Short    string    `json:"short"`
	Created  time.Time `json:"created"`
}

func NewData(original, short string, created time.Time) Data {
	return Data{
		Original: original,
		Short:    short,
		Created:  created,
	}
}
