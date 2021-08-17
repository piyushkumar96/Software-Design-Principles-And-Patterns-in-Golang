package main

import (
	"fmt"
)

// This code is following LSP as we arw not calling external function to set the new discount for InHouseProduct
// We directly call getDiscount to get discount for InHouseProduct.

type Product struct {
	discount float32
}

func (p Product) getDiscount() float32 {
	return p.discount
}

type InHouseProduct struct {
	Product
}

func (ihp InHouseProduct) getDiscount() float32 {
	ihp.applyExtraDiscount()
	return ihp.discount
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
	fmt.Println(p2.getDiscount())

}
