# payutils-http-fiber

[Fiber](https://github.com/gofiber/fiber) HTTP driver for
[payutils](https://go.gh.ink/payutils/v3).

Adapts a GoFiber router so payutils can register provider callback routes on it.
Standard-library `http.HandlerFunc`s are bridged to GoFiber via Fiber's
`adaptor.HTTPHandler`.

## Module

```
go.gh.ink/payutils/http/fiber/v3
```

```bash
go get go.gh.ink/payutils/http/fiber/v3
```

## Usage

Blank-import the driver (it self-registers under the name `fiber`) and pass a
GoFiber router (`*fiber.App` or a route group) in `Config.Instances`:

```go
import (
    "github.com/gofiber/fiber/v3"

    "go.gh.ink/payutils/v3/client"
    "go.gh.ink/payutils/v3/model"

    httpFiber "go.gh.ink/payutils/http/fiber/v3"
    _ "go.gh.ink/payutils-pay-alipay/v3"
)

app := fiber.New()

c, err := client.NewClient(model.Config{
    Endpoint:    "https://api.example.com",
    Instances:   map[string]any{httpFiber.Name: app}, // accepts fiber.Router
    Credentials: model.C{ /* ... */ },
    Contract:    myContract{},
})
```

payutils then registers `POST /{provider}/callback` on the router for each
configured provider.

## Accepted instance

Anything implementing `fiber.Router` — `*fiber.App` and route groups
(`app.Group("/pay")`). Passing an unsupported value makes `NewInstance` return
`errors.ErrUnsupportedInstance`.

## Supported verbs

`Get` · `Post` · `Put` · `Patch` · `Delete` · `Head` · `Options` · `Any`
(`Any` maps to Fiber's `All`). payutils only uses `Post` for callbacks today;
the full set is implemented for future use.

## License

See [LICENSE](LICENSE).
