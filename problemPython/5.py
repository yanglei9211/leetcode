class Solution(object):
    def pre(self,s):
        lenS = len(s)
        i = 0
        res = []
        while i < lenS:
            res.append('#')
            res.append(s[i])
            i+=1
        res.append('#')
        return res
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        ss = self.pre(s)

        lenS = len(ss)


        p = [1 for i in range(lenS+1)]
        id = 0
        mx = 0
        i = 0
        while i < lenS:
            if mx > i :
                p[i] = min(p[id*2-i],p[id]+id-i)
            else : p[i] = 1
            while i - p[i] >= 0 and i+p[i] < lenS and ss[i+p[i]] == ss[i-p[i]]: p[i] += 1
            if (p[i] + i > mx):
                mx = p[i] + i
                id = i
            i += 1
        ans = 0
        i = 0
        id = 0

        while i < lenS:
            if (ans < p[i]):
                ans = p[i] - 1
                id = i // 2
            i += 1

        id = id - ans // 2

        return s[id:id+ans]


