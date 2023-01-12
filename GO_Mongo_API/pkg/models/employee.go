package models

import (
	"context"

	"github.com/shuklaritvik06/GoProjects/GO_Mongo_API/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

type Employee struct {
	Name          string      `bson:"name" json:"name"`
	Employee_ID   string      `bson:"employee_id" json:"employee_id"`
	Salary        int64       `bson:"salary" json:"salary"`
	Address       interface{} `bson:"address" json:"address"`
	Team          string      `bson:"team" json:"team"`
	Work_Location string      `bson:"work_location" json:"work_location"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetEmployee() ([]Employee, error) {
	var results []Employee
	cur, err := db.Database("Employee").Collection("employee").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results, nil
}

func GetEmployeeByID(id string) (Employee, error) {
	var employee Employee
	err := db.Database("Employee").Collection("employee").FindOne(context.TODO(), bson.D{
		primitive.E{Key: "employee_id", Value: id},
	}).Decode(&employee)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func DeleteEmployee(id string) (*mongo.DeleteResult, error) {
	result, err := db.Database("Employee").Collection("employee").DeleteOne(context.TODO(), bson.D{
		primitive.E{Key: "employee_id", Value: id},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (employee *Employee) CreateEmployee() (*mongo.InsertOneResult, error) {
	result, err := db.Database("Employee").Collection("employee").InsertOne(context.TODO(), employee)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (employee *Employee) UpdateEmployee(id string) {
	db.Database("Employee").Collection("employee").FindOneAndUpdate(context.TODO(), bson.D{
		primitive.E{Key: "employee_id", Value: id},
	}, bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: employee.Name},
			primitive.E{Key: "salary", Value: employee.Salary},
			primitive.E{Key: "address", Value: employee.Address},
			primitive.E{Key: "team", Value: employee.Team},
			primitive.E{Key: "work_location", Value: employee.Work_Location},
		}},
	})
}
