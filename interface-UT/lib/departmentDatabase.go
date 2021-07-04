package lib

import (
	"sync"
	"log"

	"company-svc/models"
)

type DepartmentDatabase struct {
	mu    sync.RWMutex
	DepDB models.Department
}

type DepartmentInterface interface {
	AddDepartment(name string) error
}

func (dep *DepartmentDatabase) AddDepartment(name string) error {
	dep.mu.Lock()
	dep.DepDB.Name = name
	log.Println(dep.DepDB)
	dep.mu.Unlock()
	return nil
}

func CreateDepartmentObject() *DepartmentDatabase {
	var tempDepartmentDatabase = DepartmentDatabase{
		mu:    sync.RWMutex{},
		DepDB: models.Department{},
	}
	return &tempDepartmentDatabase
}
