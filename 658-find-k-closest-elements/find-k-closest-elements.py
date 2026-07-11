import heapq
from typing import List

class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        maxheap = []

        for num in arr:
            heapq.heappush_max(maxheap, (abs(num - x), num))
            if len(maxheap)>k:
                heapq.heappop_max(maxheap)

        ans = []

        for _ in range(k):
            ans.append(heapq.heappop_max(maxheap)[1])

        return sorted(ans)