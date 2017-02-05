class Solution(object):
	letter = [[' '],[],['a','b','c'],['d','e','f'],['g','h','i'],['j','k','l'],['m','n','o'],['p','q','r','s'],['t','u','v'],['w','x','y','z'],['*'],['#']];
	ans = []
	Nstring = ""
	def dfs(self,s,i):
		if i == len(s):
			if self.Nstring <> '':
				self.ans.append(self.Nstring)
			#print self.Nstring
			return

		k = ord(s[i]) - ord('0')
		for j in self.letter[k]:
			self.Nstring += j
			self.dfs(s,i+1)
			self.Nstring = self.Nstring[:-1]




	def letterCombinations(self, digits):
		self.ans = []
		self.Nstring = ""
		self.dfs(digits,0)

		return self.ans