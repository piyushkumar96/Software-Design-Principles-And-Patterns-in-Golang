## Some important points for SRP

* **Closed for Modification**:- New features getting added to the software component, should not have to modify the existing code
* **Open for Extension**:- A software component should be extendable to add a new feature or to add a new behaviour to it.
* Ease of add new features.
* OCP often requires decoupling which in turn, automatically follows SRP.

> Note:- Do not follow OCP blindly otherwise it will lead to end up with huge numbers of structs which will make overall design complex.

## Code

1. This code is not following OCP as if any new type of customer profile will come in future then InsurancePremiumDiscountCalculator code will be changed, so we are violating the principle i.e closed for modification and open for extensibility.  [Code-not-following-SRP](/1.%20SOLID%20Principle/2.%20Open%20Close%20Principle/1.a.code-not-following-OCP.go)
2. This code is following OCP as if any new type of customer profile will come in future then InsurancePremiumDiscountCalculator code will not be changed, so we are following the principle i.e closed for modification and open for extensibility  [Code-following-SRP](/1.%20SOLID%20Principle/2.%20Open%20Close%20Principle/1.b.code-following-OCP.go)