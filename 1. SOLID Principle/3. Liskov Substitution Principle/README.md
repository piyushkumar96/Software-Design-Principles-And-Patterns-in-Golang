## Some important points for LSP

* Objects should be replaceable with their subtypes without affecting the correctness of the program.
* Change the "Is-A" way of thinking.
* If it looks like a duck and quacks like a duck but it needs batteries, you probably have the wrong abstraction.

>Note:- Golang is **safe from LSP**, because in **golang we don’t have inheritance**, we have a more powerful tool that is the composition. Anyway composition does not allow to substitute parent struct by child struct.
 

## Code

1. This code is not following LSP as we have to call external function to set the new discount for InHouseProduct. After then, we call getDiscount to get discount for InHouseProduct.  [Code-not-following-LSP](/1.%20SOLID%20Principle/3.%20Liskov%20Substitution%20Principle/1.a.code-not-following-LSP.go)
2. This code is following LSP as we arw not calling external function to set the new discount for InHouseProduct. We directly call getDiscount to get discount for InHouseProduct. [Code-following-LSP](/1.%20SOLID%20Principle/3.%20Liskov%20Substitution%20Principle/1.b.code-following-LSP.go)