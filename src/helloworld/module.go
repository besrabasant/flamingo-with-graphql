package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"

	"flamingo.me/flamingo-app/src/helloworld/interfaces"
)

type Module struct{}

func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct {
	helloController *interfaces.HelloController
}

func (r *routes) Inject(controller *interfaces.HelloController) *routes {
	r.helloController = controller

	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/hello", "hello")
	registry.HandleGet("hello", r.helloController.Get)

	registry.MustRoute("/greetme", "helloWorld.greetme")
	registry.HandleGet("helloWorld.greetme", r.helloController.GreetMe)
}
