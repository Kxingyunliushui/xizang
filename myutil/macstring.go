package mayutil

import (
	"strings"
)

func Utils_mac(v string) string {
	var final_mac string = ""
	maclen := len(v)
	if len(v) >= 12 {
		if 12 == maclen {
			//logger.Error(" <%s> -> <%s>", k, v)
			final_mac = v
		} else {
			t_array := strings.Split(v, ":")
			//logger.Info("*********** after mac111:", m[k][0], len(t_array))
			//var final_mac string
			if 6 <= len(t_array) {
				for i := 0; i < 6; i++ {
					final_mac = final_mac + t_array[i]
				}
			} else {
				t_array := strings.Split(v, "-")
				//logger.Info("----------------", m[k][0], len(t_array))
				if 6 <= len(t_array) {
					for i := 0; i < 6; i++ {
						final_mac = final_mac + t_array[i]
					}
				}
			}
		}
	}

	if len(final_mac) == 12 {
		return strings.ToUpper(final_mac)
	}

	return ""
}
