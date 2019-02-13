package thetest

import (
	"strconv"
	"testing"
	"time"

	"github.com/nicob/db/thefileops"
)

func TestSafeCocurrency(t *testing.T) {
	var path = "C:\\Users\\nicob\\go\\src\\github.com\\nicob\\db\\test.txt"
	x := thefileops.NewSafeFile(path)
	x.SafeDeleteFile()
	x.SafeCreateFile()
	for i := 0; i < 1001; i++ {
		go x.SafeWriteFile(strconv.Itoa(i) + "\n")
	}
	time.Sleep(time.Second)
	x.SafeReadFile()
}
