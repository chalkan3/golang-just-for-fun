package main

import (
	Routes "./Aplications/Routes"
	Entity "./Domain/Entities"
	Infra "./Infrastructure"
	"go.uber.org/dig"
)

// Todo: Externar o container no futuro
func buildContainer() *dig.Container {
	container := dig.New()
	buildEntityProvider(container)
	buildRoutesProvider(container)
	buildInfraProvider(container)

	return container
}

func buildEntityProvider(container *dig.Container) {
	container.Provide(Entity.NewRoutesConfig)
	container.Provide(Entity.NewServerConfig)
}

func buildRoutesProvider(container *dig.Container) {
	container.Provide(Routes.NewPriceRoute)
	container.Provide(Routes.NewRoute)
}

func buildInfraProvider(container *dig.Container) {
	container.Provide(Infra.NewRoutes)
	container.Provide(Infra.NewServer)
}

func main() {
	container := buildContainer()
	err := container.Invoke(func(server *Infra.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
