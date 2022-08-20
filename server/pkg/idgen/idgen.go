package idgen

import "github.com/yitter/idgenerator-go/idgen"

func init() {
	idgen.SetIdGenerator(idgen.NewIdGeneratorOptions(1))
}

func NewID() int64 {
	return idgen.NextId()
}
