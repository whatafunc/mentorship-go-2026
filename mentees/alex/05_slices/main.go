package main

// "lo"
// "slices"

func CompareProductLists(oldList, newList []string) (added, removed []string) {
	common := intersect(oldList, newList) // в А и в Б

	for _, item := range oldList {
		if !contains(common, item) {
			removed = append(removed, item)
		}
	}

	for _, item := range newList {
		if !contains(common, item) {
			added = append(added, item)
		}
	}

	return added, removed
}

// intersect returns the common elements between two slices.
func intersect(a, b []string) []string {
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	var common []string
	for _, item := range b {
		if _, found := m[item]; found {
			common = append(common, item)
		}
	}
	return common
}

// contains checks if a slice contains a string.
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	oldList := []string{"apple", "banana", "cherry"}
	newList := []string{"banana", "cherry", "date", "fig"}
	//oldList := []string{"apple", "banana", "orange"}
	//newList := []string{"banana", "kiwi", "apple"}

	added, removed := CompareProductLists(oldList, newList)

	println("Added products:")
	for _, product := range added {
		println(product)
	}

	println("\nRemoved products:")
	for _, product := range removed {
		println(product)
	}

}
