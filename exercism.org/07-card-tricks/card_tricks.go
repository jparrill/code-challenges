package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    var cards = []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if (len(slice) <= index) || index < 0 {
        return -1
    }

    return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if (len(slice)-1 < index) || (index < 0) {
        slice = append(slice, value)
    } else if len(slice)-1 == index {
		tmpSlice := slice[:index]
        tmpSlice = append(tmpSlice, value)
    } else {
		tmpSlice := slice[:index]
        tmpSlice = append(tmpSlice, value)
        tmpSlice = append(tmpSlice, slice[index+1])
    }

    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    var newSlice = make([]int, len(slice))
	if len(values) == 0 {
        return slice
    }

    newSlice = values
    newSlice = append(newSlice, slice...)
    return newSlice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if (len(slice)-1 < index) || (index < 0) {
        return slice
    } else {
		tmpSlice := slice[:index]
        tmpSlice = append(tmpSlice, slice[index+1:]...)
        return tmpSlice
    }

}
