package printer

import (
	"fmt"

	"github.com/google/uuid"
)

func PrinterNewUUID() string {
	id := uuid.New()

	return fmt.Sprintf("Generated uuid: %s\n", id)
}
