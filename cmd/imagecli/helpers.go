// cmd/imagecli/helpers.go
package imagecli

import (
	"fmt"
	"strconv"
	"strings"
)

func parseSize(sizeStr string) (int, int, error) {
	parts := strings.Split(sizeStr, "x")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат размера")
	}
	width, err1 := strconv.Atoi(parts[0])
	height, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("неверные значения ширины или высоты")
	}
	return width, height, nil
}
