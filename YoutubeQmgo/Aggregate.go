package main

import (
	"context"
	"fmt"
	"qmgoTest/database"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	matchStage := bson.D{{"$match", []bson.E{{"age", bson.D{{"$gt", 20}}}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$name"}, {"email", bson.D{{"$push", "$email"}}}, {"age", bson.D{{"$sum", "$age"}}}}}}
	var showEithInfo []bson.M
	err := database.QmgoConnection.Aggregate(context.TODO(), []bson.D{matchStage, groupStage}).All(&showEithInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", showEithInfo)

	for _, showEithInfo := range showEithInfo {
		fmt.Printf("%+v\n", showEithInfo)
	}
}
