func missingNumber(nums []int) int {
    for i:=0;i<=len(nums);i++{
        if slices.Index(nums,i) == -1{
            return i
        }
    }
    return -1   
    
}