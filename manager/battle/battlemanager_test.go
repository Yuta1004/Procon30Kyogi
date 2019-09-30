package battle

import (
	"testing"
	"time"
)

func TestBManagerExec(t *testing.T) {
	ticker := time.NewTicker(10 * time.Second)
	go BManagerExec("A")
	<-ticker.C
}
