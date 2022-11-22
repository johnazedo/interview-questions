package medium

import datastructures.BinaryMatrix

object LeftMostColumnWithAtLeastAOne {

    fun leftMostColumnWithOne(binaryMatrix: BinaryMatrix):Int {
        val dimension = binaryMatrix.dimensions()
        var pivot: Int
        var number: Int
        var prevNumber: Int
        var result: Int = -1
        var smallest = 9999

        // O(N*logN) Solution
        for(i in 0 until dimension[0]) {
            var left = 0
            var right = dimension[1] - 1

            // Put binary search here
            while(left <= right) {
                pivot = left + (right - left)/2
                number = binaryMatrix.get(i, pivot)

                if(number == 1) {
                    prevNumber = if(pivot > 0) binaryMatrix.get(i, pivot-1) else 0
                    if(prevNumber == 0) {
                        if(pivot < smallest){
                            result = pivot
                            smallest = pivot
                        }
                        break
                    }
                    right = pivot - 1
                }
                if(number == 0) left = pivot + 1
            }
        }

        return result
    }


    fun oofNSquareSolution(binaryMatrix: BinaryMatrix): Int {
        val dimension = binaryMatrix.dimensions()

        // O(n) Solution
        for(c in 0 until dimension[1]){
            for(r in 0 until dimension[0]) {
                if(binaryMatrix.get(r, c) == 1 ) {
                    return c
                }
            }
        }

        return -1
    }
}