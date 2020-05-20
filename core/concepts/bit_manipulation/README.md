1. Get ith bit: `(n & (1 << i)) != 0`
2. Set ith bit: `n | (1 << i)`
3. Clear ith bit: `n & (~(1 << i))`
4. Check if n is a power of 2: `(n != 0 ) && ( n & (n-1)) == 0 `
5. Count the number of ones in the binary representation of n:
```Go
    count := 0

	for x != 0 {
		x = x & (x - 1)
		count++
	}
```
6. Check if the ith bit is set in the binary form of n: `(n & (1 << i)) != 0`
7. Generate all the possible subsets of a set
```C++
    for(int i = 0;i < (1 << N); ++i)
        {
            for(int j = 0;j < N;++j)
                if(i & (1 << j))
                    cout << A[j] << ‘ ‘;
            cout << endl;
```
8. Find the largest power of 2 (most significant bit in binary form), which is less than or equal to the given number n.
```Go
    func LargestPowerOfTwo16(n int16) int16 {
        n = n | (n >> 1)
        n = n | (n >> 2)
        n = n | (n >> 4)
        n = n | (n >> 8)

        return (n + 1) >> 1
    }
```
9. Returns the rightmost 1 in binary representation of x:
    - `x ^ ( x & (x-1))`
    - `x & (-x)`


https://www.youtube.com/watch?v=NLKQEOgBAnw
https://www.hackerearth.com/practice/basic-programming/bit-manipulation/basics-of-bit-manipulation/tutorial/