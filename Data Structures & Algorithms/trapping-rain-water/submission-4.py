class Solution:
	def trap(self, height: List[int]) -> int:
		leftMax = [0]
		rightMax = [0]
		for i, v in enumerate(height):
			if i == 0:
				leftMax.append(v)
				continue
			leftMax.append(max(leftMax[-1], v))
		for i in range(len(height)-1, -1, -1):
			if i == len(height)-1:
				rightMax.append(height[i])
				continue
			rightMax.append(max(rightMax[-1], height[i]))
		
		rightMax.reverse()
		water = 0
		for i, v in enumerate(height):
			if leftMax[i] > v and rightMax[i] > v:
				water += (min(leftMax[i], rightMax[i]) - v)
		return water
			