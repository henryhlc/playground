package oree

type OreeI interface {
	Trails() TrailsI
}

type TrailId string

type Trail struct {
	Description string
}

type OrderedListI[D any, H any, I comparable] interface {
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

	Update(H, D)

	Delete(H)
}

// For methods that accepts TrailI, one should only pass
// TrailI's that are obtrained from the same instance
// of TrailsI.
type TrailsI interface {
	OrderedListI[Trail, TrailI, TrailId]
}

type TrailI interface {
	Id() TrailId
	Data() Trail
	Update(Trail)
}
