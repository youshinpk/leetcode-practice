func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
	
	f0, f1 = 1, 2
	
	for i := 3; i <= n; i++{
		f2 := f0 + f1
		f0 = f1
		f1 = f2
	}
	return f1
}
