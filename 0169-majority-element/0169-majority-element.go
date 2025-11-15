func majorityElement(nums []int) int {
    c:= make(map[int]int)

    for _,v:= range nums{
        c[v]++
    }
    
    most:=0
    index:=0
    for k,v:= range c{
        if v >most{
            most = v
            index = k
        }
    }

    return index
    
}