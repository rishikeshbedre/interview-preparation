package operations

import (
	"company-svc/lib"
	"company-svc/models"
)

type Operations struct {
	EmpObj lib.EmployeeInterface
	DepObj lib.DepartmentInterface
}

type OperationsExecution interface {
	AddEntry(empDetails models.OpRequest) error
}

func (op *Operations) AddEntry(empDetails models.OpRequest) error {
	err := op.EmpObj.AddEmployee(empDetails.EmployeeName, empDetails.EmployeeCity)
	op.DepObj.AddDepartment(empDetails.DepartmentName)
	return err
}

func CreateOperationsObject(empObj lib.EmployeeInterface, depObj lib.DepartmentInterface) *Operations {
	var tempOperation = Operations{
		EmpObj: empObj,
		DepObj: depObj,
	}
	return &tempOperation
}
