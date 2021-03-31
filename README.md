# Black Box Router

Package blackbox provides simple to use, extensible, flexible, and easily
testable typesafe server routing for your web apps and microservices.

Blackbox does that by utilising a functional approach to request handling,
where routing nodes (type Router) provide filtering handling and return an
exposed Response instance that can be tested and examined.

- For more information on testing, please see the [test](test/) folder.
- To see a working example, look at this [example](example/).

## Why Black Box

1. Seamless interop with the standard `net/http.Server`: `blackbox.Router`
   implements the `net/http.Handler` interface
2. Pleasure to test: due to the convenient `Handler` interface definition,
   `blackbox.Handler` actually **returns** `blackbox.Response` instead of
   writing it to `net/http.ResponseWriter`
3. Well tested with an **average coverage of 85%** across all packages
4. Extendable architecture: if you need more filters or additional middleware,
   just implement your own by implementing interfaces like `blackbox.Filter`.
5. As a bonus, `blackbox.Response` supports JSON serialization out of the box:
   just call `resp.EncodeJSON(object)` and you're golden!
