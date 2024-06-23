package resolvers

import (
	"context"
	"github.com/les-cours/user-api/api/users"
	"github.com/les-cours/user-api/graph/models"
	gprcToGraph "github.com/les-cours/user-api/grpcToGraph"
	"github.com/les-cours/user-api/permisions"
	"log"
)

func (r *mutationResolver) Init(ctx context.Context) ([]*models.Notification, error) {

	student, err := permisions.Student(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.UserClient.InitStudent(ctx, &users.IDRequest{
		Id: student.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Notifications(res), nil
}

func (r *queryResolver) Students(ctx context.Context, in models.GetStudentsRequest) ([]*models.Student, error) {

	//var user *types.UserToken
	//if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
	//	return nil, ErrPermissionDenied
	//}
	//
	//if !user.Read.USER {
	//	return nil, ErrPermissionDenied
	//}

	var students []*models.Student
	res, err := r.UserClient.GetStudents(ctx, &users.GetStudentsRequest{
		FilterType:  in.FilterType,
		FilterValue: in.FilterValue,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	for _, student := range res.Students {
		students = append(students, gprcToGraph.Student(student))
	}

	return students, nil
}

func (r *queryResolver) Student(ctx context.Context, studentID string) (*models.Student, error) {

	student, err := r.UserClient.GetStudent(ctx, &users.GetStudentRequest{
		StudentID: studentID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.Student(student), nil

}

func (r *mutationResolver) Signup(ctx context.Context, in models.StudentSignupRequest) (*models.SignupResponse, error) {

	res, err := r.UserClient.StudentSignup(ctx, &users.StudentSignupRequest{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Password:  in.Password,
		Dob:       in.Dob,
		Gender:    in.Gender,
		GradID:    in.GradID,
		CityID:    int32(in.CityID),
	})
	if err != nil {
		log.Println("students.go 64", err.Error())
		return nil, ErrApi(err)
	}
	return &models.SignupResponse{
		Succeeded: true,
		AccessToken: &models.AccessToken{
			Token:     res.AccessToken.Token,
			ExpiresAt: int(res.AccessToken.ExpiresAt),
			TokenType: res.AccessToken.TokenType,
		},
	}, nil
}

func (r *mutationResolver) EmailConfirmation(ctx context.Context, in models.EmailConfirmationRequest) (*models.OperationStatus, error) {

	res, err := r.UserClient.EmailConfirmation(ctx, &users.EmailConfirmationRequest{
		AccountID: in.AccountID,
		Code:      int64(in.Code),
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.OperationStatus{
		Succeeded: res.Completed,
	}, nil
}
