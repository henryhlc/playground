package oreejson

import (
	"time"

	"github.com/henryhlc/playground/go/oree"
)

type contextKind int

const (
	UnknownK contextKind = iota
	AreaK
	TrailK
	NoneK
)

func kindOf(in interface{}) contextKind {
	switch in.(type) {
	case oree.TrailI:
		return TrailK
	case oree.AreaI:
		return AreaK
	case nil:
		return NoneK
	default:
		return UnknownK
	}
}

type BlockJD struct {
	Description    string        `json:"description"`
	StartTime      time.Time     `json:"start"`
	Duration       time.Duration `json:"duration"`
	TargetDuration time.Duration `json:"targetDuration"`
	ContextKind    contextKind   `json:"contextKind"`
	ContextSpec    string        `json:"contextSpec"`
}

type BlockOJ struct {
	*BlockJD
	oreeJson OreeJson
	id       oree.BlockId
}

func NewBlockJD(block oree.Block) *BlockJD {
	contextKind := kindOf(block.Context)
	var contextSpec string
	switch h := block.Context.(type) {
	case AreaOJ:
		contextSpec = string(h.Id())
	case TrailOJ:
		contextSpec = string(h.Id())
	case nil:
		contextSpec = ""
	}
	return &BlockJD{
		Description:    block.Description,
		StartTime:      block.StartTime,
		Duration:       block.Duration,
		TargetDuration: block.TargetDuration,
		ContextKind:    contextKind,
		ContextSpec:    contextSpec,
	}
}

func BlockFromData(d *BlockJD, oj OreeJson, id oree.BlockId) oree.BlockI {
	return BlockOJ{
		BlockJD:  d,
		oreeJson: oj,
		id:       id,
	}
}

func (b BlockOJ) Id() oree.BlockId {
	return b.id
}

func (b BlockOJ) Data() (oree.Block, bool) {
	block := oree.Block{
		Description:    b.Description,
		StartTime:      b.StartTime,
		Duration:       b.Duration,
		TargetDuration: b.TargetDuration,
	}
	var ok bool
	switch b.ContextKind {
	case NoneK:
		return block, true
	case AreaK:
		block.Context, ok = b.oreeJson.Areas().WithId(oree.AreaId(b.ContextSpec))
		return block, ok
	case TrailK:
		block.Context, ok = b.oreeJson.Trails().WithId(oree.TrailId(b.ContextSpec))
		return block, ok
	}
	return block, false
}

func (b BlockOJ) Update(block oree.Block) {
	b.Description = block.Description
	b.StartTime = block.StartTime
	b.Duration = block.Duration
	b.TargetDuration = block.TargetDuration
	switch at := block.Context.(type) {
	case AreaOJ:
		b.ContextKind = AreaK
		b.ContextSpec = string(at.Id())
	case TrailOJ:
		b.ContextKind = TrailK
		b.ContextSpec = string(at.Id())
	}
	blocks := b.oreeJson.Blocks().(BlocksOJ)
	blocks.FixOrder(ListItem[oree.BlockId, BlockJD]{
		Id:   b.id,
		Elem: b.BlockJD,
	})
}
