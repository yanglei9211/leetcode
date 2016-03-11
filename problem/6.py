class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if (numRows == 1) : return s

        res = ''
        lenS = len(s)
        step = numRows * 2 - 2
        res = s[::step]

        for i in range(1,numRows - 1):
            for j in range(i,lenS,step):
                res = res + s[j]
                if j + (step - i * 2) < lenS :
                    res = res + s[j + step - i * 2]

        i = numRows -1
        res = res + s[i::step]
        return res