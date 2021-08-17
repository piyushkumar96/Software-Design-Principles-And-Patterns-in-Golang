package main

import (
	"fmt"
)

// This code is not following LSP as we have to call external function to set the new discount for InHouseProduct
// After then we call getDiscount to get discount for InHouseProduct

type Product struct {
	discount float32
}

func (p Product) getDiscount() float32 {
	return p.discount
}

type InHouseProduct struct {
	Product
}

func (ihp *InHouseProduct) applyExtraDiscount() {
	ihp.discount = ihp.discount * 1.5
}

func main() {

	p1 := Product{
		discount: 2.0,
	}
	p2 := InHouseProduct{
		p1,
	}

	fmt.Println(p1.getDiscount())
	p2.applyExtraDiscount()
	fmt.Println(p2.getDiscount())

}
