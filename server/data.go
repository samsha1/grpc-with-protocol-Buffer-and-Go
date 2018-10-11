package main

import(
	pb "google.golang.org/grpc/examples/grpcdemo/pb"
)

var employee = []pb.Employee {
	pb.Employee{
		Id:1,
		BadgeNumber:2000,
		FirstName:"Samrat",
		LastName:"Shakya",
		VacationAccrualRate: 2,
		VacationAccrued: 30,
	},
	pb.Employee{
		Id:1,
		BadgeNumber:2001,
		FirstName:"Shreeya",
		LastName:"Shakya",
		VacationAccrualRate: 3,
		VacationAccrued: 30.2,
	},
	pb.Employee{
		Id:1,
		BadgeNumber:2002,
		FirstName:"Asmit",
		LastName:"Khadka",
		VacationAccrualRate: 2.4,
		VacationAccrued: 30.1,
	},
}