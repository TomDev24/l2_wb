package main

import "fmt"

// We can multiple visitors, which satisfy this interface
type Visitor interface {
	VisitPerson(*Person)
	VisitOrganization(*Organization)
}

// Clasess which behaviours we need to extent
type Person struct {
	name  string
	email string
}

func (p *Person) Accept(v Visitor) {
	v.VisitPerson(p)
}

type Organization struct {
	name    string
	address string
}

func (o *Organization) Accept(v Visitor) {
	v.VisitOrganization(o)
}

// Concrete visitor implementation
type EmailVisitor struct{}

func (e *EmailVisitor) VisitPerson(p *Person) {
	fmt.Printf("Sending email to %s at %s\n", p.name, p.email)
}

func (e *EmailVisitor) VisitOrganization(o *Organization) {
	fmt.Printf("Sending mail to %s at %s\n", o.name, o.address)
}

func main() {
	p1 := Person{name: "Alice", email: "alices@example.com"}
	p2 := Person{name: "Bob", email: "bob@example.com"}
	org := Organization{name: "Acme Inc.", address: "123 Main St."}

	visitor := &EmailVisitor{}

	p1.Accept(visitor)
	p2.Accept(visitor)
	org.Accept(visitor)
}
