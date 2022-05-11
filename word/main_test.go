package word

import (
	"fmt"
	"testing"

	"github.com/josephphyo/words/quote"
)

func TestUseCount(t *testing.T) {
	x := UseCount("phyo dther waiyan phyo phyo dther phyo phyo phyo phyo dther")
	for k, v := range x {
		switch k {
		case "phyo":
			if v != 7 {
				t.Error("Got", v, "expected", 7)
			}
		case "waiyan":
			if v != 1 {
				t.Error("Got", v, "expected", 1)
			}
		case "dther":
			if v != 3 {
				t.Error("Got", v, "Expected", 3)
			}
		}
	}
}

func TestCount(t *testing.T) {
	x := Count("phyo dther waiyan phyo phyo dther phyo phyo phyo phyo dther")
	if x != 11 {
		t.Error("Got", x, "Expected", 11)
	}
}

func ExampleCount() {
	fmt.Println(Count("phyo dther waiyan phyo phyo dther phyo phyo phyo phyo dther"))
	// output :
	// 11
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("phyo dther waiyan phyo phyo dther phyo phyo phyo phyo dther")
	}
}
