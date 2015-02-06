package DataAccess

import (
	"github.com/couchbaselabs/go-couchbase"
)


type CouchbaseDal struct {
	Location string
}

func GetBytes(dal *CouchbaseDal, key string) ([]byte) {
	c, err := couchbase.Connect("http://localhost:8091/")
	if err != nil {
		panic(err)
	}

	pool, err := c.GetPool("default")
	if err != nil {
		panic(err)
	}

	bucket, err := pool.GetBucket("SecretThing")
	if err != nil {
		panic(err)
	}

	bytes, err := bucket.GetRaw(key)
	if err != nil {
		panic(err)
	}

	return bytes
}

func SetBytes(dal *CouchbaseDal, key string, value []byte) (err string) {
	c, err := couchbase.Connect("http://localhost:8091/")
		if err != nil {
			panic(err)
		}

		pool, err := c.GetPool("default")
		if err != nil {
			panic(err)
		}

		bucket, err := pool.GetBucket("SecretThing")
		if err != nil {
			panic(err)
		}

		//decoder := json.NewDecoder(req.Body)
	    //var t string   
	    //err = decoder.Decode(&t)
	    //if err != nil {
	    //    panic(err)
	    //}

		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(bytes))

		added, err := bucket.AddRaw(key, 1, bytes)
		if err != nil {
			panic(err)
		}

		if !added {
			panic("add failed!")
		}

		return nil
}
