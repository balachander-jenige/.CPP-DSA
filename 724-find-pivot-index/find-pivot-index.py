class Solution:

    def pivotIndex(self, nums: List[int]) -> int:
        total = sum(nums)
        left = 0

        for i, num in enumerate(nums):
            # total - left - num gives the sum of elements to the right of index i
            if left == total - left - num:
                return i
            left += num

        return -1