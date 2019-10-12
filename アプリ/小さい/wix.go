package main

import (
	"fmt"
	"github.com/gorilla/mux" //go get -u github.com/gorilla/mux
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var port string = ":8107"

func YourHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Gorilla!\n"))
	y := time.Now()
	z := y.Hour()

	if z < 19 && z > 6 {
		fmt.Fprintf(w, "こにちわ\n", r.URL.Path)

	} else {
		fmt.Fprintf(w, "こんばんわ\n")
	}

	a := "人一丨口日目儿見凵山出十八木未丶来大亅了子心土冂田思二丁彳行寸寺時卜上丿刀分厶禾私中彐尹事可亻何自乂又皮彼亠方生月門間扌手言女本乙气気干年三耂者刂前勹勿豕冖宀家今下白勺的云牛物立小文矢知入乍作聿書学合"

	s := strings.Split(a, "")

	fmt.Println("せーべーのポーとト番号は" + port + "です")

	fmt.Fprintf(w, "勉強する漢字は:\n ")

	for i := 1; i <= 1; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r := r1.Intn(99)
		fmt.Fprintf(w, s[r])
		time.Sleep(5 * time.Second)

	}

}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(port, r))
}
