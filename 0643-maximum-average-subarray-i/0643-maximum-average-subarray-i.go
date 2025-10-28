func findMaxAverage(nums []int, k int) float64 {
    var wSum, mSum float64

    for i := 0; i < k; i++ {
        wSum += float64(nums[i])
    }
    mSum = wSum

    for i := k; i < len(nums); i++ {
        wSum = wSum - float64(nums[i-k]) + float64(nums[i])
        if wSum > mSum {
            mSum = wSum
        }
    }

    return mSum / float64(k)
}
