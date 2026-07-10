import heapq

class Solution:
    def kClosest(self, points, k):
        maxheap = []

        for p in points:
            distance = p[0] * p[0] + p[1] * p[1]

            heapq.heappush_max(maxheap, (distance, p))

            if len(maxheap) > k:
                heapq.heappop_max(maxheap)

        return [point for distance, point in maxheap]

        