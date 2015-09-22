// Package permute provides a function that generates permutations of any collection that satisfies sort.Interface.
package permute

import "sort"

// Next generates the next permutation of the
// sortable collection x in lexical order.  It returns false
// once the permutations are exhausted, which is when x is
// in complete descending order.
//
// Knuth, Donald (2011), "Section 7.2.1.2: Generating All Permutations",
// The Art of Computer Programming, volume 4A.
//
// The test is on http://play.golang.org/p/ljft9xhOEn
func Next(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
