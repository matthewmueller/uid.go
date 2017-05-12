package uid

import (
	"strconv"

	"github.com/spaolacci/murmur3"
)

// Hash incoming to a simple unique ID
func Hash(incoming string) string {
	h64 := murmur3.New64()
	h64.Write([]byte(incoming))
	result := h64.Sum64()
	return strconv.FormatUint(result, 36)
}
