func maxProduct(nums []int) int {

   res := nums[0]
   maxend :=nums[0]
   minend :=nums[0]

   for i:=1; i<len(nums); i++{
      v1 :=nums[i]
      v2 := maxend * nums[i]
      v3 := minend * nums[i]

      maxend = max(v1,v2,v3)
      minend = min(v1,v2,v3)

      res = max(res,maxend,minend)

   }
 return res
    
}