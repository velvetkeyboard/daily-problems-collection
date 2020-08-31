class Solution:
    def lengthOfLongestSubstring(self, s):
        ret = 0
        acc = ''
        for c in s:
            if c in acc:
              if len(acc) > ret:
                  ret = len(acc)
              acc = c
            else:
              acc += c
        return ret

assert Solution().lengthOfLongestSubstring('abrkaabcdefghijjxxx') == 10
assert Solution().lengthOfLongestSubstring('xyzaab') == 4
assert Solution().lengthOfLongestSubstring('xyzzaab') == 3
# 10