package model

type Document struct {
	Id               string
	Title            string
	UserId           string // Author
	Versions         []Version
	CurrentVersionId int
	IsDeleted        bool
}

type DocumentDetails struct {
	DocumentDetails []Document
}
