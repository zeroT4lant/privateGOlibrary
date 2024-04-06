package main

func main() {

}

// 3 Задача
// AAAAAAABBBCCXYZDDDEEEAAABBBBBBBBBBBBBBB
// Подсчитать совпадения
func RLE(in string) string {
	s := []byte(in + " ")
	res := []byte{}
	cnt := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt++
			continue
		} else if cnt == 1 {
			res = append(res, s[i])
		} else {
			res = append(res, s[i], byte(cnt))
			cnt = 1
		}
	}
	return string(res)
}
