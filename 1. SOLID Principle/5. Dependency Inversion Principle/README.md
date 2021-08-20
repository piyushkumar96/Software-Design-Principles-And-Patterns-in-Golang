## Some important points for DIP

* High level modules should not depend upon low level modules. Both should depend upon abstractions.
* Abstraction should not depend upon details. Details should be depended on abstractions.

## Code

1. This code is not following DIP as high level module(Investigate) is coupled with low level module data(relations). This means high level module depends upon low level module not on abstraction. [Code-not-following-DIP](/1.%20SOLID%20Principle/5.%20Dependency%20Inversion%20Principle/1.a.code-not-following-DIP.go)
2. This code is following DIP as high level module(Investigate) is not coupled with low level module data(relations). High level module is accessing using RelationshipRepository which is not low level module. This means high level module does not depends upon low level module, Both depend on abstraction i.e. RelationshipRepository. [Code-following-DIP](/1.%20SOLID%20Principle/5.%20Dependency%20Inversion%20Principle/1.b.code-following-DIP.go)