package main

import "fmt"

// This code is following DIP as high level module(Investigate) is not coupled with low level module data(relations).
// high level module is accessing using RelationshipRepository which is not low level module
// this means high level module does not depends upon low level module, Both depend on abstraction i.e. RelationshipRepository.

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

type RelationshipRepository interface {
	FindAllChildren(name string) []*Person
}

func (rs *Relationships) FindAllChildren(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}

	return result
}

type Research struct {
	relRepo RelationshipRepository
}

// This is high level module
func (res *Research) Investigate(personName string) {
	// here we are not accessing the low level module data i.e. relations directly
	for _, child := range res.relRepo.FindAllChildren(personName) {

		fmt.Println(personName + " is parent of " + child.name)

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

	research := Research{&relationships}
	research.Investigate("John")
}
