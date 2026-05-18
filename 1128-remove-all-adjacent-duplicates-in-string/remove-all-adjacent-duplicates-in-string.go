
func removeDuplicates(s string) string {
    result := []rune{}

    for _, Value := range s{
        
        if len(result) >0 &&  Value == result[len(result)-1]{
            result = result[:len(result)-1]
        }else{
            result = append(result, Value)
        }
    }
    return string(result)
    
}