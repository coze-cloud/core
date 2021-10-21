# core
⚙️ Core module shared across our backend projects

## Installation

Adding *core* to a Go project is as easy as calling this command

```shell
go get github.com/cozy-hosting/core
```

## Using the module

The project must be based on the [uber-go/fx](https://github.com/uber-go/fx) application framework

```go
package main

import "go.uber.org/fx"

func main() {
    // Creates a new fx application
    fx.New(
        // Add the core module to the container
        core.Module,
        fx.Invoke(
            // Also add this, if you want to use GraphQL
            core.UseGraphQl,
        ),
    ).Run()
}
```

## Contained sub-modules

The list of uber/fx modules that are currently available:

- Logrus - https://github.com/sirupsen/logrus
- Clerk database (MongoDB) - https://github.com/cozy-hosting/clerk
- Messenger messaging queue (RabbitMQ) - https://github.com/cozy-hosting/messenger
- Labstack Echo webserver - https://github.com/labstack/echo
- GraphQL Schema / Handling - https://github.com/graphql-go/graphql