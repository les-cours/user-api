package gprcToGraph

import (
	"github.com/les-cours/user-api/api/users"
	"github.com/les-cours/user-api/graph/models"
)

func Student(student *users.Student) *models.Student {
	return &models.Student{
		StudentID:          student.StudentID,
		Username:           student.Username,
		Firstname:          student.Firstname,
		Lastname:           student.Lastname,
		Gender:             student.Gender,
		DateOfBirth:        student.DateOfBirth,
		Status:             student.Status,
		Avatar:             &student.Avatar,
		NotificationStatus: student.NotificationStatus,
		OnlineStatus:       student.OnlineStatus,
		DefaultAvatar:      student.Avatar,
		CityID:             student.CityID,
	}

}

func Teachers(teachersApi *users.Teachers) []*models.Teacher {
	var teachers = make([]*models.Teacher, 0)
	for _, teacher := range teachersApi.Teachers {
		teachers = append(teachers, Teacher(teacher))
	}
	return teachers
}

func Teacher(teacher *users.Teacher) *models.Teacher {
	return &models.Teacher{
		TeacherID:    teacher.TeacherID,
		CityID:       int(teacher.CityID),
		Firstname:    teacher.Firstname,
		Lastname:     teacher.Lastname,
		Gender:       teacher.Gender,
		DateOfBirth:  teacher.DateOfBirth,
		Description:  teacher.Description,
		Avatar:       teacher.Avatar,
		Email:        teacher.Email,
		OnlineStatus: teacher.OnlineStatus,
		Username:     teacher.Username,
	}
}

func Notifications(notificationApi *users.Notifications) []*models.Notification {
	var notifications = make([]*models.Notification, 0)
	for _, notification := range notificationApi.Notifications {
		notifications = append(notifications, Notification(notification))
	}
	return notifications
}
func Notification(notification *users.Notification) *models.Notification {
	return &models.Notification{
		ID:      notification.Id,
		Title:   notification.Title,
		Content: notification.Content,
	}
}
