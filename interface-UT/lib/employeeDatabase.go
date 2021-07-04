package lib

import (
	"sync"
	"log"

	"company-svc/models"
)

type EmployeeDatabase struct {
	mu    sync.RWMutex
	EmpDB models.Employee
}

type EmployeeInterface interface {
	AddEmployee(name string, city string) error
}

func (emp *EmployeeDatabase) AddEmployee(name string, city string) error {
	emp.mu.Lock()
	emp.EmpDB.Name = name
	emp.EmpDB.City = city
	log.Println(emp.EmpDB)
	emp.mu.Unlock()
	return nil
}

func CreateEmployeeObject() *EmployeeDatabase {
	var tempEmployeeDatabase = EmployeeDatabase{
		mu:    sync.RWMutex{},
		EmpDB: models.Employee{},
	}
	return &tempEmployeeDatabase
}
