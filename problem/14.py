class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        lenS = len(strs)
        if lenS == 0: return ''
        if lenS == 1: return strs[0]
        minLen = 2<<29
        
        for i in range(lenS):
            minLen = min(len(strs[i]),minLen)

        ans = 0
        for i in range(minLen):
            for j in range(1,lenS):
                if strs[j][i] <> strs[j-1][i]:
                    ans = i
                    print ans
                    return strs[0][0:ans]

        return strs[0][0:minLen]