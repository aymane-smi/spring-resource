package structs

import "strings"

const (
	Int  = 1
	Long = 2
	Uuid = 3
)

type Entity struct {
	Name      string
	NameLower string
	TypeId    int
	RepoType  string
}

func (e *Entity) SetNameLower() {
	e.NameLower = strings.ToLower(e.Name)
}
