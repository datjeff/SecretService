package main

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/cummingsi1993@gmail.com/SecretService/Models"
	"github.com/cummingsi1993@gmail.com/SecretService/Controllers"
	//"encoding/json"
	"net/http"
) 

func main() {
  	martiniClassic := martini.Classic()
	//martiniClassic.RunOnAddr(":3000")
	secretThingCollection := Models.NewSecretThingCollection()
	_ = secretThingCollection //temporarily
	
	//Controllers.RegisterRestfulService(secretThingCollection, martiniClassic)

	martiniClassic.Use(auth.BasicFunc(Controllers.IsAuthorized))

	martiniClassic.Get("/SecretThing/:key", func(params martini.Params) string {
		//DataAccess.Get(new(CouchbaseDal), params[key])
		return "woot"
	})

	martiniClassic.Put("/SecretThing/:key", func(params martini.Params, req *http.Request) string {
		

		return "ok"
	})

	martiniClassic.Get("/hello/:name", func(params martini.Params) string {
	  return "Hello " + params["name"]
	})

	martiniClassic.Run()
}










