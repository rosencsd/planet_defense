package ecs

// Components is the global component list containing all entities and components
var Components []IComponent

func init() {
	Components = []IComponent{}
}

// Component has an id and can have Data
type IComponent interface {
	//	GetId() int
	IsEntity() bool // should always return false
}

// Component has the id and the interface functions
type Component struct {
	//	id int
}

// GetId returns the id of the component
//func (c *Component) GetId() int {
//	return c.id
//}

func (c *Component) IsEntity() bool {
	return false
}

// Name is an IComponent for storing an entity's name
type Name struct {
	Str string
	Component
}

// Mutable is an IComponent for indicating that an entity is allowed to be changed (not it has no fields, if its in the list the entity is mutable)
type Mutable struct {
	Component
}

func NewComponent(comp IComponent) (ID int) {
	ID = len(Components)
	Components = append(Components, comp)
	return ID
}

// Entities

// Entity has an id and a list of components an entity is also a component
// In retrospect I don't think we need an entity interface just the type
// type IEntity interface {
// 	GetId() int
// 	GetCompIDs() []int
// 	IsEntity() bool // should always return true
// }

// Entity is a struct that satisfies IComponent and has a list of component IDs
type Entity struct {
	//	id    int
	compIDs []int
}

//func (e *Entity) GetId() int {
//	return e.id
//}

func (e Entity) GetCompIDs() []int {
	return e.compIDs
}

func (e Entity) IsEntity() bool {
	return true
}

// NewEntity creates new named entities with components
func NewEntity(name string, compIDs []int) (ID int) {
	ID = NewComponent(&Name{Str: name})
	ID = NewComponent(&Entity{compIDs: append(compIDs, ID)})
	return ID
}

// Query Functions

// Systems
