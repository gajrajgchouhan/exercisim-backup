package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	u := map[string]int {
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
    return u
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    newBill:=make(map[string]int)
    return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item string, unit string) bool {
    v,c:=units[unit]
    if !c{
        return false
    }
    bill[item]+=v
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool{
    v,c:=bill[item]
    if !c{
        return false
    }
	v1,c1:=units[unit]
    if !c1{
        return false
    }
	newV:=v-v1
    if newV<0{
        return false
    }else if newV==0{
    	delete(bill,item)
       
    }else{
    	bill[item]-=units[unit]
    }
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    v,c:=bill[item]
    return v,c
}