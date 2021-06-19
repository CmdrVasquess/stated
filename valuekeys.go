package stated

import "strings"

func vkeyStripVarEsc(k string) string {
	if k == "" {
		return k
	}
	if k[0] == '$' {
		k = k[1:]
	}
	if l := len(k); k[l-1] == ';' {
		k = k[:l-1]
	}
	return k
}

func ShipKey(k string) string { return strings.ToLower(k) }

func EconomyKey(k string) string {
	k = strings.ToLower(k)
	k = vkeyStripVarEsc(k)
	k = strings.TrimPrefix(k, "economy_")
	return k
}

func GovernmentKey(k string) string {
	k = strings.ToLower(k)
	k = vkeyStripVarEsc(k)
	k = strings.TrimPrefix(k, "government_")
	return k
}

func SysSecurityKey(k string) string {
	k = strings.ToLower(k)
	k = vkeyStripVarEsc(k)
	k = strings.TrimPrefix(k, "system_security_")
	return k
}

func PortSvcKey(k string) string {
	return strings.ToLower(k)
}
