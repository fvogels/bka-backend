package filters

import "github.com/Masterminds/squirrel"

func And(filters ...Filter) Filter {
	return conjunction{children: filters}
}

type conjunction struct {
	children []Filter
}

func (filter conjunction) Build() squirrel.Sqlizer {
	operands := make([]squirrel.Sqlizer, len(filter.children))

	for index, filter := range filter.children {
		operands[index] = filter.Build()
	}

	return squirrel.And(operands)
}
