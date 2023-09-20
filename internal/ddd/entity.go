package ddd

type IDer interface {
	ID() string
}

type EntityNamer interface {
	EntityName() string
}

type Entity struct {
	id   string
	name string
}

var _ interface {
	IDer
	EntityNamer
	IDSetter
	NameSetter
} = (*Entity)(nil)

func NewEntity(id, name string) Entity {
	return Entity{
		id:   id,
		name: name,
	}
}

func (e Entity) ID() string         { return e.id }
func (e Entity) EntityName() string { return e.name }
func (e Entity) SetID(string)       {}
func (e Entity) SetName(string)     {}
