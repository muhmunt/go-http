package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ini index")
}

func parseUrl(w http.ResponseWriter, r *http.Request) {
	var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"

	// fmt.Println(r.URL.Query()["url"])
	// fmt.Println(r.URL.Query()["age"])
	// return
	var u, e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println("url: %s\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s \n", name, age)
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		var data = map[string]string{
// 			"Name":    "Tony Stark",
// 			"Message": "How are you",
// 		}

// 		var t, err = template.ParseFiles("template.html")
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		t.Execute(w, data)
// 	})

// 	http.HandleFunc("/index", index)
// 	http.HandleFunc("/parse", parseUrl)

// 	fmt.Println("starting server at http://localhost:8080/")
// 	http.ListenAndServe(":8080", nil)
// }

// bermain dengan json dan object
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name" : "Johny Will", "Age": 38}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User: ", data.FullName)
	fmt.Println("Age: ", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("User: ", data1["Name"])
	fmt.Println("Age: ", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Print(decodedData)
	fmt.Println("User: ", decodedData["Name"])
	fmt.Println("Age: ", decodedData["Age"])

	// array json to array object
	var jsonString2 = `[
		{"Name" : "Kally", "Age":47},
		{"Name" : "Irving", "Age":27}
	]`

	var dataArray []User

	var err2 = json.Unmarshal([]byte(jsonString2), &dataArray)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	fmt.Println("user 1 : ", dataArray[0].FullName)
	fmt.Println("user 2 : ", dataArray[1].FullName)

	// encode object to json string
	var object = []User{{"Larry Thing", 39}, {"Yurrin", 34}}
	var jsonData2, err3 = json.Marshal(object)

	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	var jsonString3 = string(jsonData2)
	fmt.Println(jsonString3)

}
