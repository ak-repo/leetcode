func twoSum(nums []int, target int) []int {


    for i,v:= range nums{
        want:= target - v
        for j,v2:= range nums{
            if v2 == want && i!= j{
                return []int{i,j}
            }
        }

    }

    return []int{-1,-1}
    
}