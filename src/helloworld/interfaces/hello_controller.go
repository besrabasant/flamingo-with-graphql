package interfaces

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	HelloController struct {
		responder *web.Responder
	}

	helloViewData struct {
		Name string
	}
)

func (controller *HelloController) Inject(responder *web.Responder) *HelloController {
	controller.responder = responder

	return controller
}

func (controller *HelloController) Get(ctx context.Context, r *web.Request) web.Result {
	return controller.responder.Render("hello", helloViewData{
		Name: "World",
	})
}

func (controller *HelloController) GreetMe(ctx context.Context, r *web.Request) web.Result {
	name, err := r.Query1("name")

	if err != nil {
		name = "World (default)"
	}

	return controller.responder.Render("hello", helloViewData{
		Name: name,
	})
}
