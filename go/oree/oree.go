package oree

type TrailId string
type Trail struct {
	Description string
}

type StepId string

type StepStatus int

const (
	Active StepStatus = iota
	Pinned
	Archived
	NotFound
)

type Step struct {
	Description string
}

type OreeI interface {
	Trails() TrailsI
}

type TrailsI interface {
	ListI[Trail, TrailI, TrailId]
}

type TrailI interface {
	Id() TrailId
	Data() Trail
	Update(Trail)
	StepWithId(StepId) (StepI, StepStatus)
	StepsWithStatus(StepStatus) StepsI
}

type StepsI interface {
	ListI[Step, StepI, StepId]
}

type StepI interface {
	Id() StepId
	Data() Step
	Update(Step)
	UpdateStatus(StepStatus)
	Status() StepStatus
}

type ListI[D any, H any, I comparable] interface {
	Len() int

	CreateFront(D) H
	CreateBack(D) H
	CreateBefore(D, H) H
	CreateAfter(D, H) H

	WithId(I) (H, bool)
	FirstN(int) []H
	LastN(int) []H
	NAfter(int, H) []H
	NBefore(int, H) []H

	PlaceFront(H)
	PlaceBack(H)
	PlaceBefore(H, H)
	PlaceAfter(H, H)

	Delete(H)
}
