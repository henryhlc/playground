package oreejson_test

import (
	"fmt"
	"testing"

	"github.com/henryhlc/playground/go/oree/oreejson"
)

func createTestItems(n int) []oreejson.ListItem[string, string] {
	items := make([]oreejson.ListItem[string, string], n)
	for i := range n {
		elem := fmt.Sprintf("elem %v", i)
		items[i] = oreejson.ListItem[string, string]{
			Id:   fmt.Sprintf("id %v", i),
			Elem: &elem,
		}
	}
	return items
}

func matchTestItem(actual, expected oreejson.ListItem[string, string], t *testing.T) bool {
	isMatch := true
	if actual.Id != expected.Id {
		isMatch = false
		t.Errorf("Item Id does not match, actual: %v, expected: %v", actual.Id, expected.Id)
	}
	if *actual.Elem != *expected.Elem {
		isMatch = false
		t.Errorf("Item Elem pointee does not match, actual: %v, expected: %v", *actual.Elem, *expected.Elem)
	}
	return isMatch
}

func matchTestItems(actual, expected []oreejson.ListItem[string, string], t *testing.T) bool {
	isMatch := true
	if len(actual) != len(expected) {
		isMatch = false
		t.Errorf("Length different, actual: %v, expected: %v", len(actual), len(expected))
	}
	for i := range max(len(actual), len(expected)) {
		switch {
		case i < len(actual) && i < len(expected):
			itemMatch := matchTestItem(actual[i], expected[i], t)
			if !itemMatch {
				isMatch = false
				t.Errorf("Mismatch found at index %v", i)
			}
		case i < len(actual):
			isMatch = false
			t.Errorf("Actual has additional item with id \"%v\" and elem value \"%v\"", actual[i].Id, *actual[i].Elem)
		case i < len(expected):
			isMatch = false
			t.Errorf("Expected has additional item with id \"%v\" and elem value \"%v\"", expected[i].Id, *expected[i].Elem)
		}
	}
	return isMatch
}

func TestLen(t *testing.T) {
	items := createTestItems(2)
	ol := oreejson.NewOrderedListJD[string, string]()
	if actual := ol.Len(); actual != 0 {
		t.Errorf("Len() expected 0 actual %v", actual)
	}
	ol.PlaceItemFront(items[0])
	if actual := ol.Len(); actual != 1 {
		t.Errorf("Len() expected 1 actual %v", actual)
	}
	ol.PlaceItemFront(items[1])
	if actual := ol.Len(); actual != 2 {
		t.Errorf("Len() expected 2 actual %v", actual)
	}
}

func TestPlaceItemBack(t *testing.T) {
	ol := oreejson.NewOrderedListJD[string, string]()
	items := createTestItems(2)
	ol.PlaceItemBack(items[0])
	if !matchTestItems(ol.LastNItems(3), items[:1], t) {
		t.Errorf("^ after one place back.")
	}
	ol.PlaceItemBack(items[1])
	if !matchTestItems(ol.LastNItems(3), items, t) {
		t.Errorf("^ after two place back.")
	}
	ol.PlaceItemBack(items[0])
	expected := []oreejson.ListItem[string, string]{
		items[1], items[0],
	}
	if !matchTestItems(ol.LastNItems(3), expected, t) {
		t.Errorf("^ after placing existing item back.")
	}
}

func TestPlaceItemFront(t *testing.T) {
	ol := oreejson.NewOrderedListJD[string, string]()
	items := createTestItems(2)
	ol.PlaceItemFront(items[1])
	if !matchTestItems(ol.FirstNItems(3), items[1:], t) {
		t.Errorf("^ after one place front.")
	}
	ol.PlaceItemFront(items[0])
	if !matchTestItems(ol.FirstNItems(3), items, t) {
		t.Errorf("^ after two place front.")
	}
	ol.PlaceItemFront(items[1])
	expected := []oreejson.ListItem[string, string]{
		items[1], items[0],
	}
	if !matchTestItems(ol.FirstNItems(3), expected, t) {
		t.Errorf("^ after placing existing item front.")
	}
}

func TestPlaceItemBefore(t *testing.T) {
	ol := oreejson.NewOrderedListJD[string, string]()
	items := createTestItems(3)
	ol.PlaceItemFront(items[2])
	if !matchTestItems(ol.LastNItems(3), items[2:], t) {
		t.Errorf("^ after one place front.")
	}
	ol.PlaceItemBefore(items[1], items[2])
	if !matchTestItems(ol.LastNItems(3), items[1:], t) {
		t.Errorf("^ after place item before one existing item.")
	}
	ol.PlaceItemBefore(items[0], items[2])
	expected := []oreejson.ListItem[string, string]{
		items[1], items[0], items[2],
	}
	if !matchTestItems(ol.LastNItems(3), expected, t) {
		t.Errorf("^ after place item between existing items.")
	}
}

func TestPlaceItemAfter(t *testing.T) {
	ol := oreejson.NewOrderedListJD[string, string]()
	items := createTestItems(3)
	ol.PlaceItemBack(items[0])
	if !matchTestItems(ol.FirstNItems(3), items[:1], t) {
		t.Errorf("^ after one place back.")
	}
	ol.PlaceItemAfter(items[1], items[0])
	if !matchTestItems(ol.FirstNItems(3), items[:2], t) {
		t.Errorf("^ after place item after one existing item.")
	}
	ol.PlaceItemAfter(items[2], items[0])
	expected := []oreejson.ListItem[string, string]{
		items[0], items[2], items[1],
	}
	if !matchTestItems(ol.FirstNItems(3), expected, t) {
		t.Errorf("^ after place item between existing items.")
	}
}

func TestItemWithId(t *testing.T) {
	ol := oreejson.NewOrderedListJD[string, string]()
	items := createTestItems(2)
	ol.PlaceItemBack(items[0])
	item, ok := ol.ItemWithId(items[0].Id)
	if !ok {
		t.Errorf("Item in the list, returns not ok.")
	} else if !matchTestItem(items[0], item, t) {
		t.Errorf("^ Returned item does not match with inserted.")
	}
	_, ok = ol.ItemWithId(items[1].Id)
	if ok {
		t.Errorf("Item not in the list, returns ok.")
	}
}

func TestNItems(t *testing.T) {
	ol := oreejson.NewOrderedListJD[string, string]()
	items := createTestItems(5)
	ol.PlaceItemBack(items[0])
	ol.PlaceItemBack(items[1])
	ol.PlaceItemBack(items[2])
	ol.PlaceItemBack(items[3])
	ol.PlaceItemBack(items[4])

	actual := ol.FirstNItems(2)
	if !matchTestItems(actual, items[:2], t) {
		t.Errorf("^ after FirstNItems")
	}
	actual = ol.NItemsAfter(4, actual[1])
	if !matchTestItems(actual, items[2:], t) {
		t.Errorf("^ after NItemsAfter")
	}
	actual = ol.LastNItems(3)
	if !matchTestItems(actual, items[2:], t) {
		t.Errorf("^ after LastNItems")
	}
	actual = ol.NItemsBefore(3, actual[0])
	if !matchTestItems(actual, items[:2], t) {
		t.Errorf("^ after NItemsBefore")
	}
}

func TestDeleteItem(t *testing.T) {
	ol := oreejson.NewOrderedListJD[string, string]()
	items := createTestItems(3)

	ol.DeleteItem(items[0]) // should be no-op
	ol.PlaceItemFront(items[0])
	ol.DeleteItem(items[1]) // should be no-op
	ol.DeleteItem(items[0])
	if ol.Len() > 0 {
		t.Errorf("Non-zero length after deleting only item.")
	}

	ol.PlaceItemBack(items[0])
	ol.PlaceItemBack(items[1])
	ol.DeleteItem(items[0])
	if ol.Len() != 1 || !matchTestItems(ol.FirstNItems(3), items[1:2], t) {
		t.Errorf("^ after deleting 0 from 0-1.")
	}

	ol.PlaceItemFront(items[0])
	ol.DeleteItem(items[1])
	if ol.Len() != 1 || !matchTestItems(ol.FirstNItems(3), items[0:1], t) {
		t.Errorf("^ after deleting 1 from 0-1.")
	}

	ol.PlaceItemBack(items[2])
	ol.PlaceItemBack(items[1])
	ol.DeleteItem(items[2])
	if ol.Len() != 2 || !matchTestItems(ol.FirstNItems(3), items[0:2], t) {
		t.Errorf("^ after deleting 2 from 0-2-1.")
	}
}
