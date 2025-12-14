func twoSum(nums []int, target int) []int {

    m:= map[int]int{}

    for i,v:= range nums{
        diff:= target - v
        if j,ok:= m[diff]; ok{
            return []int{i,j}
        }
        m[v] = i
    }
    return []int{}
}