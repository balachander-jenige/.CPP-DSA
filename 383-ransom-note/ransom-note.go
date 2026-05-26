func canConstruct(ransomNote string, magazine string) bool {

   m := make(map[rune]int)
   for _,v := range magazine{
    m[v] +=1
   }
   for _,v := range ransomNote{
    m[v] -=1
    if m[v] < 0{
        return false
    }
   }
   return true
}