package Models
import(
	"github.com/go-martini/martini"
	"encoding/json"
	"net/http"
	"strconv"
)

func (st *SecretThingCollection) GetPath() string{
	return "/SecretThing"
}

func (st *SecretThingCollection) RestfulDelete(params martini.Params)(int, string){
	var id,err = strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}
	
	err = st.RemoveSecretThing(id)
	if err == nil {
		return http.StatusOK,"Successfully removed thing with id " + params["id"]
	}else{
		return  http.StatusBadRequest, err.Error()
	}
}

func (st *SecretThingCollection) RestfulGet(params martini.Params)(int, string){
	var id,err = strconv.Atoi(params["id"])
	if err != nil {
		return http.StatusBadRequest, "Failed to parse ID"
	}

	var thing, error = st.GetSecretThing(id)

	if error == nil {
		return http.StatusOK, thing.Content
	}else {
		return http.StatusBadRequest, error.Error()
	}

}

//Update
func (st *SecretThingCollection) RestfulPut(params martini.Params, req *http.Request)(int, string){

	var idString = params["id"]
	var result =  http.StatusBadRequest
	var message = "Bad request"

	decoder := json.NewDecoder(req.Body)
    var t SecretThing   
    err := decoder.Decode(&t)
    if err != nil {
        panic(message)
    }

	if idString != "" {
		var id,err = strconv.Atoi(idString)
		if err != nil {
			return http.StatusBadRequest, message
		}
		err = st.UpdateSecretThing(id,t.Content)
		if err != nil {
			message = err.Error()
		}else{
			message = "Successfully updated thing with id " + idString
			result = http.StatusOK
		}
	}else {		
		var newId = st.AddSecretThing(t.Content)
		message = "Successfully put thing with new id " + strconv.Itoa(newId)
		result =  http.StatusOK
	}
	
	return result, message
}
//Create new
func (st *SecretThingCollection) RestfulPost(params martini.Params, req *http.Request)(int, string){
	return http.StatusBadRequest, "Post Not Supported."
}