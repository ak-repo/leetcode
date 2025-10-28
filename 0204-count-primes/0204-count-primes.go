func countPrimes(n int) int {

    if n <=2{
        return 0
    }
    isComposite := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !isComposite[i] {
			for j := i * i; j <= n; j += i {
				isComposite[j] = true
			}
		}
	}

primes:= 0
	for i := 2; i < n; i++ {
		if !isComposite[i] {
			primes++
		}
	}

    return primes
    
}