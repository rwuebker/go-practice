package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Id        bson.ObjectId `bson:"_id"`
}

// function to handle the route of hello world
func handleIt(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("this is r.Body: %+v \n", r.Body)

	var person Person

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error occurred", err)
	}

	// if err := r.Body.Close(); err != nil {
	//   panic(err)
	// }

	if err := json.Unmarshal(body, &person); err != nil {
		log.Println(err)
	}

	fmt.Println("this is person:", person)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(person); err != nil {
		log.Println(err)
	}
}

func main() {
	// setting db url
	//url := "mongodb://localhost"
	//if url == "" {
	//log.Fatal("no connection string provided")
	//}

	// establish a session with the url
	//sess, err := mgo.Dial(url)
	//if err != nil {
	//panic(err)
	//}
	//defer sess.Close()

	// SetSafe makes sure errors are thrown if errors occur
	//sess.SetSafe(&mgo.Safe{})

	// set the variable collection to the goUser collection in the test DB
	//collection := sess.DB("test").C("goUser")

	// create a document modeled after the Person structure
	//doc := Person{Id: bson.NewObjectId(), FirstName: "jackson", LastName: "browne"}

	// insert the doc and check for an err
	//err = collection.Insert(doc)

	// if error print and exit
	//if err != nil {
	//fmt.Printf("Can't insert document: %v\n", err)
	//os.Exit(1)
	//}

	// does an update on the collect and throws error if error
	//err = collection.Update(bson.M{"_id": doc.Id}, bson.M{"$inc": bson.M{"count": 1}})
	//if err != nil {
	//fmt.Printf("****Can't update document %v\n", err)
	//os.Exit(1)
	//}

	// declare a new Person object in the updatedperson variable
	var updatedperson Person

	// this is another method to find a person
	//err = sess.DB("test").C("goUser").Find(bson.M{}).One(&updatedperson)

	// finds a person in the goUsers database
	//err = collection.Find(bson.M{}).One(&updatedperson)
	//if err != nil {
	//fmt.Printf("got an error finding a doc %v\n")
	//os.Exit(1)
	//}

	// prints the document to the console
	fmt.Printf("Found document: %+v\n", updatedperson)

	// function to test

	//myVariable := funcToTest()
	//fmt.Println("this is myVariable:", myVariable)
	// handles the /helloWorld endpoint
	http.HandleFunc("/helloWorld", handleIt)

	// creates a listener on 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("there was an error")
		panic(err)
	}
}

func funcToTest() string {
	return "whats up"
}
