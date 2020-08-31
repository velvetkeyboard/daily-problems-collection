class Solution:
    def isValid(self, s):
        # Fill this in.
        expected = {
            '(': ')',
            '{': '}',
            '[': ']',
            }
        if len(s):
            if s.count('(') != s.count(')') or \
                s.count('{') != s.count('}') or \
                s.count('[') != s.count(']'):
                return False
            else:
                expecting = []
                for c in s:
                    if c in expected.keys():
                        expecting.append(expected[c])
                    else:
                        if c == expecting[-1]:
                            expecting.pop()
                return len(expecting) == 0
        else:
            return True

# Test Program
assert Solution().isValid("()(){(())") == False
assert Solution().isValid("") == True
assert Solution().isValid("([{}])()") == True
assert Solution().isValid("([)()") == False
assert Solution().isValid("([)]()") == False
assert Solution().isValid("()())") == False