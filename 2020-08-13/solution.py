# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2, c = 0):
        def convert_to_str(list_node):
            if list_node.next:
                return str(list_node.val) + convert_to_str(list_node.next)
            return str(list_node.val)
        l1 = convert_to_str(l1)[::-1]
        l2 = convert_to_str(l2)[::-1]
        l3 = int(l1) + int(l2)
        l3 = str(l3)
        l3 = l3[::-1]
        ret = ListNode(0)
        pivot = ret
        for i, digit in enumerate(l3):
            pivot.val = int(digit)
            if i < len(l3) - 1:
                pivot.next = ListNode(0)
                pivot = pivot.next
        return ret

l1 = ListNode(2)
l1.next = ListNode(4)
l1.next.next = ListNode(3)

l2 = ListNode(5)
l2.next = ListNode(6)
l2.next.next = ListNode(4)

result = Solution().addTwoNumbers(l1, l2)
while result:
    print(result.val, end="")
    result = result.next
# 7 0 8
