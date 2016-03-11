class Solution(object):
		def threeSumClosest(self, nums, target):
			nums.sort()
			length = len(nums)
			temp = 1<<28
			res = 0
			for i in range(length):
				j = i+1
				k = length -1
				while (j < k):
					if j == i : j += 1
					if k == i : k -= 1
					if k == j:
						k -= 1
						continue
					sum = nums[i] + nums[j] + nums[k]
					if (abs(sum-target)<temp):
						temp = abs(sum-target)
						res = sum
					if (sum > target): k -= 1
					else: j += 1

			return res