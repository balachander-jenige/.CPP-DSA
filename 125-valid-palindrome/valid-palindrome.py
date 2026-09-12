class Solution:
    def isPalindrome(self, s: str) -> bool:
        fs=""
        for i in s.lower():
            if i.isalnum():
                fs+=i
        
        return fs==fs[::-1]