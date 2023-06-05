package index

import "sync"

type SegmentPolicy interface {
	NewSegment(s *SegmentManager) bool
}

type DefaultSegmentPolicy struct {
}

func NewDefaultSegmentPolicy() *DefaultSegmentPolicy {
	return &DefaultSegmentPolicy{}
}

func (p *DefaultSegmentPolicy) NewSegment(s *SegmentManager) bool {
	return true
}

type Segment struct {
	idx  *Index
	meta *SegmentMeta
}

func (s *Segment) AfterUpdateDoc() {
	s.meta.MaxDoc = s.meta.MaxDoc + 1
}

type SegmentManager struct {
	mergeSegments []*Segment
	segments      []*Segment
	segmentPolicy SegmentPolicy
	mux           sync.Mutex
}

func NewSegmentManager(segmentPolicy SegmentPolicy) *SegmentManager {
	return &SegmentManager{
		segmentPolicy: segmentPolicy,
	}
}

func (s *SegmentManager) NewSegment(idx *Index, meta *SegmentMeta) *Segment {
	return &Segment{
		meta: meta,
		idx:  idx,
	}
}

func (s *SegmentManager) PubSegment() {

}
