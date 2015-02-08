package main

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/datjeff/SecretService/models"
	"github.com/datjeff/SecretService/controllers"
	"os"
	"fmt"
) 

func main() {
  	martiniClassic := martini.Classic()
	
	secretThingCollection := models.NewSecretThingCollection()
	
	controllers.RegisterRestfulService(secretThingCollection, martiniClassic)

	martiniClassic.Use(auth.BasicFunc(controllers.IsAuthorized))
	fmt.Println(os.Getenv("PORT"))
	martiniClassic.RunOnAddr(":"+ os.Getenv("PORT"))
}










