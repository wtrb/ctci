package bitmnp

// LargestPowerOfTwo16 returns the largest power of 2 (most significant bit in binary form),
// which is less than or equal to the given number n (16 bits).
// Time comeplexity: O(1)
// Space comeplexity: O(1)
func LargestPowerOfTwo16(n int16) int16 {
	n = n | (n >> 1)
	n = n | (n >> 2)
	n = n | (n >> 4)
	n = n | (n >> 8)

	return (n + 1) >> 1
}

// LargestPowerOfTwo32 returns the largest power of 2 (most significant bit in binary form),
// which is less than or equal to the given number n (32 bits).
// Time comeplexity: O(1)
// Space comeplexity: O(1)
func LargestPowerOfTwo32(n int32) int32 {
	n = n | (n >> 1)
	n = n | (n >> 2)
	n = n | (n >> 4)
	n = n | (n >> 8)
	n = n | (n >> 16)

	return (n + 1) >> 1
}
