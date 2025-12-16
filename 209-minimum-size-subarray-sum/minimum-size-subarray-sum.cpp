class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int sum=0;
        int min=INT_MAX;
        bool t=true;
        int i=0,j=0;
        int n=nums.size();
        while(j<n){
            sum+=nums[j];
            while(sum >= target && i<n){
                t=false;
                if(min>(j-i+1)){
                    min=j-i+1;
                }
                sum=sum-nums[i];
                i+=1;
            }
            j+=1;
        }
        if (t)  {
            return 0;

        }else{
              return min;

        } 
        
    }
};