class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        i = 0

        vis = []
        while (i < 256):
        	vis.append(0)
        	i += 1
        ns= []
        ans = 1
        i = 0
        j = 0
        lenS = len(s)
        if (lenS == 0) : return 0
        vis[ord(s[0])] = 1

        while (i<lenS and j < lenS):
            while (j < lenS-1 and vis[ord(s[j+1])] == 0):
        	   j += 1
        	   vis[ord(s[j])] = 1

            ans = max(ans, j - i + 1)

            vis[ord(s[i])] = 0
            i += 1
        return ans