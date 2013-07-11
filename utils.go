package main

func orit(a ...string) string {
	var ret string
	for _, v := range a {
		if v != "" {
			ret = v
		}
	}
	return ret
}
