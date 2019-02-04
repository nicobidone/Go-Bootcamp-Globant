package thetest

import "testing"
import "github.com/nicob/db/thefileops"

func TestFileOps(t *testing.T) {

	thefileops.CreateFile()
	thefileops.WriteFile()
	thefileops.ReadFile()
	thefileops.DeleteFile()
}
