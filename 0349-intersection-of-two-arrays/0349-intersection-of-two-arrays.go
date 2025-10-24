func intersection(nums1, nums2 []int) []int {
    m := make(map[int]bool)
    for _, v := range nums1 {
        m[v] = true
    }

    res := make(map[int]bool)
    for _, v := range nums2 {
        if m[v] {
            res[v] = true
        }
    }

    out := []int{}
    for k := range res {
        out = append(out, k)
    }

    return out
}
