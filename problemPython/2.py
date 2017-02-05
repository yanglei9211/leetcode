# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """

        ans = []
        while l1<>None or l2 <> None:
        	k = 0
        	if l1 <> None :
        		k = l1.val
        		l1 = l1.next
        	if l2 <> None :
        		k += l2.val
        		l2 = l2.next
        	ans.append(k)

        lenNum = len(ans)
        """for i in range (0,lenNum):
        	if i == lenNum -1 and ans[i] > 10 : ans.append(0)
        	if (ans[i] > 10):
        		ans[i+1] = ans[i+1] + ans[i] // 10
        		ans[i] = ans[i] % 10
        """

        rtype = ListNode(ans[lenNum-1])
        i = lenNum-2

        while i>=0:
        	adds = ListNode(ans[i])
        	adds.next = rtype
        	rtype = adds
        	i-=1
        tmp = rtype
        while (tmp <> None):
        	if tmp.val >= 10:
        		tmp.val -= 10
        		if (tmp.next <> None):
        			tmp.next.val += 1
        		else:
        			adds = ListNode(1)
        			tmp.next = adds
        	tmp = tmp.next
        """
        while (rtype <> None):
        	print rtype.val
        	rtype = rtype.next
        """
        return rtype