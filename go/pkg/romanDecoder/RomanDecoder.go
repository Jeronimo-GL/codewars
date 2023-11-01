package RomanDecoder

func Decode(roman string) int {
	return parseRoman(roman, 0)
}

func parseRoman(roman string, acc int) int{
	if len(roman) > 0 {
		switch roman[0]{
		case 'I':
			if len(roman) >1  {
				switch roman[1]{
				case 'V':
					return acc + 4
				case  'X': 
					return acc + 9
				default:
					return parseRoman(roman[1:], acc + 1)
				}
			}else{
				return acc + 1
			}
		case 'V':
			return parseRoman(roman[1:], acc + 5)
		case 'X':
			if len(roman)>1 {
				switch roman[1]{
				case 'L':
					return parseRoman(roman[2:], acc + 40)
				case 'C':
					return parseRoman(roman[2:], acc + 90)
				default:
					return parseRoman(roman[1:], acc + 10)
				}
				
			}else{
				return acc +10
			}
		case 'L':
			return parseRoman(roman[1:], acc + 50)
		case 'C':
			if len(roman)>1 {
				switch roman[1]{
				case 'D':
					return parseRoman(roman[2:], acc + 400)
				case 'M':
					return parseRoman(roman[2:], acc + 900)
				default:
					return parseRoman(roman[1:], acc + 100)
				}
				
			}else{
				return acc +100
			}
		case 'D':
			return parseRoman(roman[1:], acc + 500)
		case 'M':
			return parseRoman(roman[1:], acc + 1000)
		}
	}
	return acc + 0
}
