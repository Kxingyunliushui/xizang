package myutil

func Util_time(tm string) string {

	length := len(tm)
	if length == 16 {
		return tm
	} else if length < 16 {
		for i := length; i < 16; i++ {
			tm = tm + "0"
		}
		return tm
	} else if length > 16 {
		return tm[:15]
	}
	return ""
}
