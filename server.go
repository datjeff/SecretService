package main

import(
	"github.com/go-martini/martini"
	"SecretService/Models"
	"SecretService/Controllers"
) 

func main() {
  martiniClassic := martini.Classic()
        secretThingCollection := Models.NewSecretThingCollection()
        Controllers.RegisterRestfulService(secretThingCollection, martiniClassic)
        martiniClassic.Run()
}










