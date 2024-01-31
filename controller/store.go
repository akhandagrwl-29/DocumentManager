package controller

import "DocumentManager/model"

type store struct {
	createUser       model.CMD
	createDocument   model.CMD
	updateDocument   model.CMD
	deleteDocument   model.CMD
	getUserDocuments model.CMD
	revertToVersion  model.CMD
}

var Store *store
var DocumentDetails *model.DocumentDetails
var UserDetails *model.UserDetails

func InitUser() {
	UserDetails = new(model.UserDetails)
}

func InitDocumentDetails() {
	DocumentDetails = new(model.DocumentDetails)
}

func InitStore() {
	Store = new(store)
	Store.createUser = NewCreateUser(Store)
	Store.createDocument = NewCreateDocument(Store)
	Store.updateDocument = NewUpdateDocument(Store)
	Store.deleteDocument = NewDeleteDocument(Store)
	Store.getUserDocuments = NewGetUserDocuments(Store)
	Store.revertToVersion = NewRevertToVersion(Store)

}

// Getter methods
func (st *store) CreateUser() model.CMD {
	return st.createUser
}
func (st *store) CreateDocument() model.CMD {
	return st.createDocument
}
func (st *store) UpdateDocument() model.CMD {
	return st.updateDocument
}

func (st *store) DeleteDocument() model.CMD {
	return st.deleteDocument
}

func (st *store) GetUserDocuments() model.CMD {
	return st.getUserDocuments
}

func (st *store) RevertToVersion() model.CMD {
	return st.revertToVersion
}
