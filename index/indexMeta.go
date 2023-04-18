package index

type Meta struct {
	Segments []*SegmentMeta `json:"segments"`
}

type SegmentMeta struct {
	Id     string `json:"id"`
	MaxDoc uint32 `json:"max_doc"`
}
