class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_map<int,int> f;\
        int maxlen=0;
        int res=0;
        int low=0;
        for(int i=0;i<s.size();i++){
            f[s[i]]++;
            while(f[s[i]]>1)
            {   
                f[s[low]]--;
                low++;
                
            }
            maxlen=i-low+1;
            res=max(res,maxlen);
        }
        return res;
        
    }
};