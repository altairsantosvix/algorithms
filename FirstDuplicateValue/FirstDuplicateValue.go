package FirstDuplicateValue


// O(n) time
// O(1) space
func FirstDuplicateValue(arr []int) int {
    for _,  value := range arr {

        absValue := absV(value)
        if arr[absValue-1] < 0 {
                return absValue
        } 
        arr[absValue-1] *= -1
    }
        return -1
}

func absV( n int ) int {
    if n < 0 {
        return -n
    }
     return n
}
