// https://labix.org/mgo
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
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("mydb").C("mycollection")

	// create new record
	err = collection.Insert(&Person{
		FirstName: "Jon",
		LastName:  "Doe",
	})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = collection.Find(bson.M{"firstname": "Jon"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastName)
}
