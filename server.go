package main

import(
	"github.com/cummingsi1993@gmail.com/go-data_access"
	"github.com/cummingsi1993@gmail.com/SecretService/encryption"
	//github.com/cummingsi1993@gmail.com/SecretService/Models"
	//"github.com/cummingsi1993@gmail.com/SecretService/Controllers"
	//"encoding/json"
	"fmt"
	"io/ioutil"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/base64"
    "strings"
    "errors"
) 

func main() {
	//Set up configuration
	url := "http://localhost:8091/"
	_ = url
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/SecretThing/{key}", ServeHTTP(SecretThingEndpoint))

	log.Fatal(http.ListenAndServe(":3000", router))
}

type SecretThing struct {
	Key string
	Value []byte
}

func SecretThingEndpoint(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("Secret thing endpoint was hit.")

	switch r.Method { 
	case "POST": 
		err := PutOrPostSecretThing(w, r)
	case "PUT": 
		err := PutOrPostSecretThing(w, r)
	case "DELETE" : 
		err := DeleteSecretThing(w, r)
	case "GET" : 
		err := GetSecretThing(w, r)
	}
	return
}

func GetAuthenticationString(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)
	fmt.Println("made it past getting the header")
	if len(auth) != 2 || auth[0] != "Basic" {
        http.Error(w, "bad syntax", http.StatusBadRequest) //This is a good strategy for handling errors
        return nil, errors.New("Bad Syntax")
    }

    payload, err := base64.StdEncoding.DecodeString(auth[1])
    return payload, err
}

func GetSecretThingFromRequest(r *http.Request) (*SecretThing, error){
	vars := mux.Vars(r)
	key := vars["key"]
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil { return nil, err }
	secretThing := SecretThing{key, bytes}
	return &secretThing, err
}


func PutOrPostSecretThing(w http.ResponseWriter, r *http.Request) (err error) {
	thing, err := GetSecretThingFromRequest(r)
	url := "http://localhost:8091/"

	authString, err := GetAuthenticationString(w, r)
	fmt.Println(string(authString))
	encrypted, err := encryption.Encrypt(encryption.Hash(authString), thing.Value)	
	fmt.Println(err)
	thing.Value = encrypted
	err = couchbase.Set(url, thing.Key, thing)
	fmt.Println(err)
	return 
}

func GetSecretThing(w http.ResponseWriter, r *http.Request) (err error) {
	url := "http://localhost:8091/"
	thing := SecretThing{}
	vars := mux.Vars(r)
	key := vars["key"]
	err = couchbase.Get(url, key, &thing)

	authString, err := GetAuthenticationString(w, r)
	decrypted, err := encryption.Decrypt(encryption.Hash(authString), thing.Value)

	w.Write(decrypted)
	return
}

func DeleteSecretThing(w http.ResponseWriter, r *http.Request) (err error) {
	url := "http://localhost:8091/"
	vars := mux.Vars(r)
	key := vars["key"]
	err = couchbase.Remove(url, key)
	return
}

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        http.Error(w, err.Error(), 500)
    }
}