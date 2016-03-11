class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        l = len(s)
        a = ['x' for i in range(l+1)]
        pt = 0
        for i in s:
            if i == '(' or i == '[' or i =='{':
                a[pt] = i
                pt += 1
            elif i == ')':
                if pt > 0 and a[pt-1] == '(':
                    pt -= 1
                else: return False
            elif i == ']':
                if pt > 0 and a[pt-1] == '[':
                    pt -= 1
                else: return False
            elif i == '}':
                if pt > 0 and a[pt-1] == '{':
                    pt -= 1
                else: return False
        if pt == 0:
            return True
        else : return False