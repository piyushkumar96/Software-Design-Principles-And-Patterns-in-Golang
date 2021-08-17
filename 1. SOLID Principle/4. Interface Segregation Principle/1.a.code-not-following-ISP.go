package main

import (
	"fmt"
)

// This code is not following ISP as we left unsupported func for specific type printers unimplemented.
// this will make our code ambiguous
// for eg. Fax is not supported for HPPrinterNScanner similarly Scan and Fax are not supported by CannonPrinter.

type Document struct {
	name string
}

type IMultiFunction interface {
	Print(d Document)
	Scan(d Document)
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

func (pns HPPrinterNScanner) Fax(d Document) {
	panic("Not supported")
}

type CannonPrinter struct {
}

func (cp CannonPrinter) Print(d Document) {
	fmt.Println("Print the document:- ", d.name)
}

func (cp CannonPrinter) Scan(d Document) {
	panic("Not supported")
}

func (cp CannonPrinter) Fax(d Document) {
	panic("Not supported")
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
	// below func call will failed the program
	// hpPNS.Fax(doc)

	cannonP := CannonPrinter{}
	cannonP.Print(doc)
	// below functions call will failed the program
	// cannonP.Scan(doc)
	// cannonP.Fax(doc)
}
