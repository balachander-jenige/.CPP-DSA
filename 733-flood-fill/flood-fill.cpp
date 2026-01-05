class Solution {

    void dfs(int i, int j, int k, int color, vector<vector<int>>& image){
        if(image[i][j]!=k) return;

        image[i][j]=color;
        if(i-1>=0) dfs(i-1, j, k, color, image);
        if(j-1>=0) dfs(i, j-1, k, color, image);
        if(i+1<image.size()) dfs(i+1, j, k, color, image);
        if(j+1<image[0].size()) dfs(i, j+1, k, color, image);
    }
public:
    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int color) {
        if(image[sr][sc]==color) return image;
        dfs(sr, sc, image[sr][sc], color, image);
        return image;
        
    }
};