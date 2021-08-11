## Some important points for SRP

* **Cohesion**:- Higher cohesion helps better adherence to the SRP
* **Coupling**:- It is defined as the level of inter dependency between the software components. <u>Loose coupling helps better adherence to SRP. </u>
* Every software component should have one and only one reason to change. Otherwise, more changes will lead to more bugs which also lead to more testing

## Code

1. Low cohesion functions calculateArea/calculatePerimeter and draw/rotate are belonged to same type i.e. Square [Code-not-following-SRP](/SOLID%20Principle/Single%20Responsibility%20Principle/1.a.code-not-following-SRP.go)
2. Now the overall cohesion increase as we separate out the functions [Code-following-SRP](/SOLID%20Principle/Single%20Responsibility%20Principle/1.b.code-following-SRP.go)
3. Tight coupling between the Student and DB. So in future if we want change the DB then it will lead to change in Student functionality. [Code-not-following-SRP](/SOLID%20Principle/Single%20Responsibility%20Principle/2.a.code-not-following-SRP.go)
4. Loose coupling between the Student and DB. So in future if we want to change the DB then it will not lead to change in Student functionality. [Code-following-SRP](/SOLID%20Principle/Single%20Responsibility%20Principle/2.b.code-following-SRP.go)
