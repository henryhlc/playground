package oreejson

type RefListJD[I comparable] struct {
	*ListJD[I, bool]
}

type RefListOJ[I comparable, H any] struct {
	ListOJ[I, bool, bool, ListItem[I, bool]]
	Resolver RefResolver[I, H]
}

type RefResolver[I comparable, H any] interface {
	ResolveRef(I) (H, bool)
	ExtractRef(H) I
	EmptyRef() I
}

func NewRefListJD[I comparable]() *RefListJD[I] {
	return &RefListJD[I]{
		ListJD: NewListJD[I, bool](),
	}
}

func RefListFromData[I comparable, H any](d *RefListJD[I], resolver RefResolver[I, H], emptyId I) RefListOJ[I, H] {
	return RefListOJ[I, H]{
		ListOJ: ListFromData[I, bool, bool, ListItem[I, bool]](
			d.ListJD,
			identityConverter[I]{
				emptyId: emptyId,
				val:     nil,
			},
		),
		Resolver: resolver,
	}
}

type identityConverter[I comparable] struct {
	emptyId I
	val     *bool
}

func (ic identityConverter[I]) emptyHandle() ListItem[I, bool] {
	return ListItem[I, bool]{
		Id:   ic.emptyId,
		Elem: ic.val,
	}
}

func (ic identityConverter[I]) newItem(d bool) ListItem[I, bool] {
	return ListItem[I, bool]{
		Id:   ic.emptyId,
		Elem: ic.val,
	}
}

func (ic identityConverter[I]) updatedItem(item ListItem[I, bool], d bool) ListItem[I, bool] {
	return item
}

func (ic identityConverter[I]) itemToHandle(item ListItem[I, bool]) ListItem[I, bool] {
	return item
}

func (ic identityConverter[I]) handleToItem(h ListItem[I, bool]) ListItem[I, bool] {
	return h
}

func (rl RefListOJ[I, H]) idItemsToHandles(ids []ListItem[I, bool]) []H {
	handles := []H{}
	for _, id := range ids {
		handle, ok := rl.Resolver.ResolveRef(id.Id)
		if ok {
			handles = append(handles, handle)
		}
	}
	return handles
}

func (rl RefListOJ[I, H]) idItemFromHandle(h H) ListItem[I, bool] {
	return ListItem[I, bool]{
		Id:   rl.Resolver.ExtractRef(h),
		Elem: nil,
	}
}

func (rl RefListOJ[I, H]) Len() int {
	return rl.ListOJ.Len()
}

func (rl RefListOJ[I, H]) WithId(id I) (H, bool) {
	return rl.Resolver.ResolveRef(id)
}

func (rl RefListOJ[I, H]) FirstN(n int) []H {
	return rl.idItemsToHandles(rl.ListOJ.FirstN(n))
}

func (rl RefListOJ[I, H]) LastN(n int) []H {
	return rl.idItemsToHandles(rl.ListOJ.LastN(n))
}
func (rl RefListOJ[I, H]) NAfter(n int, h H) []H {
	return rl.idItemsToHandles(rl.ListOJ.NAfter(n, rl.idItemFromHandle(h)))
}
func (rl RefListOJ[I, H]) NBefore(n int, h H) []H {
	return rl.idItemsToHandles(rl.ListOJ.NBefore(n, rl.idItemFromHandle(h)))
}
func (rl RefListOJ[I, H]) PlaceFront(h H) {
	rl.ListOJ.PlaceFront(rl.idItemFromHandle(h))
}
func (rl RefListOJ[I, H]) PlaceBack(h H) {
	rl.ListOJ.PlaceBack(rl.idItemFromHandle(h))
}
func (rl RefListOJ[I, H]) PlaceBefore(h H, nbr H) {
	rl.ListOJ.PlaceBefore(
		rl.idItemFromHandle(h),
		rl.idItemFromHandle(nbr),
	)
}
func (rl RefListOJ[I, H]) PlaceAfter(h H, nbr H) {
	rl.ListOJ.PlaceAfter(
		rl.idItemFromHandle(h),
		rl.idItemFromHandle(nbr),
	)
}
func (rl RefListOJ[I, H]) Delete(h H) {
	rl.ListOJ.Delete(rl.idItemFromHandle(h))
}
