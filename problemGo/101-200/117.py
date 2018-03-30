# Definition for binary tree with next pointer.
# class TreeLinkNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None

class Solution:
    # @param root, a tree link node
    # @return nothing
    # def init(self, rt):
    #     if not rt:
    #         return
    #     self.init(rt.left)
    #     rt.next = None
    #     self.init(rt.right)

    def connect(self, root):
        q = fake = TreeLinkNode(0)
        node = root
        while node:
            q.next = node.left
            if q.next:
                q = q.next
            q.next = node.right
            if q.next:
                q = q.next
            node = node.next
            if not node:
                q = fake
                node = fake.next
