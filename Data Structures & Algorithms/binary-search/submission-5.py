class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums)-1
        mid = int(left + (right - left)/2)

        while nums[mid] != target and left <= right:
            print(left, right, mid)
            if nums[mid] < target :
                left = mid+1
            else:
                right = mid-1

            mid = int(left + (right - left)/2)

        if left > right:
            return -1
        return mid