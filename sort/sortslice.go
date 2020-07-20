package sort

func SortSlice(slice []string) []string {
	var sortedslice []string

	if len(slice) != 0 {
		for h := 0; h < len(slice); h++ {

			sortedslice = append(sortedslice, slice[h])
		}
	}

	return sortedslice
}
