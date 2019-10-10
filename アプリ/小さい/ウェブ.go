package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	a := "人一丨口日目儿見凵山出十八木未丶来大亅了子心土冂田思二丁彳行寸寺時卜上丿刀分厶禾私中彐尹事可亻何自乂又皮彼亠方生月門間扌手言女本乙气気干年三耂者刂前勹勿豕冖宀家今下白勺的云牛物立小文矢知入乍作聿書学合"

	s := strings.Split(a, "")

	fmt.Println("Kanjis Frecuentes")

	fmt.Println("Random Frecuentes")

	time.Sleep(5 * time.Second)

	for i := 0; i <= 10; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r := r1.Intn(89)
		fmt.Fprintf(w, s[r])
		time.Sleep(1 * time.Second)
	}

}
