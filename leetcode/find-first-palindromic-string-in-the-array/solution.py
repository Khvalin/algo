class Solution:
    def firstPalindrome(self, words: List[str]) -> str:
        def is_palindrome(w):
            i, j = 0, len(w) - 1

            while i < j:
                if w[i] != w[j]:
                    return False
                i += 1
                j -= 1
            return True

        for w in filter(is_palindrome, words):
            return w
        return ""
