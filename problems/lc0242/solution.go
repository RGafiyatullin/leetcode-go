package leetcode

func isAnagram(s string, t string) bool {
	s_digest := digest(s)
	t_digest := digest(t)

	for rune, count := range s_digest {
		if t_digest[rune] != count {
			return false
		}
		delete(t_digest, rune)
	}

	return len(t_digest) == 0
}

func digest(in string) map[rune]uint {
	m := map[rune]uint{}

	for _, char := range in {
		m[char]++
	}

	return m
}
