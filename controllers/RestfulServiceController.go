package controllers

import(
<<<<<<< HEAD
	"github.com/datjeff/SecretService/interfaces"
=======
	"github.com/cummingsi1993@gmail.com/SecretService/Interfaces"
>>>>>>> upstream/master
	"github.com/go-martini/martini"
)

func RegisterRestfulService(restModel interfaces.IRestful, classicMartini *martini.ClassicMartini){
	path := restModel.GetPath()

	classicMartini.Get(path, restModel.RestfulGet)
	classicMartini.Get(path+"/:id", restModel.RestfulGet)

	classicMartini.Put(path, restModel.RestfulPut)
	classicMartini.Put(path+"/:id", restModel.RestfulPut)

	classicMartini.Post(path, restModel.RestfulPost)
	classicMartini.Post(path+"/:id", restModel.RestfulPost)

	classicMartini.Delete(path, restModel.RestfulDelete)
	classicMartini.Delete(path+"/:id", restModel.RestfulDelete)
}
