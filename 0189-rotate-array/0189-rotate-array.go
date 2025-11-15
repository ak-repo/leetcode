func rotate(nums []int, k int)  {


    j:=len(nums)-1

    for k > 0{
        last:= nums[j]
        for i:=j; i>0;i--{
            nums[i] = nums[i-1]
        }
        nums[0] = last
        k--
    }


  
    
}