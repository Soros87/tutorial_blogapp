package types

const (
	// ModuleName defines the module name
	ModuleName = "blogapp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blogapp"
)

var (
	ParamsKey = []byte("p_blogapp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
