class Solution(object):
    def cov(self,s):
        if s == 'I': return 1
        elif s == 'V': return 5
        elif s == 'X': return 10
        elif s == 'L': return 50
        elif s == 'C': return 100
        elif s == 'D': return 500
        elif s == 'M': return 1000

    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        lenS = len(s)
        a = [0 for i in range(lenS + 2)]
        for i in range(lenS):
            a[i] = self.cov(s[i])

        res = 0
        for i in range(lenS):
            if a[i] < a[i+1]:res -= a[i]
            else : res += a[i]
        return res