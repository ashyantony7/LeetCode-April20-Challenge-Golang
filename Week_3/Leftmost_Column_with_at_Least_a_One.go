/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */
package main

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {

	dim := binaryMatrix.Dimensions()
	row, col := 0, dim[1]-1
	var set bool

	for (row < dim[0]) && (col >= 0) {
		if binaryMatrix.Get(row, col) == 1 {
			set = true
			col--
		} else {
			row++
		}
	}

	if set {
		return col + 1
	} else {
		return -1
	}

}
