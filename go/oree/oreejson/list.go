package oreejson

type ItemHandleConverter[I comparable, E any, D any, H any] interface {
	emptyHandle() H
	newItem(D) ListItem[I, E]
	itemToHandle(ListItem[I, E]) H
	handleToItem(H) ListItem[I, E]
}

type ListOJ[I comparable, E any, D any, H any] struct {
	*ListJD[I, E]
	Converter ItemHandleConverter[I, E, D, H]
}

func ListFromData[I comparable, E any, D any, H any](
	d *ListJD[I, E],
	converter ItemHandleConverter[I, E, D, H]) ListOJ[I, E, D, H] {
	return ListOJ[I, E, D, H]{
		ListJD:    d,
		Converter: converter,
	}
}

func (li ListOJ[I, E, D, H]) itemsToHandles(items []ListItem[I, E]) []H {
	handles := make([]H, len(items))
	for i, item := range items {
		handles[i] = li.Converter.itemToHandle(item)
	}
	return handles
}

func (li ListOJ[I, E, D, H]) Len() int {
	return li.ListJD.Len()
}

func (li ListOJ[I, E, D, H]) CreateFront(d D) H {
	item := li.Converter.newItem(d)
	li.PlaceItemFront(item)
	return li.Converter.itemToHandle(item)
}

func (li ListOJ[I, E, D, H]) CreateBack(d D) H {
	item := li.Converter.newItem(d)
	li.PlaceItemBack(item)
	return li.Converter.itemToHandle(item)
}

func (li ListOJ[I, E, D, H]) CreateBefore(d D, nbr H) H {
	item := li.Converter.newItem(d)
	nbrItem := li.Converter.handleToItem(nbr)
	li.PlaceItemBefore(item, nbrItem)
	return li.Converter.itemToHandle(item)
}

func (li ListOJ[I, E, D, H]) CreateAfter(d D, nbr H) H {
	item := li.Converter.newItem(d)
	nbrItem := li.Converter.handleToItem(nbr)
	li.PlaceItemAfter(item, nbrItem)
	return li.Converter.itemToHandle(item)
}

func (li ListOJ[I, E, D, H]) WithId(id I) (H, bool) {
	item, ok := li.ItemWithId(id)
	if !ok {
		return li.Converter.emptyHandle(), false
	}
	return li.Converter.itemToHandle(item), true
}

func (li ListOJ[I, E, D, H]) FirstN(n int) []H {
	return li.itemsToHandles(li.FirstNItems(n))
}

func (li ListOJ[I, E, D, H]) LastN(n int) []H {
	return li.itemsToHandles(li.LastNItems(n))
}

func (li ListOJ[I, E, D, H]) NAfter(n int, h H) []H {
	item := li.Converter.handleToItem(h)
	return li.itemsToHandles(li.NItemsAfter(n, item))
}

func (li ListOJ[I, E, D, H]) NBefore(n int, h H) []H {
	item := li.Converter.handleToItem(h)
	return li.itemsToHandles(li.NItemsBefore(n, item))
}

func (li ListOJ[I, E, D, H]) PlaceFront(h H) {
	item := li.Converter.handleToItem(h)
	li.PlaceItemFront(item)
}

func (li ListOJ[I, E, D, H]) PlaceBack(h H) {
	item := li.Converter.handleToItem(h)
	li.PlaceItemBack(item)
}

func (li ListOJ[I, E, D, H]) PlaceBefore(h H, nbr H) {
	item := li.Converter.handleToItem(h)
	nbrItem := li.Converter.handleToItem(nbr)
	li.PlaceItemBefore(item, nbrItem)
}

func (li ListOJ[I, E, D, H]) PlaceAfter(h H, nbr H) {
	item := li.Converter.handleToItem(h)
	nbrItem := li.Converter.handleToItem(nbr)
	li.PlaceItemAfter(item, nbrItem)
}

func (li ListOJ[I, E, D, H]) Delete(h H) {
	item := li.Converter.handleToItem(h)
	li.DeleteItem(item)
}
