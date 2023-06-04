# Auth
The HTTP controller package provides a structured way of handling HTTP requests and responses in the Nutrai API. It's responsible for the flow of incoming HTTP requests, binding of request data to a defined input model, validation of the input data, and execution of the actual handler function with validated input data. The controller is where we set that an endpoint is public or private.

## Controller

The `Controller` interface allows us to create route handlers with specific behaviors. The `IsPublic` method is one of the key features here, which returns a boolean value. When set to `false`, the handler will require authentication, meaning that the route is private. On the other hand, if it's set to `true`, the handler will not require authentication, making the route public. 

```go
type Controller interface {
	Handle(ctx *gin.Context)
	Method() http.Method
	Path() Path
	Params() Params
	IsPublic() bool
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
The Nutrai API uses middleware to add authentication to its endpoints. The middleware will check whether the endpoint is marked as public or not. If it's not public, the authMiddleware will be invoked before the handler function, ensuring only authenticated users can access the endpoint. This is implemented in the addAuthIfNeeded function of the routes package:

```go
func (r *routes) addAuthIfNeeded(
	controller controller.Controller,
	handlers []gin.HandlerFunc,
) []gin.HandlerFunc {
	if !controller.IsPublic() {
		handlers = slice.Unshift(handlers, authMiddleware.MakeAuth().Handle).Build()
	}
	return handlers
}
```

This way, we ensure that only authenticated users can access private routes.