class Solution {
public:
    bool canBeEqual(vector<int>& target, vector<int>& arr) {
        if(target.size()!=arr.size()){
            return false;
        }

        sort(target.begin(),target.end());
        sort(arr.begin(),arr.end());
        if(target==arr){
            return true;
        }
        /*for(int i=0;i<target.size();i++){
            if(target[i]!=arr[i]){
                return false;
            }
        }
        return true;*/
        return 0;

        
    }
};