class Solution:

    def subarraySum(self, nums: List[int], k: int) -> int:
        count = 0
        prefixsum = 0
        prefixcount = {0: 1}

        for num in nums:
            prefixsum += num

            # Checks if (prefixsum - k) exists in the dictionary and adds its frequency
            if (prefixsum - k) in prefixcount:
                count += prefixcount[prefixsum - k]

            # Python's dict.get() handles missing keys safely (defaulting to 0)
            prefixcount[prefixsum] = prefixcount.get(prefixsum, 0) + 1

        return count