package controller

import (
	"DocumentManager/errors"
	"DocumentManager/model"
	"fmt"
	"time"
)

type createDocument struct {
	*store
}

func NewCreateDocument(st *store) *createDocument {
	return &createDocument{st}
}

func (st *createDocument) Execute(command *model.Command) error {
	args, ok := command.Arguments.(model.CreateDocument)

	if !ok {
		return errors.ErrInvalidCommand
	}

	isDocumentExist := isDocumentExist(args.Id)
	if isDocumentExist {
		return errors.ErrUserAlreadyExists
	}

	DocumentDetails.DocumentDetails = append(DocumentDetails.DocumentDetails, model.Document{
		Id:     args.Id,
		Title:  args.Title,
		UserId: args.UserId,
		Versions: []model.Version{
			{
				Id:        1,
				Content:   args.Content,
				Timestamp: time.Now().String(),
			},
		},
		CurrentVersionId: 1,
	})

	fmt.Println("Successfully created document with id:", args.Id)
	fmt.Println("List of Documents are:", DocumentDetails)
	return nil
}

func isDocumentExist(id string) bool {
	for _, document := range DocumentDetails.DocumentDetails {
		if document.Id == id && document.IsDeleted == false {
			return true
		}
	}
	return false
}
