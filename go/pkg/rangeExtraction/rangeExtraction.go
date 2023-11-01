package rangeExtraction

import (
	"strconv"
)

func appendRange(prev string, start int, end int) string {
	if prev != "" {
		prev += ","
	}

	if (start==end){
		prev +=  strconv.Itoa(start)
	}else if (start +1 == end) {
		prev +=  strconv.Itoa(start) + "," + strconv.Itoa(end)
	} else {
		prev +=  strconv.Itoa(start) + "-" + strconv.Itoa(end)
	}
	return prev
}

func Solution(list []int) string {

	if(len(list) > 0) {
		start := list[0]
		end := list[0]
		
		retVal:=""
		for _, next := range list[1:] {
			if next > end+1 {
				retVal = appendRange(retVal, start, end)			
				start = next
				end = next
			}
			end = next
		}
		retVal = appendRange(retVal, start , end )
		return retVal
	}else{
		return ""
	}
}
