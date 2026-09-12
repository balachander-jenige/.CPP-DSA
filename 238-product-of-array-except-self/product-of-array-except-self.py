class Solution:

    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        prefix = [1] * n
        suffix = [1] * n

        # 1. Fill Prefix Array (products of elements to the left)
        for i in range(1, n):
            prefix[i] = prefix[i - 1] * nums[i - 1]

        # 2. Fill Suffix Array (products of elements to the right)
        for i in range(n - 2, -1, -1):
            suffix[i] = suffix[i + 1] * nums[i + 1]

        # 3. Combine both
        return [prefix[i] * suffix[i] for i in range(n)]