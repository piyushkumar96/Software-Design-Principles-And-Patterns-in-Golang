package main

import (
	"fmt"
)

// This code is following ISP as we break the fat interface to smaller one.
// So now there is no need to link the unsupported functions.

type Document struct {
	name string
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxing interface {
	Fax(d Document)
}

type XeroxWorkCentre struct {
}

func (xwc XeroxWorkCentre) Print(d Document) {
	fmt.Println("Print the document:- ", d.name)
}

func (xwc XeroxWorkCentre) Scan(d Document) {
	fmt.Println("Scan the document:- ", d.name)
}

func (xwc XeroxWorkCentre) Fax(d Document) {
	fmt.Println("Fax the document:- ", d.name)
}

type HPPrinterNScanner struct {
}

func (pns HPPrinterNScanner) Print(d Document) {
	fmt.Println("Print the document:- ", d.name)
}

func (pns HPPrinterNScanner) Scan(d Document) {
	fmt.Println("Scan the document:- ", d.name)
}

type CannonPrinter struct {
}

func (cp CannonPrinter) Print(d Document) {
	fmt.Println("Print the document:- ", d.name)
}

func main() {
	doc := Document{
		name: "Sample",
	}
	// all functionalities for xeroxWC will work fine
	xeroxWC := XeroxWorkCentre{}
	xeroxWC.Print(doc)
	xeroxWC.Scan(doc)
	xeroxWC.Fax(doc)

	hpPNS := HPPrinterNScanner{}
	hpPNS.Print(doc)
	hpPNS.Scan(doc)

	cannonP := CannonPrinter{}
	cannonP.Print(doc)
}
