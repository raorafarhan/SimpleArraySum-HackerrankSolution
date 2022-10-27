func simpleArraySum(ar []int32) int32 {
	var temp int32
	for _, v := range ar {
		temp = temp + v
	}
	return temp
}