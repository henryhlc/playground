package oree

import "time"

type BlockId string
type Block struct {
	Description    string
	StartTime      time.Time
	Duration       time.Duration
	TargetDuration time.Duration
	Context        interface{} // Area, Trail or nil
}

type BlocksI interface {
	SortedListI[Block, BlockI, BlockId]
}

type BlockI interface {
	Id() BlockId
	Data() (Block, bool)
	Update(Block)
}

type OpenSession struct {
	StartTime time.Time
	Trail     TrailI
	Step      StepI
}

type OpenSessionManagerI interface {
	Data() (OpenSession, bool)
	Open(OpenSession)
	Close(time.Time)
}

type SessionId string
type Session struct {
	StartTime time.Time
	Duration  time.Duration
	Trail     TrailI
	Step      StepI
}

type SessionI interface {
	Id() SessionId
	Data() (Session, bool)
	Update(Session)
}

type SessionsI interface {
	SortedListI[Session, SessionI, SessionId]
}

type AreaId string
type Area struct {
	Description string
}

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
	Blocks() BlocksI
	Trails() TrailsI
	Areas() AreasI
	Sessions() SessionsI
	OpenSessionManager() OpenSessionManagerI
}

type AreaI interface {
	Id() AreaId
	Data() Area
	Update(Area)
	Trails() RefListI[TrailI, TrailId]
}

type AreasI interface {
	ListI[Area, AreaI, AreaId]
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

type SortedListI[D any, H any, I comparable] interface {
	Len() int
	Create(D) H

	WithId(I) (H, bool)
	FirstN(int) []H
	LastN(int) []H
	NAfter(int, H) []H
	NBefore(int, H) []H

	Delete(H)
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

type RefListI[H any, I comparable] interface {
	Len() int

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
