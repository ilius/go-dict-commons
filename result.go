package commons

type SearchResultIface interface {
	Terms() []string
	Score() uint8
	DictName() string
	DefinitionsHTML() []string
	ResourceDir() string
	EntryIndex() uint64
}

const (
	ResultFlag_FixAudio = 1 << iota
	ResultFlag_FixFileSrc
	ResultFlag_FixWordLink
	ResultFlag_ColorMapping
	ResultFlag_Web
)

type SearchResultItem struct {
	Data []byte
	Type rune
}

type SearchResultLow struct {
	Items        func() []*SearchResultItem
	F_Terms      []string
	F_Score      uint8
	F_EntryIndex uint64
}

func (res *SearchResultLow) Terms() []string {
	return res.F_Terms
}

func (res *SearchResultLow) Score() uint8 {
	return res.F_Score
}

func (res *SearchResultLow) EntryIndex() uint64 {
	return res.F_EntryIndex
}
