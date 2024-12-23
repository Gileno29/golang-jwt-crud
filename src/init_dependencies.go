package main

func initDependenciesx(database *mongo.Database)(controller.UserControllerInterface, error){

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service), nil
}