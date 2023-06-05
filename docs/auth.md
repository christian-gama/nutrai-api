# Auth
The HTTP controller package provides a structured way of handling HTTP requests and responses in the Nutrai API. It's responsible for the flow of incoming HTTP requests, binding of request data to a defined input model, validation of the input data, and execution of the actual handler function with validated input data. The controller is where we set that an endpoint is public or private.

## Controller

The `Controller` interface allows us to create route handlers with specific behaviors. The `Security` method is one of the key features here, which returns a security config. When set to nil, the handler will use the default authentication method, which is JWT, meaning that the route is private. On the other hand, if it's set to any other value, the handler will use the specified authentication method.

```go
type Controller interface {
	Handle(ctx *gin.Context)
	Method() http.Method
	Path() Path
	Params() Params
	Security() Security
	RPM() int
}
```

## Input with User

One key detail in the controller design is how the current user is passed into the handler function. This is done through the `Input` struct, where you can add a `User` field with the `ctx:"current_user"` tag. This tag informs the application to automatically set the user in the input.

```go
type Input struct {
	User *user.User `ctx:"current_user" json:"-"`
}
```

The `User` object is populated in the `Handle` function of `controllerImpl`, specifically through the `http.ExtractCurrentUser(ctx, &c.input)` function call. It fetches the authenticated user from the context and assigns it to the input.

```go
func (c controllerImpl[Input]) Handle(ctx *gin.Context) {
	http.ExtractBody(ctx, &c.input)
	http.ExtractQuery(ctx, &c.input)
	http.ExtractParams(ctx, &c.input)
	http.ExtractCurrentUser(ctx, &c.input) // This is where the user is set

	// Rest of the function...
}
```

## Auth Middleware
The Nutrai API uses middleware to add authentication to its endpoints. The middleware will use the `Security` method of the controller to determine if the route is public or private. If it's private, the middleware will use the authentication method specified in the `Security` method. Otherwise, it will skip the authentication process. 

```go
func (r *routes) addAuthIfNeeded(
	c controller.Controller,
	handlers []gin.HandlerFunc,
) []gin.HandlerFunc {
	if c.Security().Middleware() != nil {
		handlers = slice.Unshift(handlers, c.Security().Middleware().Handle).Build()
	}

	return handlers
}
```

## Setting the Security Method
At boot time, the application will set the security method of each `controller.Security`. If the method is not set at boot time, it will exit the application with an error.
