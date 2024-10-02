package oreejson

type ItemHandleConverter[I comparable, E any, D any, H any] interface {
	emptyHandle() H
	newItem(D) ListItem[I, E]
	updatedItem(ListItem[I, E], D) ListItem[I, E]
	itemToHandle(ListItem[I, E]) H
	handleToItem(H) ListItem[I, E]
}

type OrderedListJson[I comparable, E any, D any, H any] struct {
	*OrderedListJsonData[I, E]
	Converter ItemHandleConverter[I, E, D, H]
}

func OrderedListJsonFromData[I comparable, E any, D any, H any](
	d *OrderedListJsonData[I, E],
	converter ItemHandleConverter[I, E, D, H]) OrderedListJson[I, E, D, H] {
	return OrderedListJson[I, E, D, H]{
		OrderedListJsonData: d,
		Converter:           converter,
	}
}

func (ol OrderedListJson[I, E, D, H]) itemsToHandles(items []ListItem[I, E]) []H {
	handles := make([]H, len(items))
	for i, item := range items {
		handles[i] = ol.Converter.itemToHandle(item)
	}
	return handles
}

func (ol OrderedListJson[I, E, D, H]) Len() int {
	return ol.OrderedListJsonData.Len()
}

func (ol OrderedListJson[I, E, D, H]) CreateFront(d D) H {
	item := ol.Converter.newItem(d)
	ol.PlaceItemFront(item)
	return ol.Converter.itemToHandle(item)
}

func (ol OrderedListJson[I, E, D, H]) CreateBack(d D) H {
	item := ol.Converter.newItem(d)
	ol.PlaceItemBack(item)
	return ol.Converter.itemToHandle(item)
}

func (ol OrderedListJson[I, E, D, H]) CreateBefore(d D, nbr H) H {
	item := ol.Converter.newItem(d)
	nbrItem := ol.Converter.handleToItem(nbr)
	ol.PlaceItemBefore(item, nbrItem)
	return ol.Converter.itemToHandle(item)
}

func (ol OrderedListJson[I, E, D, H]) CreateAfter(d D, nbr H) H {
	item := ol.Converter.newItem(d)
	nbrItem := ol.Converter.handleToItem(nbr)
	ol.PlaceItemAfter(item, nbrItem)
	return ol.Converter.itemToHandle(item)
}

func (ol OrderedListJson[I, E, D, H]) WithId(id I) (H, bool) {
	item, ok := ol.ItemWithId(id)
	if !ok {
		return ol.Converter.emptyHandle(), false
	}
	return ol.Converter.itemToHandle(item), true
}

func (ol OrderedListJson[I, E, D, H]) FirstN(n int) []H {
	return ol.itemsToHandles(ol.FirstNItems(n))
}

func (ol OrderedListJson[I, E, D, H]) LastN(n int) []H {
	return ol.itemsToHandles(ol.LastNItems(n))
}

func (ol OrderedListJson[I, E, D, H]) NAfter(n int, h H) []H {
	item := ol.Converter.handleToItem(h)
	return ol.itemsToHandles(ol.NItemsAfter(n, item))
}

func (ol OrderedListJson[I, E, D, H]) NBefore(n int, h H) []H {
	item := ol.Converter.handleToItem(h)
	return ol.itemsToHandles(ol.NItemsBefore(n, item))
}

func (ol OrderedListJson[I, E, D, H]) PlaceFront(h H) {
	item := ol.Converter.handleToItem(h)
	ol.PlaceItemFront(item)
}

func (ol OrderedListJson[I, E, D, H]) PlaceBack(h H) {
	item := ol.Converter.handleToItem(h)
	ol.PlaceItemBack(item)
}

func (ol OrderedListJson[I, E, D, H]) PlaceBefore(h H, nbr H) {
	item := ol.Converter.handleToItem(h)
	nbrItem := ol.Converter.handleToItem(nbr)
	ol.PlaceItemBefore(item, nbrItem)
}

func (ol OrderedListJson[I, E, D, H]) PlaceAfter(h H, nbr H) {
	item := ol.Converter.handleToItem(h)
	nbrItem := ol.Converter.handleToItem(nbr)
	ol.PlaceItemAfter(item, nbrItem)
}

func (ol OrderedListJson[I, E, D, H]) Update(h H, d D) {
	item := ol.Converter.updatedItem(ol.Converter.handleToItem(h), d)
	ol.UpdateItem(item)
}

func (ol OrderedListJson[I, E, D, H]) Delete(h H) {
	item := ol.Converter.handleToItem(h)
	ol.DeleteItem(item)
}
