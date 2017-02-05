class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        s = str(x)
        lenS = len(s)

        for i in range(0,lenS // 2):
             if (s[i] != s[lenS - 1 - i]): return False
        return True