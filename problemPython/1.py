class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        """index = {};
        idx = 0;
        for i in nums:
        	index[i] = idx
        	idx+=1 	"""
        lenNums = len(nums)
        for i in range(0,lenNums):
        	ansx = nums[i];
        	ansy = target - ansx;
        	res = []

        	if (ansy in nums):
        		for j in range(0,lenNums):
        			if ansx == nums[j]:
        				res.append(j)
        				break
        		for j in range(0,lenNums):
        			if ansy == nums[j] and res[0] <> j :
        				res.append(j)
        				break

        		if len(res)>1 : return res
        return res
