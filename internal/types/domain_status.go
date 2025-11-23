package types

type DomainStatus string

const (
	DomainStatusCreated    DomainStatus = "created"
	DomainStatusProcessing DomainStatus = "processing"
	DomainStatusCompleted  DomainStatus = "completed"
	DomainStatusFailed     DomainStatus = "failed"
)
