package main

import (
	//"moneh/factories"
	"context"
	"fmt"
	"log"
	"moneh/packages/database"
	"moneh/routes"

	"firebase.google.com/go/db"
)

func main() {
	// Run App
	database.Init()

	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))

	// Run Seeders
	// factories.Factory()
}

func GetDataFromFirebaseDB(client *db.Client) {
	type UserScore struct {
		Score int `json:"score"`
	}

	ref := client.NewRef("user_scores")

	var s UserScore
	if err := ref.Get(context.TODO(), &s); err != nil {
		log.Fatalln("error in reading from firebase DB: ", err)
	}
	fmt.Println("retrieved user's score is: ", s.Score)
}
func DeleteDataFromFirebaseDB(client *db.Client) {
	ref := client.NewRef("user_scores/1")

	if err := ref.Delete(context.TODO()); err != nil {
		log.Fatalln("error in deleting ref: ", err)
	}
	fmt.Println("user's score deleted successfully:)")
}
