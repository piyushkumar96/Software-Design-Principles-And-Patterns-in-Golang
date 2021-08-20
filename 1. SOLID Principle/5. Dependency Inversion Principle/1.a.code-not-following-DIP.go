package main

import "fmt"

// This code is not following DIP as high level module(Investigate) is coupled with low level module data(relations).
// this means high level module depends upon low level module not on abstraction.

type Relationship int

const (
	Parent Relationship = iota
	Child
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{
		parent,
		Parent,
		child,
	})

	rs.relations = append(rs.relations, Info{
		child,
		Child,
		parent,
	})
}

type Research struct {
	relationships Relationships
}

// This is high level module
func (res *Research) Investigate(personName string) {
	// here we are accessing the low level module data i.e. relations
	relations := res.relationships.relations
	for _, rel := range relations {
		if rel.from.name == personName && rel.relationship == Parent {
			fmt.Println(personName + " is parent of " + rel.to.name)
		}
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{relationships}
	research.Investigate("John")
}
