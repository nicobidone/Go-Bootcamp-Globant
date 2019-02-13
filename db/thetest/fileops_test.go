package thetest

import "testing"
import "github.com/nicob/db/thefileops"

func TestFileOps(t *testing.T) {
	var path = "C:\\Users\\nicob\\go\\src\\github.com\\nicob\\db\\test.txt"
	thefileops.CreateFile(path)
	thefileops.WriteFile(path, "Hello \nWorld\n")
	thefileops.ReadFile(path)
	thefileops.DeleteFile(path)
}
