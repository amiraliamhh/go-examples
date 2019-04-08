package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func logOnFail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("mydb").C("mycollection")

	var person = Person{
		FirstName: "Jon",
		LastName:  "Doe",
		Age:       20,
	}

	// remove all records before inserting
	info, err := collection.RemoveAll(nil)
	logOnFail(err)
	fmt.Printf("data is removed, info: %+v\n", info)

	// create new record
	err = collection.Insert(&person)
	logOnFail(err)

	// query records
	result := Person{}
	err = collection.Find(bson.M{"firstname": "Jon"}).One(&result)
	logOnFail(err)
	fmt.Printf("data is inserted: %+v\n", result)

	// update
	newPerson := Person{
		FirstName: "Joe",
		LastName:  "Smith",
		Age:       33,
	}
	err = collection.Update(person, newPerson)
	logOnFail(err)

	err = collection.Find(bson.M{"firstname": "Joe"}).One(&result)
	logOnFail(err)
	fmt.Printf("data is updated: %+v\n", result)

}
