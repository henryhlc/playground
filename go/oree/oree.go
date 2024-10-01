package oree

type Trail struct {
	Description string
}

type TrailId string

type OreeI interface {
	CreateTrail(Trail) TrailId
}

type Oree struct {
	OreeI
}
