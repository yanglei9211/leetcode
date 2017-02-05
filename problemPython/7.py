import math
class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        flag = 0
        if (x < 0) :
        	flag = -1
        	x = -x
        res = 0
        while x > 0:
        	res = res * 10
        	res = res + x % 10

        	x = x // 10
        if res > math.pow(2,31) :return 0
        if (flag == -1): res = -res
        return res