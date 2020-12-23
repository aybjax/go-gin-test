package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Json struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// type

func main() {
	fmt.Println("Go Redis Tutorial")

	/*
	** setup
	 */
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	handleError(err)
	fmt.Println(pong)

	/*
	** set
	 */
	// we can call set with a `Key` and a `Value`.
	err = client.Set("name", "Elliot", 0).Err()
	// if there has been an error setting the value
	// handle the error
	handleError(err)

	/**
	** get
	 */

	val, err := client.Get("name").Result()
	handleError(err)

	fmt.Println("name =>", val)

	/**
	** complex
	**/

	json, err := json.Marshal(Json{Name: "Elliot", Age: 25})
	handleError(err)

	err = client.Set("id1234", json, 0).Err()
	handleError(err)
	val, err = client.Get("id1234").Result()
	handleError(err)
	fmt.Println("complex =>", val)

}

func handleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
