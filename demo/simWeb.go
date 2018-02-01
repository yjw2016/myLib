package main
import(
	"net/http"
	"fmt"
	"log"
	"strings"
	"time"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()   //解析参数，默认是不会解析的
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("r.Form:url_long:",r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("Key:", k)
		fmt.Println("Val:", strings.Join(v, ""))
	}
	map1 := map[int]string{1:"a",2:"b",3:"c"}
	fmt.Fprintf(w, "Hello Cofox!" + time.Now().String(),map1)
	//...此处返回类似java中@ResponseBody？？

}
func main() {
	p := &Person{"joel", 45}
	http.HandleFunc("/"+p.Name, p.sayCN)
	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}


type Person struct {
	Name    string
	Age     int
}

func (ps *Person) sayCN(rw http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(rw, "你好啊,%v,你%d岁了吗？\n", ps.Name, ps.Age)
}