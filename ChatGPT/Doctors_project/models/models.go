package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct{
	DoctorId primitive.ObjectID 		`json:"doctorid,omitempty" bson:"_doctorid,omitempty"`
	DoctorName string 					`json:"doctorname,omitempty" bson:"doctorname,omitempty"`
	DoctorSpecialist string 			`json:"doctorspecialist,omitempty" bson:"doctorspecialist,omitempty"`
	DoctorQualification string 			`json:"doctorqualification,omitempty" bson:"doctorqualification,omitempty"`
	DoctorSalary string					`json:"doctorsalary,omitempty" bson:"doctorsalary,omitempty"`
}