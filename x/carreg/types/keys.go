package types

const (
	// ModuleName defines the module name
	ModuleName = "carreg"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_carreg"
)

var (
	ParamsKey = []byte("p_carreg")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
    IvcKey = "Ivc/value/"
)

const (
    IvcCountKey = "Ivc/count/"
)