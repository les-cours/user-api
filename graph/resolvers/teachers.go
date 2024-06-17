package resolvers

import (
	"context"
	"github.com/les-cours/user-api/api/users"
	"github.com/les-cours/user-api/graph/models"
	gprcToGraph "github.com/les-cours/user-api/grpcToGraph"
	"github.com/les-cours/user-api/permisions"
)

func (r *mutationResolver) InviteTeacher(ctx context.Context, in models.InviteTeacherRequest) (*models.OperationStatus, error) {

	//ctx.Value()
	_, err := r.UserClient.InviteTeacher(ctx, &users.InviteTeacherRequest{
		Email:    in.Email,
		Subjects: in.Subjects,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}

func (r *mutationResolver) SignupTeacher(ctx context.Context, in models.TeacherSignupRequest) (*models.SignupResponse, error) {

	//ctx.Value()
	res, err := r.UserClient.TeacherSignup(ctx, &users.TeacherSignupRequest{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Password:  in.Password,
		Dob:       in.Dob,
		Gender:    in.Gender,
		CityID:    int32(in.CityID),
		TeacherID: in.TeacherID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.SignupResponse{
		Succeeded: true,
		AccessToken: &models.AccessToken{
			Token: res.Token,
		},
	}, nil
}

func (r *mutationResolver) UpdateTeacher(ctx context.Context, in models.UpdateTeacherRequest) (*models.Teacher, error) {

	//user, _ := ctx.Value("user").(*types.UserToken)
	//if !user.Update.USER {
	//	return nil, ErrPermissionDenied
	//}

	_, err := r.UserClient.UpdateTeacher(ctx, &users.UpdateTeacherRequest{
		TeacherID:   in.TeacherID,
		CityID:      int32(in.CityID),
		Firstname:   in.Firstname,
		Lastname:    in.Lastname,
		Gender:      in.Gender,
		DateOfBirth: in.DateOfBirth,
		Description: in.Description,
		Avatar:      in.Avatar,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.Teacher{
		TeacherID:   in.TeacherID,
		CityID:      in.CityID,
		Firstname:   in.Firstname,
		Lastname:    in.Lastname,
		Gender:      in.Gender,
		DateOfBirth: in.DateOfBirth,
		Description: in.Description,
		Avatar:      in.Avatar,
	}, nil

}

func (r *queryResolver) GetTeachers(ctx context.Context) ([]*models.Teacher, error) {

	_, err := permisions.Admin(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.UserClient.GetTeachers(ctx, &users.Empty{})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Teachers(res), nil

}

func (r *queryResolver) GetTeacher(ctx context.Context, teacherID string) (*models.Teacher, error) {

	_, permissionErr := permisions.Admin(ctx)
	if permissionErr != nil {
		return nil, permissionErr
	}
	res, err := r.UserClient.GetTeacher(ctx, &users.IDRequest{
		Id: teacherID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.Teacher(res), nil

}
