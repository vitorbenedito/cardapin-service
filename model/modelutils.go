package model

type RootJSON interface {
	GetId() uint
}

type ConflictChecker interface {
	HasConflict() (error, []string)
}

func getNotFoundValues(firstSlice []uint, otherSlice []uint) []uint {
	notUsedMap := make(map[uint]bool)
	notUsedSlice := make([]uint, 0)
	for _, value := range firstSlice {
		found := false
		for _, otherValue := range otherSlice {
			if value == otherValue {
				found = true
			}
		}
		if !found {
			notUsedMap[value] = true
		}
	}
	for k := range notUsedMap {
		notUsedSlice = append(notUsedSlice, k)
	}
	return notUsedSlice
}
