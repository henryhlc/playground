package oreejson

import (
	"testing"
)

type TestConverter struct {
	NextId int
}

func (tc *TestConverter) emptyHandle() ListItem[int, int] {
	return ListItem[int, int]{}
}

func (tc *TestConverter) newItem(d int) ListItem[int, int] {
	id := tc.NextId
	tc.NextId++
	k := d
	return ListItem[int, int]{
		Id:   id,
		Elem: &k,
	}
}

func (tc *TestConverter) itemToHandle(item ListItem[int, int]) ListItem[int, int] {
	return item
}

func (tc *TestConverter) handleToItem(item ListItem[int, int]) ListItem[int, int] {
	return item
}

type TestComparator struct{}

func (tc TestComparator) Compare(a *int, b *int) int {
	return *a - *b
}

func TestCreate(t *testing.T) {
	sl := SortedListFromData[int, int, int, ListItem[int, int]](
		NewListJD[int, int](),
		&TestConverter{NextId: 0},
		TestComparator{},
	)
	sl.Create(1)
	sl.Create(3)
	sl.Create(5)
	sl.Create(7)
	sl.Create(2)
	sl.Create(4)
	sl.Create(6)
	items := sl.FirstN(5)
	if len(items) != 5 {
		t.Errorf("Number of items expected 5, actual %v", len(items))
	}
	for i := range min(len(items), 5) {
		expected := i + 1
		if actual := *items[i].Elem; actual != expected {
			t.Errorf("Item %v expected %v, actual %v", i, expected, actual)
		}
	}
}

func TestFixOrder(t *testing.T) {
	sl := SortedListFromData[int, int, int, ListItem[int, int]](
		NewListJD[int, int](),
		&TestConverter{NextId: 0},
		TestComparator{},
	)
	sl.Create(1)
	sl.Create(3)
	item := sl.Create(4)
	*item.Elem = 2
	sl.FixOrder(item)

	items := sl.FirstN(5)
	if len(items) != 3 {
		t.Errorf("Number of items expected 3, actual %v", len(items))
	}
	for i := range min(len(items), 5) {
		expected := i + 1
		if actual := *items[i].Elem; actual != expected {
			t.Errorf("Item %v expected %v, actual %v", i, expected, actual)
		}
	}

}
