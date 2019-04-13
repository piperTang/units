# units包使用方法
导入包：

`github.com/units`

新建一个测试项目项目：

unitsTest/main.go

```go
package main
import (
	"fmt"
	"os"
	"strconv"
	"ch2/units"
)
func main(){
	for _,arg:=range os.Args[1:]{
		t,err:=strconv.ParseFloat(arg,64)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"cf:%v\n",err)
			os.Exit(1)
		}
		p:=units.Pound(t)
		k:=units.Kg(t)
		fmt.Printf("%s=%s,%s=%s\n",p,units.PToK(p),k,units.KToP(k))
	}
}
```

使用包：

```go
go build main.go
//输入一个数字进行转换
./main 1
//得到以下结果
1英镑=0.45公斤,1公斤=2.2222222222222223英镑
```

