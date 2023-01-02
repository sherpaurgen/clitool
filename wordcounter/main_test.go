package main
import(
	"bytes"
	"testing"
)
func TestCountWords(t *testing.T){
	b:=bytes.NewBufferString("hi there are five words\n")
	exp:=5
	res:=count(b)
	if res!=exp{
		t.Errorf("expected %d but got %d",exp,res)
	}
}
