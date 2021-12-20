package bloom

type Bloom interface {
	Location(uint64) (uint64, uint64)
	Add([]byte)
	Exists([]byte) bool
	AddString(string)
	ExistsString(string) bool
}
