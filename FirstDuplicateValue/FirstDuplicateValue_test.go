//Implements testCase 

package FirstDuplicateValue

import "testing"

func TestFirstDuplicateValue( t *testing.T){
	input_array := []int{1, 2, 3, 4, 5,5}

	n := FirstDuplicateValue( input_array)  
	
	expected := 5;

	if n != expected {
	  t.Errorf("\n First Duplicate Value of Array %d expected %d, but is %d \n", input_array, expected, n )
	}

}