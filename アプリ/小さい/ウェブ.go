package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"github.com/gojp/kana" //go get github.com/gojp/kana
)

var port string = ":8080"

func main() {

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	y := time.Now()
	z := y.Hour()

	if z < 19 && z > 6 {
		fmt.Fprintf(w, "こにちわ %s さん", r.URL.Path)

	} else {
		fmt.Fprintf(w, "こんばんわ %s さん", r.URL.Path)
	}

	a := "人一丨口日目儿見凵山出十八木未丶来大亅了子心土冂田思二丁彳行寸寺時卜上丿刀分厶禾私中彐尹事可亻何自乂又皮彼亠方生月門間扌手言女本乙气気干年三耂者刂前勹勿豕冖宀家今下白勺的云牛物立小文矢知入乍作聿書学合"

	s := strings.Split(a, "")

	fmt.Println("せーべーのポーとト番号は" + port + "です")
	k := kana.RomajiToHiragana("kanji")
	fmt.Fprintf(w, "勉強する漢字("+ k +  ")は:                        ")

	for i := 1; i <= 1; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r := r1.Intn(99)
		fmt.Fprintf(w, s[r])
		time.Sleep(5 * time.Second)

	}

	
}
