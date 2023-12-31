# Platform as Prompt

Experimental project for investigating generative AI for an entire platform:

- A new paradigm, PaP (Platform as Prompts), can software architects describe good enough an entire platform with production quality?
- Can PaP reduce the numbers of engineers to keep a platform running?
- Do we need new software patterns?
- Is it good enough to support iterative design, additions and modifications?
- Can it support monitoring, debugging and fixing bugs?
- Can it be used for the frontend?

## Context

Constant incremental delivery is the main paradigm in the tech companies. Basically, it acknowledges:

- Humans are incapable of describing **accurately** what to build when the information is incomplete.
- Even with complete information of a system, development teams can't build **properly** a large system in one go.

The solution is to build the system incrementally, and deliver it to the users as soon as possible. The users can then provide feedback, and the product team can improve the system accordingly.

The incremental delivery is a good solution, but it has a few drawbacks:

- The focus of the technical/product design is in the short term.
- Iterating in complexity hops is difficult (example: from a single server to a distributed system).
- Refactors are often postponed in favor of new features.

With the PaP (Platform as Prompts) paradigm, we're looking to improve a few aspects of the incremental delivery:

- The capacity of a development team to deliver larger systems in one iteration. This should help putting the focus of the technical design in the **mid term** as opposed to the short term. Technical debt should be reduced and the quality of the design should be improved. As the design time counts as development time, developers/architects can invest more time in designing the solution. The issue we can have in here is the is to fall into the analysis paralysis. 
- The capacity to iterate and refactor existing functionality, as the system is rebuild entirely at each iteration with the updated PROMPTS

## Approach

Typically, generative AIs tend to be stochastic. That implies certain randomness in the responses and the code generated, but more importantly, it means **the code generation process is not deterministic**. The AI generated code might be correct (functionally or semantically) for a specific generation and wrong in the next one.

As a solution, the PaP framework must be built around:

- (1) The capacity to re-generate small pieces (modules) of the system only wihtout affecting others.
- (2) The capacity to accurately, quickly test that modules can speak correctly between them.
- (3) The capacity to accurately, quickly test that generated code does comply with the functional descriptions.
- (4) Being deterministic in the interfaces between modules, specifiying them manually.


---
**The PaP framework can then be described as:**

A framework for building and mantaining a platform given the definitions of:
- data: when databases should be used
- interfaces: of the services, as modules
- component test: as a functional description (for methods and UI components)
- integration tests: as the behavior the system has to have together
---



