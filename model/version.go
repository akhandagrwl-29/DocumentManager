package model

type Version struct {
	Id        int
	Content   string
	Timestamp string
	State     string
}

const (
	StateDraft     = "DRAFT"
	StatePublished = "PUBLISHED"
)
