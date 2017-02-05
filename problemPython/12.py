class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        nums = [1000,500,100,50,10,5,1]
        res = ''
        x = num // 1000

        for i in range(x) : res = res + 'M'
        num = num % 1000
        x = num // 100
        ti = 0
        if x == 9:
            res = res + 'CM'
        elif x == 4:
            res = res + 'CD'
        else:

            if (x>=5):
                ti = x - 5
                res = res + 'D'
            else:
                ti = x

            for i in range(ti) : res = res + 'C'

        num = num % 100
        x = num // 10
        if x == 9:
            res = res + 'XC'
        elif x == 4:
            res = res + 'XL'
        else:
            if (x >= 5) :
                ti = x - 5
                res = res + 'L'
            else : ti = x
            for i in range(ti) : res = res +'X'

        num = num % 10
        x = num
        if x == 9:
            res = res + 'IX'
        elif x == 4: res = res + 'IV'
        else :
            if (x >= 5):
                ti = x - 5
                res = res + 'V'
            else: ti = x
            for  i in range(ti) :res = res + 'I'
        return res




