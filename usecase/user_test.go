package usecase

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/guntoroyk/go-user-api/entity"
	"github.com/guntoroyk/go-user-api/lib/validator"
	"github.com/guntoroyk/go-user-api/mocks"
	"github.com/guntoroyk/go-user-api/repository"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		userRepo repository.UserRepoItf
	}
	tests := []struct {
		name string
		args args
		want UserUsecaseItf
	}{
		{
			name: "TestNewUserUsecase",
			args: args{
				userRepo: nil,
			},
			want: &userUsecase{
				userRepo:  nil,
				validator: validator.GetValidator(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.userRepo, validator.GetValidator()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUsers(t *testing.T) {
	type fields struct {
		userRepo func(ctrl *gomock.Controller) repository.UserRepoItf
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.User
		wantErr bool
	}{
		{
			name: "success get users",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().GetUsers().Return([]*entity.User{
						{
							ID:        1,
							Name:      "User",
							Username:  "user",
							Phone:     "081234567890",
							Role:      "user",
							CreatedAt: "2021-01-01 00:00:00",
							UpdatedAt: "2021-01-01 00:00:00",
						},
					}, nil)
					return mockUserRepo
				},
			},
			want: []*entity.User{
				{
					ID:        1,
					Name:      "User",
					Username:  "user",
					Phone:     "081234567890",
					Role:      "user",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
			},
			wantErr: false,
		},
		{
			name: "failed get users",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().GetUsers().Return(nil, errors.New("error"))
					return mockUserRepo
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &userUsecase{
				userRepo: tt.fields.userRepo(ctrl),
			}
			got, err := c.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUser(t *testing.T) {
	type fields struct {
		userRepo func(ctrl *gomock.Controller) repository.UserRepoItf
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "success get user",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().GetUser(&entity.UserFilter{
						ID: 1,
					}).Return(&entity.User{
						ID:        1,
						Name:      "User",
						Username:  "user",
						Phone:     "081234567890",
						Role:      "user",
						CreatedAt: "2021-01-01 00:00:00",
						UpdatedAt: "2021-01-01 00:00:00",
					}, nil)
					return mockUserRepo
				},
			},
			args: args{
				id: 1,
			},
			want: &entity.User{
				ID:        1,
				Name:      "User",
				Username:  "user",
				Phone:     "081234567890",
				Role:      "user",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			wantErr: false,
		},
		{
			name: "failed get user",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().GetUser(&entity.UserFilter{
						ID: 1,
					}).Return(nil, errors.New("error"))
					return mockUserRepo
				},
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed get user not found",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().GetUser(&entity.UserFilter{
						ID: 1,
					}).Return(nil, sql.ErrNoRows)
					return mockUserRepo
				},
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &userUsecase{
				userRepo: tt.fields.userRepo(ctrl),
			}
			got, err := c.GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_CreateUser(t *testing.T) {
	type fields struct {
		userRepo func(ctrl *gomock.Controller) repository.UserRepoItf
	}
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "success create user",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().GetUser(&entity.UserFilter{Username: "user"}).Return(nil, sql.ErrNoRows)
					mockUserRepo.EXPECT().CreateUser(&entity.User{
						Name:     "User",
						Username: "user",
						Phone:    "081234567890",
						Password: "1234",
						Role:     "user",
					}).Return(&entity.User{
						ID:        1,
						Name:      "User",
						Username:  "user",
						Phone:     "081234567890",
						Password:  "1234",
						Role:      "user",
						CreatedAt: "2021-01-01 00:00:00",
						UpdatedAt: "2021-01-01 00:00:00",
					}, nil)
					return mockUserRepo
				},
			},
			args: args{
				user: &entity.User{
					Name:     "User",
					Username: "user",
					Phone:    "081234567890",
					Password: "1234",
					Role:     "user",
				},
			},
			want: &entity.User{
				ID:        1,
				Name:      "User",
				Username:  "user",
				Phone:     "081234567890",
				Password:  "1234",
				Role:      "user",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
		},
		{
			name: "failed create user error validation",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					return mockUserRepo
				},
			},
			args: args{
				user: &entity.User{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed create user",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().GetUser(&entity.UserFilter{Username: "user"}).Return(nil, sql.ErrNoRows)
					mockUserRepo.EXPECT().CreateUser(&entity.User{
						Name:     "User",
						Username: "user",
						Phone:    "081234567890",
						Password: "1234",
						Role:     "user",
					}).Return(nil, errors.New("error"))
					return mockUserRepo
				},
			},
			args: args{
				user: &entity.User{
					Name:     "User",
					Username: "user",
					Password: "1234",
					Phone:    "081234567890",
					Role:     "user",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &userUsecase{
				userRepo:  tt.fields.userRepo(ctrl),
				validator: validator.GetValidator(),
			}
			got, err := c.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_UpdateUser(t *testing.T) {
	type fields struct {
		userRepo func(ctrl *gomock.Controller) repository.UserRepoItf
	}
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "success update user",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().UpdateUser(&entity.User{
						ID:       1,
						Name:     "User",
						Username: "user",
						Password: "1234",
						Phone:    "081234567890",
						Role:     "user",
					}).Return(&entity.User{
						ID:        1,
						Name:      "User",
						Username:  "user",
						Phone:     "081234567890",
						Password:  "1234",
						Role:      "user",
						CreatedAt: "2021-01-01 00:00:00",
						UpdatedAt: "2021-01-01 00:00:00",
					}, nil)
					return mockUserRepo
				},
			},
			args: args{
				user: &entity.User{
					ID:       1,
					Name:     "User",
					Username: "user",
					Password: "1234",
					Phone:    "081234567890",
					Role:     "user",
				},
			},
			want: &entity.User{
				ID:        1,
				Name:      "User",
				Username:  "user",
				Phone:     "081234567890",
				Password:  "1234",
				Role:      "user",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
		},
		{
			name: "failed create user error validation",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					return mockUserRepo
				},
			},
			args: args{
				user: &entity.User{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &userUsecase{
				userRepo:  tt.fields.userRepo(ctrl),
				validator: validator.GetValidator(),
			}
			got, err := c.UpdateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_DeleteUser(t *testing.T) {
	type fields struct {
		userRepo func(ctrl *gomock.Controller) repository.UserRepoItf
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success delete user",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().DeleteUser(1).Return(nil)
					return mockUserRepo
				},
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "error delete not found",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) repository.UserRepoItf {
					mockUserRepo := mocks.NewMockUserRepoItf(ctrl)
					mockUserRepo.EXPECT().DeleteUser(1).Return(sql.ErrNoRows)
					return mockUserRepo
				},
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := &userUsecase{
				userRepo: tt.fields.userRepo(ctrl),
			}
			if err := c.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
