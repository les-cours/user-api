// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/users/user.proto

package users

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetStudents(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Students, error)
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Student, error)
	AddStudent(ctx context.Context, in *StudentAddRequest, opts ...grpc.CallOption) (*Student, error)
	UpdateStudentStatus(ctx context.Context, in *UpdateStudentStatusRequest, opts ...grpc.CallOption) (*UpdateStudentStatusResponse, error)
	UpdateStudent(ctx context.Context, in *StudentUpdateRequest, opts ...grpc.CallOption) (*Student, error)
	DeleteStudent(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	DeleteStudents(ctx context.Context, in *MultiStudentsDeleteRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	UpdateStudentOnlineStatus(ctx context.Context, in *UpdateStudentOnlineStatusRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	StudentSignup(ctx context.Context, in *StudentSignupRequest, opts ...grpc.CallOption) (*StudentSignupResponse, error)
	EmailConfirmation(ctx context.Context, in *EmailConfirmationRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	InviteTeacher(ctx context.Context, in *InviteTeacherRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	TeacherSignup(ctx context.Context, in *TeacherSignupRequest, opts ...grpc.CallOption) (*TeacherSignupResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*User, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	PasswordResetLinkHealth(ctx context.Context, in *PasswordResetLinkHealthRequest, opts ...grpc.CallOption) (*PasswordResetLinkHealthResponse, error)
	ChangePasswordByLink(ctx context.Context, in *PasswordChangeByLinkRequest, opts ...grpc.CallOption) (*PasswordChangeByLinkResponse, error)
	DoesEmailExist(ctx context.Context, in *DoesEmailExistRequest, opts ...grpc.CallOption) (*DoesEmailExistResponse, error)
	DoesUserNameExist(ctx context.Context, in *DoesUserNameExistRequest, opts ...grpc.CallOption) (*DoesUserNameExistResponse, error)
	IsSignupLinkValid(ctx context.Context, in *IsSignupLinkValidRequest, opts ...grpc.CallOption) (*IsSignupLinkValidResponse, error)
	UserPasswordReset(ctx context.Context, in *UserPasswordResetRequest, opts ...grpc.CallOption) (*OperationStatus, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetStudents(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Students, error) {
	out := new(Students)
	err := c.cc.Invoke(ctx, "/users.UserService/GetStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/users.UserService/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddStudent(ctx context.Context, in *StudentAddRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/users.UserService/AddStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateStudentStatus(ctx context.Context, in *UpdateStudentStatusRequest, opts ...grpc.CallOption) (*UpdateStudentStatusResponse, error) {
	out := new(UpdateStudentStatusResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/UpdateStudentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateStudent(ctx context.Context, in *StudentUpdateRequest, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/users.UserService/UpdateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteStudent(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteStudents(ctx context.Context, in *MultiStudentsDeleteRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateStudentOnlineStatus(ctx context.Context, in *UpdateStudentOnlineStatusRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/users.UserService/UpdateStudentOnlineStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) StudentSignup(ctx context.Context, in *StudentSignupRequest, opts ...grpc.CallOption) (*StudentSignupResponse, error) {
	out := new(StudentSignupResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/StudentSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EmailConfirmation(ctx context.Context, in *EmailConfirmationRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/users.UserService/EmailConfirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) InviteTeacher(ctx context.Context, in *InviteTeacherRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/users.UserService/InviteTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) TeacherSignup(ctx context.Context, in *TeacherSignupRequest, opts ...grpc.CallOption) (*TeacherSignupResponse, error) {
	out := new(TeacherSignupResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/TeacherSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PasswordResetLinkHealth(ctx context.Context, in *PasswordResetLinkHealthRequest, opts ...grpc.CallOption) (*PasswordResetLinkHealthResponse, error) {
	out := new(PasswordResetLinkHealthResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/PasswordResetLinkHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePasswordByLink(ctx context.Context, in *PasswordChangeByLinkRequest, opts ...grpc.CallOption) (*PasswordChangeByLinkResponse, error) {
	out := new(PasswordChangeByLinkResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/ChangePasswordByLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DoesEmailExist(ctx context.Context, in *DoesEmailExistRequest, opts ...grpc.CallOption) (*DoesEmailExistResponse, error) {
	out := new(DoesEmailExistResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/DoesEmailExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DoesUserNameExist(ctx context.Context, in *DoesUserNameExistRequest, opts ...grpc.CallOption) (*DoesUserNameExistResponse, error) {
	out := new(DoesUserNameExistResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/DoesUserNameExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsSignupLinkValid(ctx context.Context, in *IsSignupLinkValidRequest, opts ...grpc.CallOption) (*IsSignupLinkValidResponse, error) {
	out := new(IsSignupLinkValidResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/IsSignupLinkValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserPasswordReset(ctx context.Context, in *UserPasswordResetRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/users.UserService/UserPasswordReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetStudents(context.Context, *GetStudentRequest) (*Students, error)
	GetStudent(context.Context, *GetStudentRequest) (*Student, error)
	AddStudent(context.Context, *StudentAddRequest) (*Student, error)
	UpdateStudentStatus(context.Context, *UpdateStudentStatusRequest) (*UpdateStudentStatusResponse, error)
	UpdateStudent(context.Context, *StudentUpdateRequest) (*Student, error)
	DeleteStudent(context.Context, *DeleteStudentRequest) (*OperationStatus, error)
	DeleteStudents(context.Context, *MultiStudentsDeleteRequest) (*OperationStatus, error)
	UpdateStudentOnlineStatus(context.Context, *UpdateStudentOnlineStatusRequest) (*OperationStatus, error)
	StudentSignup(context.Context, *StudentSignupRequest) (*StudentSignupResponse, error)
	EmailConfirmation(context.Context, *EmailConfirmationRequest) (*OperationStatus, error)
	InviteTeacher(context.Context, *InviteTeacherRequest) (*OperationStatus, error)
	TeacherSignup(context.Context, *TeacherSignupRequest) (*TeacherSignupResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	GetUserByID(context.Context, *GetUserByIDRequest) (*User, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	PasswordResetLinkHealth(context.Context, *PasswordResetLinkHealthRequest) (*PasswordResetLinkHealthResponse, error)
	ChangePasswordByLink(context.Context, *PasswordChangeByLinkRequest) (*PasswordChangeByLinkResponse, error)
	DoesEmailExist(context.Context, *DoesEmailExistRequest) (*DoesEmailExistResponse, error)
	DoesUserNameExist(context.Context, *DoesUserNameExistRequest) (*DoesUserNameExistResponse, error)
	IsSignupLinkValid(context.Context, *IsSignupLinkValidRequest) (*IsSignupLinkValidResponse, error)
	UserPasswordReset(context.Context, *UserPasswordResetRequest) (*OperationStatus, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetStudents(context.Context, *GetStudentRequest) (*Students, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}
func (UnimplementedUserServiceServer) GetStudent(context.Context, *GetStudentRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedUserServiceServer) AddStudent(context.Context, *StudentAddRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (UnimplementedUserServiceServer) UpdateStudentStatus(context.Context, *UpdateStudentStatusRequest) (*UpdateStudentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudentStatus not implemented")
}
func (UnimplementedUserServiceServer) UpdateStudent(context.Context, *StudentUpdateRequest) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedUserServiceServer) DeleteStudent(context.Context, *DeleteStudentRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedUserServiceServer) DeleteStudents(context.Context, *MultiStudentsDeleteRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudents not implemented")
}
func (UnimplementedUserServiceServer) UpdateStudentOnlineStatus(context.Context, *UpdateStudentOnlineStatusRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudentOnlineStatus not implemented")
}
func (UnimplementedUserServiceServer) StudentSignup(context.Context, *StudentSignupRequest) (*StudentSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StudentSignup not implemented")
}
func (UnimplementedUserServiceServer) EmailConfirmation(context.Context, *EmailConfirmationRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailConfirmation not implemented")
}
func (UnimplementedUserServiceServer) InviteTeacher(context.Context, *InviteTeacherRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteTeacher not implemented")
}
func (UnimplementedUserServiceServer) TeacherSignup(context.Context, *TeacherSignupRequest) (*TeacherSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeacherSignup not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByID(context.Context, *GetUserByIDRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) PasswordResetLinkHealth(context.Context, *PasswordResetLinkHealthRequest) (*PasswordResetLinkHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordResetLinkHealth not implemented")
}
func (UnimplementedUserServiceServer) ChangePasswordByLink(context.Context, *PasswordChangeByLinkRequest) (*PasswordChangeByLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePasswordByLink not implemented")
}
func (UnimplementedUserServiceServer) DoesEmailExist(context.Context, *DoesEmailExistRequest) (*DoesEmailExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoesEmailExist not implemented")
}
func (UnimplementedUserServiceServer) DoesUserNameExist(context.Context, *DoesUserNameExistRequest) (*DoesUserNameExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoesUserNameExist not implemented")
}
func (UnimplementedUserServiceServer) IsSignupLinkValid(context.Context, *IsSignupLinkValidRequest) (*IsSignupLinkValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSignupLinkValid not implemented")
}
func (UnimplementedUserServiceServer) UserPasswordReset(context.Context, *UserPasswordResetRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPasswordReset not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetStudents(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddStudent(ctx, req.(*StudentAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateStudentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateStudentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/UpdateStudentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateStudentStatus(ctx, req.(*UpdateStudentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/UpdateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateStudent(ctx, req.(*StudentUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteStudent(ctx, req.(*DeleteStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiStudentsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteStudents(ctx, req.(*MultiStudentsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateStudentOnlineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentOnlineStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateStudentOnlineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/UpdateStudentOnlineStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateStudentOnlineStatus(ctx, req.(*UpdateStudentOnlineStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_StudentSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).StudentSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/StudentSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).StudentSignup(ctx, req.(*StudentSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EmailConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EmailConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/EmailConfirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EmailConfirmation(ctx, req.(*EmailConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_InviteTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).InviteTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/InviteTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).InviteTeacher(ctx, req.(*InviteTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_TeacherSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).TeacherSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/TeacherSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).TeacherSignup(ctx, req.(*TeacherSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByID(ctx, req.(*GetUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PasswordResetLinkHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordResetLinkHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PasswordResetLinkHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/PasswordResetLinkHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PasswordResetLinkHealth(ctx, req.(*PasswordResetLinkHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePasswordByLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordChangeByLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePasswordByLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/ChangePasswordByLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePasswordByLink(ctx, req.(*PasswordChangeByLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DoesEmailExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoesEmailExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DoesEmailExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DoesEmailExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DoesEmailExist(ctx, req.(*DoesEmailExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DoesUserNameExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoesUserNameExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DoesUserNameExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DoesUserNameExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DoesUserNameExist(ctx, req.(*DoesUserNameExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsSignupLinkValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsSignupLinkValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsSignupLinkValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/IsSignupLinkValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsSignupLinkValid(ctx, req.(*IsSignupLinkValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserPasswordReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPasswordResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserPasswordReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/UserPasswordReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserPasswordReset(ctx, req.(*UserPasswordResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudents",
			Handler:    _UserService_GetStudents_Handler,
		},
		{
			MethodName: "GetStudent",
			Handler:    _UserService_GetStudent_Handler,
		},
		{
			MethodName: "AddStudent",
			Handler:    _UserService_AddStudent_Handler,
		},
		{
			MethodName: "UpdateStudentStatus",
			Handler:    _UserService_UpdateStudentStatus_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _UserService_UpdateStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _UserService_DeleteStudent_Handler,
		},
		{
			MethodName: "DeleteStudents",
			Handler:    _UserService_DeleteStudents_Handler,
		},
		{
			MethodName: "UpdateStudentOnlineStatus",
			Handler:    _UserService_UpdateStudentOnlineStatus_Handler,
		},
		{
			MethodName: "StudentSignup",
			Handler:    _UserService_StudentSignup_Handler,
		},
		{
			MethodName: "EmailConfirmation",
			Handler:    _UserService_EmailConfirmation_Handler,
		},
		{
			MethodName: "InviteTeacher",
			Handler:    _UserService_InviteTeacher_Handler,
		},
		{
			MethodName: "TeacherSignup",
			Handler:    _UserService_TeacherSignup_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserService_GetUserByID_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "PasswordResetLinkHealth",
			Handler:    _UserService_PasswordResetLinkHealth_Handler,
		},
		{
			MethodName: "ChangePasswordByLink",
			Handler:    _UserService_ChangePasswordByLink_Handler,
		},
		{
			MethodName: "DoesEmailExist",
			Handler:    _UserService_DoesEmailExist_Handler,
		},
		{
			MethodName: "DoesUserNameExist",
			Handler:    _UserService_DoesUserNameExist_Handler,
		},
		{
			MethodName: "IsSignupLinkValid",
			Handler:    _UserService_IsSignupLinkValid_Handler,
		},
		{
			MethodName: "UserPasswordReset",
			Handler:    _UserService_UserPasswordReset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/users/user.proto",
}
