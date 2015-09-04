package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
)

type Person struct {
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Id        bson.ObjectId `bson:"_id"`
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:8080", nil)

	url := "mongodb://localhost"
	if url == "" {
		fmt.Println("url not provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Unable to connect to mongo", err)
		os.Exit(1)
	}

	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("test").C("person")

	doc := Name{Id: bson.NewObjectId(), FirstName: "Jason", LastName: "Kautz"}

	err = collection.Insert(doc)

	if err != nil {
		fmt.Printf("Can't insert document:", err)
		os.Exit(1)
	}
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJSONResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "hello, %s!\n", string(response))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello,"))
}

func getJSONResponse() ([]byte, error) {
	var name Name

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &person); err != nil {
		panic(err)
	}

	fmt.Println("hiya, ", name)

}
