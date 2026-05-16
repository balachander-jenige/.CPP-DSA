func collide(value int, res []int) []int {
    if len(res) ==0{
        res=append(res, value)
        return res
    }
    i:=len(res)-1
    for (i>=0){
        if res[i] < 0 {
			break
		}
        if res[i] > -(value){
            return res;
        }else if (res[i] == -(value)){
            res=res[:i]
            return res
        }else{
            res=res[:i]
            i--
        }
    }
    res=append(res,value)
    return res
}


func asteroidCollision(asteroids []int) []int {
    var res []int;
    for _,value := range(asteroids) {
        if value < 0{
            res=collide(value,res)
        }else{
            res=append(res,value)
            }
    }
    return res;
}