package utilites

type IdGenerator interface {
	Generate() uint64
}

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate() uint64 {}
