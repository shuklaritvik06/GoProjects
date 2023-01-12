package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shuklaritvik06/GoProjects/GO_Mongo_API/pkg/models"
	"github.com/shuklaritvik06/GoProjects/GO_Mongo_API/pkg/utils"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	results, _ := models.GetEmployee()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(results)
	w.Write(res)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	results, _ := models.GetEmployeeByID(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(results)
	w.Write(res)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	response, _ := models.GetEmployeeByID(ID)
	employee := &models.Employee{}
	utils.ParseBody(r, employee)
	if employee.Name != "" {
		response.Name = employee.Name
	}
	if employee.Team != "" {
		response.Team = employee.Team
	}
	if employee.Work_Location != "" {
		response.Work_Location = employee.Work_Location
	}
	if employee.Salary != 0 {
		response.Salary = employee.Salary
	}
	response.UpdateEmployee(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal("Updated")
	w.Write(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	results, _ := models.DeleteEmployee(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(results)
	w.Write(res)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := &models.Employee{}
	utils.ParseBody(r, employee)
	results, _ := employee.CreateEmployee()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(results)
	w.Write(res)
}
