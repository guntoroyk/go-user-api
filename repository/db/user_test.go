package db

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guntoroyk/go-user-api/entity"
)

func Test_userRepo_GetUsers(t *testing.T) {
	columns := []string{"id", "name", "username", "password", "phone", "role", "created_at", "updated_at"}

	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.User
		wantErr bool
	}{
		{
			name: "success GetUsers",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectQuery(selectUsersQuery).WillReturnRows(
						sqlmock.NewRows(columns).
							AddRow(1, "User", "user", "password", "08123456789", "admin", "2006-01-01 00:00:00", "2006-01-01 00:00:00"),
					)

					return db
				}(),
			},
			want: []*entity.User{
				{
					ID:        1,
					Name:      "User",
					Username:  "user",
					Password:  "password",
					Phone:     "08123456789",
					Role:      "admin",
					CreatedAt: "2006-01-01 00:00:00",
					UpdatedAt: "2006-01-01 00:00:00",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userRepo{
				db: tt.fields.db,
			}
			got, err := c.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepo.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_GetUser(t *testing.T) {
	columns := []string{"id", "name", "username", "password", "phone", "role", "created_at", "updated_at"}

	type fields struct {
		db *sql.DB
	}
	type args struct {
		filter *entity.UserFilter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "success GetUser",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectQuery("SELECT * FROM users WHERE id= ?").
						WithArgs(1).
						WillReturnRows(
							sqlmock.NewRows(columns).
								AddRow(1, "User", "user", "password", "08123456789", "admin", "2006-01-01 00:00:00", "2006-01-01 00:00:00"),
						)

					return db
				}(),
			},
			args: args{
				filter: &entity.UserFilter{
					ID: 1,
				},
			},
			want: &entity.User{
				ID:        1,
				Name:      "User",
				Username:  "user",
				Password:  "password",
				Phone:     "08123456789",
				Role:      "admin",
				CreatedAt: "2006-01-01 00:00:00",
				UpdatedAt: "2006-01-01 00:00:00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userRepo{
				db: tt.fields.db,
			}
			got, err := c.GetUser(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_CreateUser(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			name: "success CreateUser",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectExec("INSERT INTO users").
						WithArgs("User", "user", "password", "08123456789", "admin", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")).
						WillReturnResult(sqlmock.NewResult(1, 1))

					return db
				}(),
			},
			args: args{
				user: &entity.User{
					Name:     "User",
					Username: "user",
					Password: "password",
					Phone:    "08123456789",
					Role:     "admin",
				},
			},
			want: &entity.User{
				ID:        1,
				Name:      "User",
				Username:  "user",
				Password:  "password",
				Phone:     "08123456789",
				Role:      "admin",
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
				UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userRepo{
				db: tt.fields.db,
			}
			got, err := c.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepo.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_UpdateUser(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			name: "success UpdateUser",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectExec("UPDATE users").
						WithArgs("User", "user", "password", "08123456789", "admin", time.Now().Format("2006-01-02 15:04:05")).
						WillReturnResult(sqlmock.NewResult(1, 1))

					return db
				}(),
			},
			args: args{
				user: &entity.User{
					ID:       1,
					Name:     "User",
					Username: "user",
					Password: "password",
					Phone:    "08123456789",
					Role:     "admin",
				},
			},
			want: &entity.User{
				ID:       1,
				Name:     "User",
				Username: "user",
				Password: "password",
				Phone:    "08123456789",
				Role:     "admin",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userRepo{
				db: tt.fields.db,
			}
			got, err := c.UpdateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepo.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_DeleteUser(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			name: "success DeleteUser",
			fields: fields{
				db: func() *sql.DB {
					db, mock, err := sqlmock.New()
					if err != nil {
						t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
					}

					mock.ExpectExec("DELETE FROM users").
						WithArgs(1).
						WillReturnResult(sqlmock.NewResult(1, 1))

					return db
				}(),
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userRepo{
				db: tt.fields.db,
			}
			if err := c.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userRepo.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
