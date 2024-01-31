package controller

import (
	"DocumentManager/errors"
	"DocumentManager/model"
	"fmt"
)

type revertToVersion struct {
	*store
}

func NewRevertToVersion(st *store) *revertToVersion {
	return &revertToVersion{st}
}

func (st *revertToVersion) Execute(command *model.Command) error {
	args, ok := command.Arguments.(model.RevertToVersion)

	if !ok {
		return errors.ErrInvalidCommand
	}

	isValidUser := isValidUser(args.UserId, args.Password)
	if !isValidUser {
		return errors.ErrUserNotExists
	}

	isDocumentExist := isDocumentExist(args.DocumentId)
	if !isDocumentExist {
		return errors.ErrDocumentNotExists
	}

	for idx, document := range DocumentDetails.DocumentDetails {
		if document.Id == args.DocumentId {
			if len(document.Versions) < args.VersionId {
				return errors.ErrInvalidVersion
			} else {
				DocumentDetails.DocumentDetails[idx].CurrentVersionId = args.VersionId
			}
		}
	}

	fmt.Printf("Successfully updated the version of doc: %s as %d\n", args.DocumentId, args.VersionId)
	return nil
}
