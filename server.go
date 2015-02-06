package main

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"SecretService/Models"
	"SecretService/Controllers"
) 

func main() {
  	martiniClassic := martini.Classic()
	
	secretThingCollection := Models.NewSecretThingCollection()
	
	Controllers.RegisterRestfulService(secretThingCollection, martiniClassic)

	martiniClassic.Use(auth.BasicFunc(Controllers.IsAuthorized))

	martiniClassic.Run()
}










