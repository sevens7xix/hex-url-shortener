package domain

type Data struct {
	Original string `dynamodbav:"original"`
	Short    string `dynamodbav:"short"`
}

func NewData(original, short string) Data {
	return Data{
		Original: original,
		Short:    short,
	}
}
