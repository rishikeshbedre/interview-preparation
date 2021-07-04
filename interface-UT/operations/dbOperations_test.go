package operations

import (
	"log"
	"errors"
	"company-svc/lib"
	"company-svc/models"
	"reflect"
	"testing"
)

type mockEmployeeDatabase struct {
	EmpDB models.Employee
}

var mockAddEmployee = func(name string, city string) error {
	return nil
}

func (mockEmp *mockEmployeeDatabase) AddEmployee(name string, city string) error {
	return mockAddEmployee(name, city)
}

func TestOperations_AddEntry(t *testing.T) {
	type fields struct {
		EmpObj lib.EmployeeInterface
		DepObj lib.DepartmentInterface
	}
	type args struct {
		empDetails models.OpRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFunc func()
		wantErr bool
	}{
		{
			name: "fixed_mock",
			fields: fields{
				EmpObj: &mockEmployeeDatabase{},
				DepObj: lib.CreateDepartmentObject(),
			},
			args: args{
				empDetails: models.OpRequest{
					EmployeeName: "rishi",
					EmployeeCity: "mysore",
					DepartmentName: "platform",
				},
			},
			mockFunc: func() {
				mockAddEmployee = func(name, city string) error {
					log.Println("name: arup, city: kolkata")
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "error_case",
			fields: fields{
				EmpObj: &mockEmployeeDatabase{},
				DepObj: lib.CreateDepartmentObject(),
			},
			args: args{
				empDetails: models.OpRequest{
					EmployeeName: "rishi",
					EmployeeCity: "mysore",
					DepartmentName: "platform",
				},
			},
			mockFunc: func() {
				mockAddEmployee = func(name, city string) error {
					log.Println("employee already present")
					return errors.New("employee already present")
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &Operations{
				EmpObj: tt.fields.EmpObj,
				DepObj: tt.fields.DepObj,
			}
			tt.mockFunc()
			if err := op.AddEntry(tt.args.empDetails); (err != nil) != tt.wantErr {
				t.Errorf("Operations.AddEntry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOperationsObject(t *testing.T) {
	type args struct {
		empObj lib.EmployeeInterface
		depObj lib.DepartmentInterface
	}
	tests := []struct {
		name string
		args args
		want *Operations
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateOperationsObject(tt.args.empObj, tt.args.depObj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOperationsObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
