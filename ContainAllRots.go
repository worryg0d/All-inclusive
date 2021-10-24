package kata

import "fmt"

func ContainAllRots(strng string, arr []string) bool {
	str := strng

	for i := 0; i < len(strng); i++ {
		fmt.Println(str)
		isContain := true
		for _, word := range arr {
			if word == str {
				isContain = true
				break
			}
			isContain = false
		}
		if !isContain {
			return false
		}
		str = str[1:] + string(str[0])
	}
	return true
}
