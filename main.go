package main

import (
	"fmt"
	"net/http"
	"os"
)
//write the name of site that you want check; example (go run main.go golang)
func main() {

	a:=os.Args[1]
	var res [22]int

	var arr_RUS = [...]string{".рф",

		".com.ru",

		".exnet.su",

		".net.ru",

		".org.ru",

		".pp.ru",

		".ru",

		".ru.net",

		".su",//end of russian domens
		".aero",

		".asia",

		".biz",

		".com",

		".info",

		".mobi",

		".name",

		".net",

		".org",

		".pro",

		".tel",

		".travel",

		".xxx"	/*end of international domens*/}
	i:=0

for ; i<22; i++{
		url := "http://" + a + arr_RUS[i]
		fmt.Println(url)
		site, err := http.Get(url)

		if err != nil {
			res[i] = 1
			fmt.Println(i,")",url)
			fmt.Println( "You can use this host: (", arr_RUS[i],")")
			}
		if	err == nil {
			defer site.Body.Close() //we should to close our site in the end)
			fmt.Println(i, ") ", url)
			fmt.Println(site.Status, " - host:(",  arr_RUS[i],")")
		}
	}

fmt.Println("The list of free host:")
	i = 0
	for ; i<22; i++ {

		if i == 10{
		  fmt.Println("___international domen")
	}
		if i == 0{
		  fmt.Println("___Russian domen")
	}
		if res[i] == 1{
			fmt.Println(arr_RUS[i])
		}

	}

}


