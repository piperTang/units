package units
import "fmt"
type Pound float64  //英镑
type Kg float64    //公斤
func (p Pound) String() string {return fmt.Sprintf("%g英镑",p)}
func (k Kg) String() string {return fmt.Sprintf("%g公斤",k)}