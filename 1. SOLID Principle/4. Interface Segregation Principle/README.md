## Some important points for ISP

* No client should be forced to depend on methods it does not use.
* Breaking down the bigger interface to smaller interface so to increase cohesion between the functions inside that interface.
* Techniques to identify violations of ISP
  1. Fat interfaces.
  2. Functions inside the interface with low cohesion.
  3. Empty function implementations.

## Code

1. This code is not following ISP as we left unsupported func for specific type printers unimplemented. This will make our code ambiguous. for eg. Fax is not supported for HPPrinterNScanner similarly Scan and Fax are not supported by CannonPrinter.  [Code-not-following-ISP](/1.%20SOLID%20Principle/4.%20Interface%20Segregation%20Principle/1.a.code-not-following-ISP.go)
2. This code is following ISP as we break the fat interface to smaller one.  So now there is no need to link the unsupported functions. [Code-following-ISP](/1.%20SOLID%20Principle/4.%20Interface%20Segregation%20Principle/1.b.code-following-ISP.go)