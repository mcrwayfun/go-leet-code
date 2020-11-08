package main

func reverseString(s []byte)  {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	for i:=0;i<len(s)/2;i++{
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	return
}
