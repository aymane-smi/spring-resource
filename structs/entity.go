package structs

const (
	Int  = 1
	Long = 2
	Uuid = 3
)

type Entity struct {
	Name   string
	TypeId int
}
