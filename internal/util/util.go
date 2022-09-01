package util

func StrListFrom(strPtrs []*string) []string {
	var strs []string

	for _, sPtr := range strPtrs {
		strs = append(strs, *sPtr)
	}

	return strs
}
