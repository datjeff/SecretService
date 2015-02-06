package main

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/cummingsi1993@gmail.com/SecretService/Models"
	"github.com/cummingsi1993@gmail.com/SecretService/Controllers"
) 

func main() {
  	martiniClassic := martini.Classic()
	
	secretThingCollection := Models.NewSecretThingCollection()
	
	Controllers.RegisterRestfulService(secretThingCollection, martiniClassic)

	martiniClassic.Use(auth.BasicFunc(Controllers.IsAuthorized))

	martiniClassic.Run()
}










