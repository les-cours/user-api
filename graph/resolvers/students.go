package resolvers

import (
	"context"
	"github.com/les-cours/user-api/api/users"
	"github.com/les-cours/user-api/graph/models"
	gprcToGraph "github.com/les-cours/user-api/grpcToGraph"
)

func (r *queryResolver) Students(ctx context.Context, in models.GetStudentsRequest) ([]*models.Student, error) {

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

	//ctx.Value()
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
		return nil, ErrApi(err)
	}
	return &models.SignupResponse{
		Succeeded: true,
		AccessToken: &models.AccessToken{
			Token:     res.AccessToken.Token,
			ExpiresAt: int(res.AccessToken.ExpiresAt),
			TokenType: res.AccessToken.TokenType,
		},
		RefreshToken: &models.RefreshToken{
			Token:     res.RefreshToken.Token,
			ExpiresAt: int(res.RefreshToken.ExpiresAt),
		},
		SignupToken: &models.SignupToken{
			Token:     res.SignupToken.Token,
			ExpiresAt: int(res.SignupToken.ExpiresAt),
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
