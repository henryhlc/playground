package oree

type OreeI interface {
	Trails() TrailsI
}

type TrailId string

type Trail struct {
	Description string
}

type TrailsI interface {
	CreateTrail(Trail) TrailId
	TrailWithId(TrailId) TrailI
}

type TrailI interface {
	Data() Trail
}
