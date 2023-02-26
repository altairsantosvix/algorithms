package main

import ( 
    "fmt"
)

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

func main() {

    arr := []int{1, 2, 3, 4, 4,5}

    n := FirstDuplicateValue( arr)

    fmt.Printf(" First duplicate value in array is %d \n",  n)

}

