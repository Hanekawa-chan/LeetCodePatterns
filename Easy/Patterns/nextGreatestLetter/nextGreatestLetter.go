package nextGreatestLetter

import "fmt"

func Example() {
	fmt.Println("\nFind Smallest Letter Greater Than Target")
	nums := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	fmt.Println("nums =", nums)
	result := nextGreatestLetter(nums, 'c')
	fmt.Println("result =", result)
}

func nextGreatestLetter(letters []byte, target byte) byte {
	index := search(letters, target)
	return letters[index]
}

func search(let []byte, t byte) int {
	//получаем инфу о том есть ли у нас вообще числа меньше таргета
	n := 0
	if let[n] > t {
		return n
	}
	//получаем инфу о том равно ли первое число таргету
	if let[n] == t {
		//ищем число больше таргета
		for let[n] <= t {
			n++
		}
		return n
	}
	//проверяем не больше ли последнее число таргета
	n = len(let) - 1
	if let[n] <= t {
		return 0
	}
	//получаем такую картину мира let{<t,...,>t}
	//ищем есть ли вообще в середине таргет и избавляемся от краев
	arr := let
	n = (n + 1) / 2
	num := n
	if arr[n] == t {
		for let[n] == t {
			n++
		}
		return n
	}
	for arr[n] != t {
		//чо делать, если нет таргета?
		if len(arr) == 1 {
			if let[num] > t {
				m := num - 1
				for let[m] > t {
					m -= 1
				}
				return m + 1
			}
			if let[num] < t {
				m := num + 1
				for let[m] < t {
					m += 1
				}
				return m
			}
		}
		//чо делать, если таргет меньше?
		if arr[n] > t {
			arr = arr[0:n]
			n = len(arr) / 2
			num -= n + 1
			//чо делать, если таргет больше?
		} else {
			arr = arr[n+1:]
			n = len(arr) / 2
			num += n + 1
			//чо делать, если нет таргета?
			if len(arr) == 0 {
				if let[num] > t {
					m := num - 1
					for let[m] > t {
						m -= 1
					}
					return m + 1
				}
				if let[num] < t {
					m := num + 1
					for let[m] < t {
						m += 1
					}
					return m
				}
			}
		}
	}
	if num < 0 {
		num = 1
	}
	//чо делать, если таргет больше?
	if let[num] == t {
		for let[num] == t {
			num++
		}
		return num
	}
	//чо делать, если таргет меньше?
	if let[num] > t {
		for let[num] != t {
			num--
		}
		//чо делать, если таргет больше?
	} else {
		for let[num] != t {
			num++
		}
	}
	for let[num] == t {
		num++
	}
	return num
}
