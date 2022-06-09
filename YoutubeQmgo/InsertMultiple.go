package main

// import (
// 	"context"
// 	"fmt"
// 	"qmgoTest/database"
// 	"qmgoTest/models"
// )

// func main() {
// 	userinfo := []models.User{
// 		{
// 			Id:       1,
// 			Name:     "Tom",
// 			Password: "123456",
// 			Age:      18,
// 			Email:    "tom@gmail.com",
// 		},
// 		{
// 			Id:       2,
// 			Name:     "Jerry",
// 			Password: "123456",
// 			Age:      18,
// 			Email:    "jerry@gmail.com",
// 		},
// 		{
// 			Id:       3,
// 			Name:     "Lucy",
// 			Password: "123456",
// 			Age:      18,
// 			Email:    "lucy@gmail.com",
// 		},
// 	}

// 	result, err := database.QmgoConnection.InsertMany(context.TODO(), userinfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", result)
// }
