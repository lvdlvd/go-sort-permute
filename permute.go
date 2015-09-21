// Package permute provides a function, NextPermutation, that generates permutations of any collection that satisfies sort.Interface.
package permute

import "sort"

// NextPermutation generates the next permutation of the
// sortable collection x in lexical order.  It returns false
// if the permutations are exhausted.
//
// Knuth, Donald (2011), "Section 7.2.1.2: Generating All Permutations",
// The Art of Computer Programming, volume 4A.
//
// <iframe src="http://play.golang.org/p/ljft9xhOEn" frameborder="0" style="width: 100%; height: 100%"><a href="http://play.golang.org/p/ljft9xhOEn">see this code in play.golang.org</a></iframe>
func NextPermutation(x sort.Interface) bool {
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
