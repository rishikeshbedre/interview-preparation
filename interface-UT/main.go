package main

import (
	"company-svc/lib"
	"company-svc/models"
	"company-svc/operations"
)

func main() {
	var empIntf lib.EmployeeInterface = lib.CreateEmployeeObject()
	var depIntf lib.DepartmentInterface = lib.CreateDepartmentObject()
	var opIntf operations.OperationsExecution = operations.CreateOperationsObject(empIntf, depIntf)

	var request = models.OpRequest{
		EmployeeName:   "rishi",
		EmployeeCity:   "mysore",
		DepartmentName: "platform",
	}

	opIntf.AddEntry(request)
}
