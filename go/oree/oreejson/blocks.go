package oreejson

import (
	"time"

	"github.com/henryhlc/playground/go/oree"
)

type BlocksJD struct {
	*ListJD[oree.BlockId, BlockJD]
}

type BlocksOJ struct {
	SortedListOJ[oree.BlockId, BlockJD, oree.Block, oree.BlockI]
	oreeJson OreeJson
}

func NewBlocksJD() *BlocksJD {
	return &BlocksJD{
		ListJD: NewListJD[oree.BlockId, BlockJD](),
	}
}

func BlocksFromData(d *BlocksJD, oj OreeJson) BlocksOJ {
	return BlocksOJ{
		SortedListOJ: SortedListFromData(
			d.ListJD,
			newItemBlockConverter(oj),
			BlockComparator{},
		),
		oreeJson: oj,
	}
}

func (bs BlocksOJ) LastBlockCovering(t time.Time) (oree.BlockI, bool) {
	lastList := bs.LastN(1)
	for len(lastList) > 0 {
		last := lastList[0]
		data, ok := last.Data()
		if !ok {
			continue
		}
		if data.StartTime.Before(t) && data.StartTime.Add(data.Duration).After(t) {
			return last, true
		}
		lastList = bs.NBefore(1, last)
	}
	return BlockOJ{}, false
}

type ItemBlockConverter struct {
	oreeJson OreeJson
}

func newItemBlockConverter(oj OreeJson) ItemBlockConverter {
	return ItemBlockConverter{oreeJson: oj}
}

func (c ItemBlockConverter) emptyHandle() oree.BlockI {
	return BlockOJ{}
}

func (c ItemBlockConverter) newItem(d oree.Block) ListItem[oree.BlockId, BlockJD] {
	return ListItem[oree.BlockId, BlockJD]{
		Id:   oree.BlockId(c.oreeJson.getAndIncId()),
		Elem: NewBlockJD(d),
	}
}

func (c ItemBlockConverter) itemToHandle(item ListItem[oree.BlockId, BlockJD]) oree.BlockI {
	return BlockFromData(item.Elem, c.oreeJson, item.Id)
}

func (c ItemBlockConverter) handleToItem(s oree.BlockI) ListItem[oree.BlockId, BlockJD] {
	soj := s.(BlockOJ)
	return ListItem[oree.BlockId, BlockJD]{
		Id:   s.Id(),
		Elem: soj.BlockJD,
	}
}

type BlockComparator struct{}

func (BlockComparator) Compare(a, b *BlockJD) int {
	return a.StartTime.Compare(b.StartTime)
}
