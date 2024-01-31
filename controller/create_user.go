package controller

import (
	"DocumentManager/errors"
	"DocumentManager/model"
	"fmt"
)

type createUser struct {
	*store
}

func NewCreateUser(st *store) *createUser {
	return &createUser{st}
}

func (st *createUser) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.User)

	if !ok {
		return errors.ErrInvalidCommand
	}

	isUserExists := isUserExists(args.UserId)
	if isUserExists {
		return errors.ErrUserAlreadyExists
	}

	UserDetails.UserDetails = append(UserDetails.UserDetails, model.User{
		UserName: args.UserName,
		UserId:   args.UserId,
		Password: args.Password,
	})

	fmt.Println("Successfully created user with user id:", args.UserId)
	fmt.Println("List of users are:", UserDetails)
	return nil
}

func isUserExists(userId string) bool {
	for _, user := range UserDetails.UserDetails {
		if user.UserId == userId {
			return true
		}
	}
	return false
}
