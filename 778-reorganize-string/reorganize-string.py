from collections import Counter
import heapq


class Solution:
    def reorganizeString(self, s: str) -> str:
        freq = Counter(s)

        maxheap = []

        for char, count in freq.items():
            heapq.heappush_max(maxheap, (count, char))

        res = []

        prev_count = 0
        prev_char = ""

        while maxheap:
            count, char = heapq.heappop_max(maxheap)

            # Cannot use the same character consecutively
            if char == prev_char:
                if not maxheap:
                    return ""

                count2, char2 = heapq.heappop_max(maxheap)

                res.append(char2)

                count2 -= 1

                if count2 > 0:
                    heapq.heappush_max(maxheap, (count2, char2))

                # Put the previous character back
                heapq.heappush_max(maxheap, (count, char))

                prev_char = char2

            else:
                res.append(char)

                count -= 1

                if count > 0:
                    heapq.heappush_max(maxheap, (count, char))

                prev_char = char

        return "".join(res)