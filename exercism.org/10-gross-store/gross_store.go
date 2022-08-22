package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	grossUnit := make(map[string]int)
    
    grossUnit["quarter_of_a_dozen"] = 3
    grossUnit["half_of_a_dozen"] = 6
    grossUnit["dozen"] = 12
    grossUnit["small_gross"] = 120
    grossUnit["gross"] = 144
    grossUnit["great_gross"] = 1728

    return grossUnit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    newBill := make(map[string]int)
	return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := units[unit]; !exists {
        return false
    }

    if _, exists := bill[item]; !exists {
        bill[item] = units[unit]
    } else {
    	bill[item] += units[unit]
    }

    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

    if _, exists := units[unit]; !exists {
        return false
    }

    if _, exists := bill[item]; !exists {
        return false
    }

    if (bill[item] - units[unit]) < 0 {
        return false
    } else if (bill[item] - units[unit]) == 0 {
    	delete(bill, item)
        return true
    } else {
    	bill[item] -= units[unit]
        return true
    }    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, exists := bill[item]; !exists {
        return 0, false
    }

    return bill[item], true
}
