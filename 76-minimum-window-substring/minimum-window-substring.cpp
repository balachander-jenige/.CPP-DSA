class Solution {
public:
    string minWindow(string s, string t) {
        if (s.size() < t.size()) return "";

        unordered_map<char, int> freq;
        for (char c : t) {
            freq[c]++;
        }

        int left = 0, right = 0;
        int required = t.size();
        int minLen = INT_MAX;
        int start = 0;

        while (right < s.size()) {
            char rc = s[right];

            // If character is needed
            if (freq[rc] > 0) {
                required--;
            }
            freq[rc]--;
            right++;

            // Try shrinking the window
            while (required == 0) {
                if (right - left < minLen) {
                    minLen = right - left;
                    start = left;
                }

                char lc = s[left];
                freq[lc]++;

                // If removing this char breaks the requirement
                if (freq[lc] > 0) {
                    required++;
                }
                left++;
            }
        }

        return minLen == INT_MAX ? "" : s.substr(start, minLen);
    }
};
