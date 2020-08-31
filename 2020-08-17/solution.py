class Solution: 
    def getRange(self, arr, target):
        # Fill this in.
        try:
            return [
                arr.index(target),
                len(arr) - arr[::-1].index(target) - 1,
            ]
        except ValueError:
            return [-1, -1]
  
# Test program 
arr = [1, 2, 2, 2, 2, 3, 4, 7, 8, 8] 
x = 2
assert Solution().getRange(arr, x) == [1, 4]

arr = [1, 2, 2, 2, 2, 3, 4, 7, 8, 8] 
x = 2
assert Solution().getRange(arr, x) == [1, 4]

arr = [1,3,3,5,7,8,9,9,9,15]
x = 9
assert Solution().getRange(arr, x) == [6,8]

arr = [100, 150, 150, 153]
x = 150
assert Solution().getRange(arr, x) == [1,2]

arr = [1, 2, 3, 4, 5, 6, 10]
x = 9
assert Solution().getRange(arr, x) == [-1, -1]