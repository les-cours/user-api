package gprcToGraph

import (
	"github.com/les-cours/user-api/api/users"
	"github.com/les-cours/user-api/graph/models"
)

func Student(student *users.Student) *models.Student {
	return &models.Student{
		StudentID:          student.StudentId,
		Username:           student.Username,
		Firstname:          student.Firstname,
		Lastname:           student.Lastname,
		Gender:             student.Gender,
		DateOfBirth:        student.DateOfBirth,
		Status:             student.Status,
		Avatar:             &student.Avatar,
		NotificationStatus: student.NotificationStatus,
		OnlineStatus:       student.OnlineStatus,
		DefaultAvatar:      student.StudentId,
		CityID:             student.StudentId,
	}

}
