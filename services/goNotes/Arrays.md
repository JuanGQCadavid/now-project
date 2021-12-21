[n]T -> and array of n lenght of type T

	var a [10] int
	a[0] = 58
	
	primes := [6]int{2,3,4,5,6,7}
	
By default, all vales wold have an value.

Slices 
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s[]int  = primes[1:4]
	
Damn!

	s := []struct {
		i int
		b bool
	} {
		{2,true},
		{3, false}
	}
	
	
	func main() {
		s := []int{2, 3, 5, 7, 11, 13}
		printSlice(s)

		// Slice the slice to give it zero length.
		s = s[:0]
		printSlice(s)

		// Extend its length.
		s = s[:4]
		printSlice(s)

		// Drop its first two values.
		s = s[2:]
		printSlice(s)
	}

	func printSlice(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

Append
	
	a:= []int{1,2}
	
	b:= append(a,1,2,5)