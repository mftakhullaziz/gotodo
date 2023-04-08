## todos service
An Implement Hexagonal Pattern in Go

### background
Hexagonal architecture is a software design pattern that separates the core business logic of an application from the technical implementation details. This architecture is also known as **"ports and adapters"** because it uses a series of adapters to connect the application to the outside world through a set of defined ports.

The key idea behind hexagonal architecture is to focus on the core business logic of the application, which is encapsulated in the "domain" layer. This layer defines the domain models, business rules, and use cases for the application.

The other layers of the architecture are responsible for handling technical details such as data storage, networking, and user interface. These layers are organized around the domain layer and communicate with it through defined interfaces or "ports".

Hexagonal architecture provides several benefits, including:

- Separation of concerns: By separating the business logic from the technical implementation details, hexagonal architecture allows developers to focus on the core domain logic without being distracted by implementation details.
- Testability: Hexagonal architecture makes it easier to test the business logic of an application in isolation, without being coupled to specific implementation details such as a database or user interface.
- Flexibility: By defining clear interfaces between the different layers of an application, hexagonal architecture makes it easier to modify or replace specific components without affecting the overall architecture.
- Scalability: Hexagonal architecture can be scaled horizontally by adding more instances of the same component, or vertically by adding more powerful hardware to support the application.

However, hexagonal architecture can also be more complex to implement than traditional layered architectures. It requires careful planning and design to ensure that the interfaces between the layers are well-defined and flexible enough to support future changes to the application.

### structure hexagonal
![pattern](https://github.com/mftakhullaziz/gotodos/blob/main/docs/hexago-pattern.png)

### makefile
How running using makefile

How to use makefile 'make <target>' where <target> is one of the following:
- build/service  [builds the executable]
- run/test       [runs unit tests]
- run/service    [builds and runs the program]
- clean          [cleans the build directory]
- clean/test     [cleans the cache tests]

How to use command :
- make build/service
- make run/test
- make run/service
- make clean
- make clean/test


