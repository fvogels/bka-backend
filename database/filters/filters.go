package filters

import "github.com/Masterminds/squirrel"

type Filter interface {
	Build() squirrel.Sqlizer
}
