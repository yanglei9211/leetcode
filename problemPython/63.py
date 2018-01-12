#!/usr/bin/env python
# encoding: utf-8


class Solution(object):
    def uniquePathsWithObstacles(self, obstacleGrid):
        """
        :type obstacleGrid: List[List[int]]
        :rtype: int
        """
        lx = len(obstacleGrid)
        if not lx:
            return 1
        ly = len(obstacleGrid[0])
        if not ly:
            return 1

        dp = []
        for i in range(0, lx):
            t = []
            for j in range(0, ly):
                t.append(0)
            dp.append(t)

        # for i in range(0, ly):
        #     if not obstacleGrid[0][i]:
        #         dp[0][i] = 1
        #
        # for i in range(1, lx):
        #     if not obstacleGrid[i][0]:
        #         dp[i][0] = 1
        for i in range(0, lx):
            for j in range(0, ly):
                if obstacleGrid[i][j]:
                    continue

                if i == 0 and j == 0:
                    dp[i][j] = 1
                    continue
                elif i == 0:
                    dp[i][j] = dp[i][j-1]
                elif j == 0:
                    dp[i][j] = dp[i-1][j]
                else:
                    if not obstacleGrid[i-1][j] and not obstacleGrid[i][j-1]:
                        dp[i][j] = dp[i-1][j] + dp[i][j-1]
                    elif not obstacleGrid[i-1][j] and obstacleGrid[i][j-1]:
                        dp[i][j] = dp[i-1][j]
                    elif obstacleGrid[i-1][j] and not obstacleGrid[i][j-1]:
                        dp[i][j] = dp[i][j-1]
                    else:
                        dp[i][j] = 0
        return dp[lx-1][ly-1]

s = Solution()
# a = [[0,0,0], [0,1,0], [0,0,0]]
a = [[0], [1]]
print s.uniquePathsWithObstacles(a)
