package domain

type Data struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

func NewData(original, short string) Data {
	return Data{
		Original: original,
		Short:    short,
	}
}
