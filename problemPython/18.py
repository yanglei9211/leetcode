from itertools import *
from collections import *
class Solution():
	def fourSum(self, nums, target):
		two_sum = defaultdict(list)
		ret = set()
		for (id1,val1),(id2,val2) in combinations(enumerate(nums),2):
			#print id1,val1,id2,val2
			two_sum[val1 + val2].append({id1,id2})
		keys = two_sum.keys()
		for key in keys:
			if two_sum[key] and two_sum[target - key]:
				for pair1 in two_sum[key]:
					for pair2 in two_sum[target - key]:
						if pair1.isdisjoint(pair2):
							ret.add(tuple(sorted([nums[i] for i in pair1 | pair2])))
				del two_sum[key]
				if key <> target - key:
					del two_sum[target - key]

		return [list(i) for i in ret]