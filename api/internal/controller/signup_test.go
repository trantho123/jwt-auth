package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/rds"
	"Jwtwithecdsa/api/internal/repository"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

type mockSignUp struct {
	signInput model.SignUpInput
}

func TestSighUp(t *testing.T) {
	testCase := []struct {
		name       string
		input      mockSignUp
		buildStubs func(repo *repository.MockRegistry, input mockSignUp)
		wantErr    string
	}{
		{
			name: "OK",
			input: mockSignUp{
				signInput: model.SignUpInput{
					Email:    utils.RandomEmail(),
					UserName: utils.RandomString(5),
					Password: utils.RandomString(10),
				},
			},
			buildStubs: func(repo *repository.MockRegistry, input mockSignUp) {
				// Is username exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"username": input.signInput.UserName}).Return(model.User{}, nil)
				// Is email exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"email": input.signInput.Email}).Return(model.User{}, nil)
				// Insert user
				repo.EXPECT().InsertUser(gomock.Any(), gomock.AssignableToTypeOf(&model.User{})).
					DoAndReturn(func(_ context.Context, u *model.User) error {
						require.Equal(t, input.signInput.Email, u.Email)
						require.Equal(t, input.signInput.UserName, u.Username)
						return nil
					})
				// Get user for send email
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"email": input.signInput.Email}).Return(model.User{
					Username: input.signInput.UserName,
					Email:    input.signInput.Email,
				}, nil)
			},
			wantErr: "",
		},
		{
			name: "Username already exists",
			input: mockSignUp{
				signInput: model.SignUpInput{
					Email:    utils.RandomEmail(),
					UserName: "existing_user",
					Password: utils.RandomString(10),
				},
			},
			buildStubs: func(repo *repository.MockRegistry, input mockSignUp) {
				// Is username exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"username": input.signInput.UserName}).Return(model.User{
					Username: "existing_user",
				}, nil)

			},
			wantErr: fmt.Sprintf("username '%s' already exists", "existing_user"),
		},
		{
			name: "Not Verified Email",
			input: mockSignUp{
				signInput: model.SignUpInput{
					Email:    utils.RandomEmail(),
					UserName: utils.RandomString(5),
					Password: utils.RandomString(10),
				},
			},
			buildStubs: func(repo *repository.MockRegistry, input mockSignUp) {
				// Is username exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"username": input.signInput.UserName}).Return(model.User{}, nil)
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"email": input.signInput.Email}).Return(model.User{
					Email:    input.signInput.Email,
					Verified: false,
				}, nil)
			},
			wantErr: "please verify your email",
		},
		{
			name: "Email already exists",
			input: mockSignUp{
				signInput: model.SignUpInput{
					Email:    "emailexisted@gmail.com",
					UserName: utils.RandomString(5),
					Password: utils.RandomString(10),
				},
			},
			buildStubs: func(repo *repository.MockRegistry, input mockSignUp) {
				// Is username exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"username": input.signInput.UserName}).Return(model.User{}, nil)
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"email": "emailexisted@gmail.com"}).Return(model.User{
					Email:    "emailexisted@gmail.com",
					Verified: true,
				}, nil)
			},
			wantErr: "user with that email already exists",
		},
		{
			name: "failed to insert user",
			input: mockSignUp{
				signInput: model.SignUpInput{
					Email:    utils.RandomEmail(),
					UserName: utils.RandomString(5),
					Password: utils.RandomString(10),
				},
			},
			buildStubs: func(repo *repository.MockRegistry, input mockSignUp) {
				// Is username exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"username": input.signInput.UserName}).Return(model.User{}, nil)
				// Is email exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"email": input.signInput.Email}).Return(model.User{}, nil)
				// Insert user
				repo.EXPECT().InsertUser(gomock.Any(), gomock.AssignableToTypeOf(&model.User{})).Return(
					fmt.Errorf("something bad happened"))
				// Get user for send email

			},
			wantErr: "failed to insert user",
		},
		{
			name: "failed to get user by email after insert user",
			input: mockSignUp{
				signInput: model.SignUpInput{
					Email:    utils.RandomEmail(),
					UserName: utils.RandomString(5),
					Password: utils.RandomString(10),
				},
			},
			buildStubs: func(repo *repository.MockRegistry, input mockSignUp) {
				// Is username exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"username": input.signInput.UserName}).Return(model.User{}, nil)
				// Is email exist in db
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"email": input.signInput.Email}).Return(model.User{}, nil)
				// Insert user
				repo.EXPECT().InsertUser(gomock.Any(), gomock.AssignableToTypeOf(&model.User{})).
					DoAndReturn(func(_ context.Context, u *model.User) error {
						require.Equal(t, input.signInput.Email, u.Email)
						require.Equal(t, input.signInput.UserName, u.Username)
						return nil
					})
				// Get user for send email
				repo.EXPECT().GetUser(gomock.Any(), bson.M{"email": input.signInput.Email}).Return(model.User{}, fmt.Errorf("something bad happened"))
			},
			wantErr: "something bad happened",
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := repository.NewMockRegistry(ctrl)
			tc.buildStubs(repo, tc.input)
			rds := rds.NewMockRedisService(ctrl)
			config, errConfig := utils.LoadConfig("../..")
			require.NoError(t, errConfig)
			ctr := New(repo, rds, config)

			err := ctr.SighUp(&tc.input.signInput)
			if tc.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.wantErr)
			}
		})
	}
}
