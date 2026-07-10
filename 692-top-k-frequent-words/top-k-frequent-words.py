from collections import Counter
import heapq

class Word:
    def __init__(self, word, freq):
        self.word = word
        self.freq = freq

    # __lt__ means "less than"
    # heapq calls this automatically whenever it needs
    # to compare two Word objects.
    def __lt__(self, other):

        # If frequencies are the same,
        # treat the lexicographically LARGER word as "smaller"
        # so it gets removed first from the min-heap.
        if self.freq == other.freq:
            return self.word > other.word

        # Otherwise, the word with the smaller frequency
        # has higher priority to be removed.
        return self.freq < other.freq
      

class Solution:
    def topKFrequent(self, words, k):
        freq = Counter(words)

        minheap = []

        for word, count in freq.items():

            # Push a Word object into the heap.
            # heapq may call __lt__ internally to place it
            # in the correct position.
            heapq.heappush(minheap, Word(word, count))

            if len(minheap) > k:
                # heapq compares objects using __lt__
                # to find the smallest element.
                heapq.heappop(minheap)

        result = []

        # Popping returns elements from smallest to largest,
        # so we reverse the final list.
        while minheap:
            result.append(heapq.heappop(minheap).word)

        return result[::-1]