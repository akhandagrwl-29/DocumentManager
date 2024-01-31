package controller

import (
	"DocumentManager/errors"
	"DocumentManager/model"
	"fmt"
)

type deleteDocument struct {
	*store
}

func NewDeleteDocument(st *store) *deleteDocument {
	return &deleteDocument{st}
}

func (st *deleteDocument) Execute(command *model.Command) error {

	args, ok := command.Arguments.(model.DeleteDocument)

	if !ok {
		return errors.ErrInvalidCommand
	}

	isDocumentExist := isDocumentExist(args.Id)
	if !isDocumentExist {
		return errors.ErrDocumentNotExists
	}

	isUserAuthenticated := isUserAuthenticated(args.UserId, args.Password, args.Id)
	if !isUserAuthenticated {
		return errors.ErrDocumentAlreadyExists
	}

	for idx, document := range DocumentDetails.DocumentDetails {
		if document.Id == args.Id {
			DocumentDetails.DocumentDetails[idx].IsDeleted = true
		}
	}

	fmt.Println("Successfully deleted the document with id:", args.Id)

	return nil

}

func isUserAuthenticated(userId, password, documentId string) bool {
	isUserValid := false
	isPasswordCorrect := false

	for _, document := range DocumentDetails.DocumentDetails {
		if document.UserId == userId && document.Id == documentId {
			isUserValid = true
		}
	}

	for _, user := range UserDetails.UserDetails {
		if user.UserId == userId {
			if user.Password == password {
				isPasswordCorrect = true
			}
		}
	}

	return isPasswordCorrect && isUserValid

}
