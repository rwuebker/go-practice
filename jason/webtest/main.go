package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Person struct {
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Id        bson.ObjectId `bson:"_id_`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("this is r.Body: %+v \n", r.Body)

	var person Person

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &person); err != nil {
		panic(err)
	}

	fmt.Println("this is person:", person)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriterHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		panic(err)
	}

}

fun main() {
    url := "mongodb://localhost"

    sess, err := mgo.Dial(url)

    sess.SetSafe(&mgo.Safe{})

    collection := sess.DB("test").C("goUser")

    doc := Person{Id: bson.NewObjectId(), FirstName : "bob", LastName : "jones"}

    err = collection.Insert(doc)

    err = collection.Update(bson.M{"_id": doc.Id}, bson.M{"$inc": bson.M{"cound": 1}})

    fmt.Printf("this is err %v\n", err)

    var updatedperson Person

    err = collection.Find(bson.M{}).One(&updatedperson)

    fmt.Printf("Found doc: %+v\n", updatedperson)

    myVariable := funcToTest()
    
