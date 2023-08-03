package main

import (
	"fmt"
	"testing"
)

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
// Example 1:

// Input: s = "leetcode"
// Output: 0
// Example 2:

// Input: s = "loveleetcode"
// Output: 2
// Example 3:

// Input: s = "aabb"
// Output: -1

func firstUniqChar1(s string) int {
	d := map[byte]int{}

	// Count each character.
	for i := 0; i < len(s); i++ {
		d[s[i]]++
	}

	// Find the first unique character and return.
	for i := 0; i < len(s); i++ {
		if d[s[i]] == 1 {
			return i
		}
	}

	// If there's no unique character then return -1.
	return -1
}
func firstUniqChar2(s string) int {
	// This is a constant space allocation (ie: always 26)
	var counts = make([]int, 26)

	// Insert all the character's count into counts array
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	// Find the first element with count 1
	for i := 0; i < len(s); i++ {
		if counts[s[i]-'a'] == 1 {
			// The reason for - 'A', is that it "shifts" the ascii/unicode value so that A - Z have values 0 - 25.
			// And are thus more suitable as an array index.
			//'A' - 'A' == 0
			// 'B' - 'A' == 1
			// 'C' - 'A' == 2
			return i
		}
	}

	return -1
}

func TestFirstUniqChar(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
	for _, c := range cases {
		got := firstUniqChar1(c.in)
		if got != c.want {
			t.Errorf("FirstUniqChar(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}

func main() {
	fmt.Println(firstUniqChar1("leetcode"))
}
