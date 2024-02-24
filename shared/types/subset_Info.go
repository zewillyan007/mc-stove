package types

type CacheSubsetInfo struct {
	Preview int
	Next    int
	Page    int
	Pages   int
	Total   int
	Lines   int
}

type Pair struct {
	Key   string
	Value interface{}
}
