func findNumbers(nums []int) int {
    //without std

	var count int

    for _,num:= range nums{
        digits:=0
        for num>0{
            num/=10
            digits++
        }
        if digits % 2 == 0{
            count++
        }
    }
    return count
}
