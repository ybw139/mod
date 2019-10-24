package  mod
import  (
        "fmt"
)
type  Hello struct {
        HH int
}
func (h *Hello) print() {
fmt.Println(h.HH)
}

