func countPrimes(n int) int {

    if n <=2{
        return 0
    }
    isComposite := make([]bool, n+1)
	primes := []int{}
	for i := 2; i*i <= n; i++ {
		if !isComposite[i] {
			for j := i * i; j <= n; j += i {
				isComposite[j] = true
			}
		}
	}

	for i := 2; i < n; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
		}
	}

    return len(primes)
    
}