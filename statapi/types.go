package statapi

import "html/template"

type Stat struct {
	Type, Picture template.HTML
	Value         string
}
