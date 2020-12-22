package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person test
type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	/*
	** connection
	 */
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		// log.Fatal(err)
		panic(err.Error())
	}
	fmt.Println("Connected to MongoDB!")

	/*
	** insert
	 */
	collection := client.Database("aybjax").Collection("persons")

	ruan := Person{"my - struct", 34, "my - struct"}
	// ruan := bson.M{"name": "bsonM", "age": 34, "city": "bsonM"}
	// ruan := bson.D{{"name", "bsonD"}, {"age", 34}, {"city", "bsonD"}}

	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	fmt.Printf("====>>>> %#v\n", insertResult)

	/*
	** insert many
	 */
	collection = client.Database("aybjax").Collection("persons")

	james := Person{"James", 32, "Nairobi"}
	frankie := Person{"Frankie", 31, "Nairobi"}

	// map of only this type
	trainers := []interface{}{james, frankie}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	/*
	** update
	 */
	filter := bson.D{}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	/*
	** reading single
	 */
	filter = bson.D{}
	// filter = struct{}{}
	// var result Person
	// var result bson.D
	var result bson.M

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	/*
	** read many
	 */
	findOptions := options.Find()
	findOptions.SetLimit(2)
	var results []*Person
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Person
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
		fmt.Printf("\tvals: %#v\n", elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	/*
	** delete
	 */
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	/*
	** disconnection
	 */
	err = collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Collection delete?.")
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}

}
