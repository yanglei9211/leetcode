class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()

        n = len(nums)
        if n <= 2: return []

        res = []
        i = 0
        while i < n-2 :
            l = i + 1
            r = n - 1
            while(l < r):
                if (nums[i] + nums[l] + nums[r] > 0):
                    r = r - 1
                elif (nums[i] + nums[l] + nums[r] < 0):
                    l = l + 1
                else:
                    #print nums[i],nums[l],nums[r]
                    tmp = [nums[i],nums[l],nums[r]]
                    res.append(tmp)
                    l += 1
                    r -= 1

                    while l < r and nums[l] == nums[l - 1]: l += 1
                    while l < r and nums[r] == nums[r + 1]: r -= 1
                    while i < n - 2 and nums[i] == nums[i + 1]:
                        i += 1
            i += 1


        return res