package main

import "fmt"

func stringRotation(input1, input2 string) bool {
	r := input2 + input2

	len1 := len(input1)
	len2 := len(r)
	i := 0
	j := 0

	if len1+len2 == 0 {
		return true
	}

	for i < len1 {
		j = 0
		for j < len1 {
			if i+j < len2 {
				if r[i+j] != input1[j] {
					break
				}

			} else {
				break
			}

			j++
		}

		if j == len1 {
			return true
		}
		i++
	}
	return false
}
func main() {
	inputs1 := []string{"waterbottle", "hellomynameis", "waterbottle", "waterbottle",
		" ", "",
	}
	inputs2 := []string{"erbottlewat", "nameishellomy", "water", "elttobretaw",
		" ", "",
	}
	outputs := []bool{true, true, false, false,
		true, true,
	}
	var sol bool
	var ans bool

	for i := 0; i < len(outputs); i++ {
		sol = stringRotation(inputs1[i], inputs2[i])
		ans = outputs[i]

		if sol != ans {
			fmt.Println(inputs1[i], inputs2[i], sol, ans)
		}
	}
}
