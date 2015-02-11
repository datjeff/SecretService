package controllers

import(
	"SecretService/interfaces"
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

C:\Users\Jeff\.ssh\id_rsa.pub

type "C:\Users\Jeff\.ssh\id_rsa.pub" | ssh root@104.236.196.171 "sudo sshcommand acl-add dokku JeffLaptop"