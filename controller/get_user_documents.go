package controller

import (
	"DocumentManager/errors"
	"DocumentManager/model"
	"fmt"
)

type getUserDocuments struct {
	*store
}

type userDoc struct {
	Title   string
	Content string
}

func NewGetUserDocuments(st *store) *getUserDocuments {
	return &getUserDocuments{st}
}

func (st *getUserDocuments) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.GetUserDocuments)

	if !ok {
		return errors.ErrInvalidCommand
	}

	isValidUser := isValidUser(args.UserId, args.Password)
	if !isValidUser {
		return errors.ErrUserNotExists
	}

	userDocs := []userDoc{}

	for _, document := range DocumentDetails.DocumentDetails {
		if document.UserId == args.UserId && !document.IsDeleted {
			userDocs = append(userDocs, userDoc{
				Title:   document.Title,
				Content: document.Versions[document.CurrentVersionId-1].Content,
			})
		}
	}

	fmt.Println("The user docs are:", userDocs)

	return nil
}

func isValidUser(userId, password string) bool {
	isUserValid := false

	for _, user := range UserDetails.UserDetails {
		if user.UserId == userId {
			if user.Password == password {
				isUserValid = true
			}
		}
	}

	return isUserValid

}
