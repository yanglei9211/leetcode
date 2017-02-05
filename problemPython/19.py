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
    def removeNthFromEnd(self, head, n):
        num = countLink(head)
        n = num - n

        res = ListNode(-1)
        res.next = head
        head = res

        while n > 0 and head <> None:
            n -= 1
            if head.next:
                head = head.next
        if head and head.next:
            head.next = head.next.next
        res = res.next
        printLink(res)
        return res