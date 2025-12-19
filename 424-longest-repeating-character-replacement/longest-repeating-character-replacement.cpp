class Solution {
public:
    int characterReplacement(string s, int k) {
        int low=0;
        unordered_map<char,int> f;
        int maxf=0;
        int maxl=0;
        for(int i=0;i<s.size();i++){
            f[s[i]]++;
            maxf=max(maxf,f[s[i]]);
            while((i-low+1)-maxf>k)
            {
                f[s[low]]--;
                low++;

            }
            maxl=max(maxl,(i-low+1));


        }
        return maxl;
        
    }
};