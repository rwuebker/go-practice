package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

type Person struct {
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Id        bson.ObjectId `bson:"_id"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	var person Person

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//	defer r.Body.Close()

	// if err := r.Body.Close(); err != nil {
	//   panic(err)
	// }

	if err := json.Unmarshal(body, &person); err != nil {
		panic(err)
	}

	fmt.Println("this is person:", person)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		panic(err)
	}

}
