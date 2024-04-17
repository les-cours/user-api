package resolvers

import (
	"context"
	"github.com/les-cours/user-api/api/users"
	"github.com/les-cours/user-api/graph/models"
)

func (r *queryResolver) GetStudents(ctx context.Context) ([]*models.Student, error) {

	var students []*models.Student
	res, err := r.UserClient.GetStudent(ctx, &users.GetStudentRequest{
		StudentID: "test",
	})
	if err != nil {
		return nil, err
	}

	students = append(students, &models.Student{
		ID:   res.Id,
		Name: res.Firstname,
	})
	return students, nil
}

func (r *mutationResolver) Signup(ctx context.Context, in models.SignupRequest) (*models.SignupResponse, error) {

	//ctx.Value()
	res, err := r.UserClient.StudentSignup(ctx, &users.StudentSignupRequest{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Password:  in.Password,
	})
	if err != nil {
		return nil, err
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
