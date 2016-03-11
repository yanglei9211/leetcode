class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        if (str == '') : return 0
        res = 0
        i = 0
        lenS = len(str)
        while (str[i] == ' ') : i+=1
        str = str[i:lenS]

        if (str[0] == '-' or str[0] == '+'): sstr = str[1:lenS]
        else : sstr = str


        if sstr == '' :return 0

        for i in sstr:
            if (i > '9' or i < '0'): break
            res= res * 10;
            res = res + (ord(i) - ord('0'))

        MAX_INT = 2147483647
        MIN_INT = -2147483648
        if (str[0] == '-'): res = - res
        if (res > MAX_INT): res = MAX_INT
        if (res < MIN_INT): res = MIN_INT
        return res