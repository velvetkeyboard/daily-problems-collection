/*
Hi, here's your problem today. This problem was recently asked by Amazon:

Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.

Here is the method signature:
```
class Solution:
  def minSubArrayLen(self, nums, s):
    # Fill this in

print Solution().minSubArrayLen([2, 3, 1, 2, 4, 3], 7)
# 2
```
*/
package main

import "fmt"

func minSubArrayLen(arr []int, sumLimit int) int {
  ret := len(arr)
  for i := 0; i < len(arr); i++ {
    if arr[i] == sumLimit {
      ret = 1
      continue
    }

    sum := arr[i]

    for j := i + 1; j < len(arr); j++ {
      sum += arr[j]
      if sum > sumLimit {
        break
      } else if sum == sumLimit {
        if ret > j - i {
          ret = (j - i) + 1
        }
        break
      }
    }
  }
  return ret
}

func assertEqualInt(a int, b int) {
  if a != b {
    panic(fmt.Sprintf("%d is not equal to %d", a, b))
  }
}

func main() {
  assertEqualInt(2, minSubArrayLen([]int{2, 3, 1, 2, 4, 3}, 7))
  assertEqualInt(2, minSubArrayLen([]int{2, 3, 1, 3, 4, 3}, 7))
  assertEqualInt(2, minSubArrayLen([]int{2, 3, 1, 3, 4, 3}, 5))
  assertEqualInt(3, minSubArrayLen([]int{2, 3, 4, 3, 2, 1}, 6))
  assertEqualInt(1, minSubArrayLen([]int{2, 3, 1, 2, 4, 3}, 4))
}