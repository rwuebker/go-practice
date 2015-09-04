package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func db() {

	// setting db url
	url := "mongodb://localhost"
	if url == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	// establish a session with the url
	sess, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Unable to connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer sess.Close()

	// SetSafe makes sure errors are thrown if errors occur
	sess.SetSafe(&mgo.Safe{})

	// set the variable collection to the goUser collection in the test DB
	collection := sess.DB("test").C("goUser")

	// create a document modeled after the Person structure
	doc := Person{Id: bson.NewObjectId(), FirstName: "jackson", LastName: "browne"}

	// insert the doc and check for an err
	err = collection.Insert(doc)

	// if error print and exit
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		os.Exit(1)
	}

	// does an update on the collect and throws error if error
	err = collection.Update(bson.M{"_id": doc.Id}, bson.M{"$inc": bson.M{"count": 1}})
	if err != nil {
		fmt.Printf("****Can't update document %v\n", err)
		os.Exit(1)
	}

	// declare a new Person object in the updatedperson variable
	var updatedperson Person

	// this is another method to find a person
	//err = sess.DB("test").C("goUser").Find(bson.M{}).One(&updatedperson)

	// finds a person in the goUsers database
	err = collection.Find(bson.M{}).One(&updatedperson)
	if err != nil {
		fmt.Printf("got an error finding a doc %v\n")
		os.Exit(1)
	}

	// prints the document to the console
	fmt.Printf("Found document: %+v\n", updatedperson)

}
