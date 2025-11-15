func rotate(nums []int, k int)  {


for i:=0;i<k;i++{
    r(nums)
}
      

    
}
func r(nums []int){
    j:=len(nums)-1

      last:= nums[j]
        for i:=j; i>0;i--{
            nums[i] = nums[i-1]
        }
        nums[0] = last
}