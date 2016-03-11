#Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

def printLink(head):
    while head<>None:
        print head.val,'->'
        head = head.next

def creadLink(a):
    if a == []: return None

    head = ListNode(a[0])
    node = head
    for i in a[1:] :
        node2 = ListNode(i)
        node.next = node2
        node = node.next
    return head
def countLink(head):
    cnt = 0
    while head:
        head = head.next
        cnt += 1
    return cnt

class Solution(object):
    def mergeTwoLists(self, l1, l2):
		h1 = l1
		h2 = l2
		resH = ListNode(-1)
		tmp = resH
		while h1 and h2:
			if h1.val < h2.val:
				node = ListNode(h1.val)
				h1 = h1.next
				tmp.next = node
			else:
				node = ListNode(h2.val)
				h2 = h2.next
				tmp.next = node
			tmp = tmp.next
		if h1 : tmp.next = h1
		if h2 : tmp.next = h2
		resH = resH.next
		return resH
test = Solution()
a = [2]
b = [1]
la = creadLink(a)
lb = creadLink(b)
printLink(test.mergeTwoLists(la,lb))



