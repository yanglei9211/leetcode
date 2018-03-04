#!/usr/bin/env python
# encoding: utf-8


class Solution(object):
    def __init__(self):
        self.N = 205
        self.fac = [1] * self.N
        for i in range(1, self.N):
            self.fac[i] = self.fac[i-1] * i
        print self.fac[100]

    def uniquePaths(self, m, n):
        n -= 1
        m -= 1
        if not n or not m:
            return 1
        n1 = n+m
        n2 = n
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        x1 = self.fac[n1] / self.fac[n2]
        x1 = x1 / self.fac[n1-n2]
        return x1

s = Solution()
print s.uniquePaths(80, 80)
