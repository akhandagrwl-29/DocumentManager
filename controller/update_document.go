package controller

import (
	"DocumentManager/errors"
	"DocumentManager/model"
	"fmt"
	"time"
)

type updateDocument struct {
	*store
}

func NewUpdateDocument(st *store) *updateDocument {
	return &updateDocument{st}
}

func (st *updateDocument) Execute(command *model.Command) error {
	args, ok := command.Arguments.(model.UpdateDocument)

	if !ok {
		return errors.ErrInvalidCommand
	}

	isDocumentExist := isDocumentExist(args.Id)
	if !isDocumentExist {
		return errors.ErrDocumentNotExists
	}

	isValidUser := isValidUser(args.UserId, args.Password)
	if !isValidUser {
		return errors.ErrUserNotExists
	}

	for documentIdx, document := range DocumentDetails.DocumentDetails {
		if document.Id == args.Id {
			DocumentDetails.DocumentDetails[documentIdx].Title = args.Title
			DocumentDetails.DocumentDetails[documentIdx].Versions = append(DocumentDetails.DocumentDetails[documentIdx].Versions, model.Version{
				Id:        len(document.Versions) + 1,
				Content:   args.Content,
				Timestamp: time.Now().String(),
			})
			DocumentDetails.DocumentDetails[documentIdx].CurrentVersionId = len(document.Versions) + 1
		}
	}

	fmt.Println("Successfully updated document with id:", args.Id)
	fmt.Println("Documents so far are", DocumentDetails)
	return nil
}
