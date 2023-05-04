package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DBNAME           = "student"
	COLLNAME         = "student"
)

var db *mongo.Database

// Connect establishes a connection to the database
func init() {
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

// Student
type Student struct {
	ID      string `bson:"id"`
	Name    string `bson:"name"`
	Email   string `bson:"email"`
	Address struct {
		City    string `bson:"city"`
		Zipcode string `bson:"zipcode"`
		Phone   string `bson:"phone"`
	} `bson:"address"`
}

// CreateStudent add a new student data
func CreateStudent(student Student) {
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), student)
	if err != nil {
		log.Fatal(err)
	}
}

// GetStudent
func GetStudent(studentID string) Student {
	var student Student
	filter := bson.D{{"id", studentID}}
	err := db.Collection(COLLNAME).FindOne(context.Background(), filter).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}
	return student
}

// UpdateStudent
func UpdateStudent(student Student, studentID string) {
	filter := bson.D{{"id", studentID}}
	update := bson.D{
		{"$set", bson.D{
			{"name", student.Name},
			{"email", student.Email},
			{"address.city", student.Address.City},
			{"address.zipcode", student.Address.Zipcode},
			{"address.phone", student.Address.Phone},
		}},
	}
	_, err := db.Collection(COLLNAME).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteStudent
func DeleteStudent(studentID string) {
	filter := bson.D{{"id", studentID}}
	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	student := Student{
		ID:    "1",
		Name:  "Elon Musk",
		Email: "elonmusk@gmail.com",
		Address: struct {
			City    string `bson:"city"`
			Zipcode string `bson:"zipcode"`
			Phone   string `bson:"phone"`
		}{
			City:    "New York",
			Zipcode: "10001",
			Phone:   "555-555-1212",
		},
	}

	// create a student
	CreateStudent(student)

	// get a student
	student = GetStudent("1")
	fmt.Println(student)

	// update a student
	student.Name = "Melon Musk"
	UpdateStudent(student, "1")

	// delete a student
	DeleteStudent("1")
}
