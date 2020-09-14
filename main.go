package main

import (
	Controllers "./Aplications/Controllers"
	Routes "./Aplications/Routes"
	Entity "./Domain/Entities"
	Helpers "./Domain/Helpers"
	Services "./Domain/Services"
	Infra "./Infrastructure"
	Database "./Infrastructure/Database"
	Repository "./Infrastructure/Repository"
	"go.uber.org/dig"
)

// Todo: Externar o container no futuro
func buildContainer() *dig.Container {
	container := dig.New()
	providers := [7]func(*dig.Container){
		buildEntityProvider,
		buildHelpersProvider,
		buildRepositoryProvider,
		buildServiceProvider,
		buildControllerProvider,
		buildRoutesProvider,
		buildInfraProvider,
	}

	for _, provider := range providers {
		provider(container)
	}

	return container
}

func buildHelpersProvider(container *dig.Container) {
	container.Provide(Helpers.NewHandleController)
}
func buildEntityProvider(container *dig.Container) {
	container.Provide(Entity.NewRoutesConfig)
	container.Provide(Entity.NewServerConfig)
}

func buildControllerProvider(container *dig.Container) {
	container.Provide(Controllers.NewPriceController)
}

func buildServiceProvider(container *dig.Container) {
	container.Provide(Services.NewPriceService)
}

func buildRepositoryProvider(container *dig.Container) {
	container.Provide(Repository.NewPriceRepository)
}

func buildRoutesProvider(container *dig.Container) {
	container.Provide(Routes.NewPriceRoute)
	container.Provide(Routes.NewRoute)
}

func buildInfraProvider(container *dig.Container) {
	container.Provide(Infra.NewRoutes)
	container.Provide(Infra.NewServer)
	container.Provide(Database.NewDatabase)
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
