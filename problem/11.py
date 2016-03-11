class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        res = 0
        lp = 0
        rp = len(height)-1
        while (lp < rp):
            res = max(res,min(height[lp],height[rp]) * (rp - lp))
            if (height[lp] < height[rp]): lp+=1
            else: rp-=1


        return res