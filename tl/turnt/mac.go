package turnt

import (
	"fmt"
)

//Mac2Str is
func Mac2Str(mac []byte) string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

//Str2mac is
func Str2mac(str string) []byte {
	mac := make([]byte, 6)
	fmt.Sscanf(str, "%02x:%02x:%02x:%02x:%02x:%02x", &mac[0], &mac[1], &mac[2], &mac[3], &mac[4], &mac[5])
	return mac
}
