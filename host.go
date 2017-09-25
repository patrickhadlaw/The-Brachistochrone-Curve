//
//	Analysing the Brachistochrone curve
//	file: host.go
//	description: hosts brachistochrone curve and stores/sends leaderboard data
//  port: 8080
//
//	by: Patrick Stanislaw Hadlaw
//


package main

import(
	"net/http"
	"io/ioutil"
	"io"
	"strconv"
	"errors"
	"sort"
	"fmt"
)

func main(){
	
	leaderboard := make(map[float64][]string) // Leaderboard data: time and name
	
	mux := http.NewServeMux()
	
	// FileServer in directory com/brachistochrone/
	fs := http.FileServer(http.Dir("com/files"))
	mux.Handle("/files/", http.StripPrefix("/files/", fs))
	
	// Default response: writes The Brachistochrone Curve.html file 
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		f, err := ioutil.ReadFile("com/The Brachistochrone Curve.html");
		if err != nil{
			fmt.Println(err.Error())
			io.WriteString(w, "<h1>404: File Not Found</h1>")
		}else{
			io.WriteString(w, string(f))
		}
	})
	
	// Data url obtains leaderboard values, name is the users name and time is the time taken for ball to reach finish line
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request){
		n, ok := r.URL.Query()["name"]
		if !ok{
			n[0] = "Unnamed"
		}
		if n[0] == "" {
			n[0] = "Unnamed"
		}
		m, ok := r.URL.Query()["time"]
		if !ok{
			fmt.Println("Error: missing time attrib")
		}else{
			i, err := strconv.ParseFloat(m[0], 64)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				_, ok := leaderboard[i]
				if !ok{
					leaderboard[i] = []string{n[0]}
				}else{
					leaderboard[i] = append(leaderboard[i], n[0])
				}
			}
		}
	})
	
	// Leaderboard url gives live top ten leaderboard data in table on request
	mux.HandleFunc("/leaderboard/", func(w http.ResponseWriter, r *http.Request){
		l := ""
		var keys []float64
		for k := range leaderboard {
			keys = append(keys, k)
		}
		sort.Sort(sort.Float64Slice(keys)) // Sorts leaderboard
		
		if len(keys) >= 10{ // Only show top ten
			keys = keys[:10]
		}
		for i := range keys{ // Create table
			x := leaderboard[keys[i]]
			list := "<tr><td><b>" + strconv.Itoa(i+1) + "</b>&#58; " + strconv.FormatFloat(keys[i], 'f', 6, 64) + "s</b></td><td><b>"
			first := true
			for _, j := range x{
				if !first{
					list = list + ", "
				}else{
					first = false
				}
				list = list + j
			}
			l = l + list + "</b></td></tr>"
		}
		io.WriteString(w, l) // Write response
	})
	
	// Favicon.ico gives site icon
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request){
		f, err := ioutil.ReadFile("favicon.ico");
		if err != nil{
			fmt.Println(err.Error())
			io.WriteString(w, "<h1>404: File Not Found</h1>")
		}else{
			io.WriteString(w, string(f))
		}
	})
	
	// Reset allows for remote reset of leaderboard data
	mux.HandleFunc("/reset/", func(w http.ResponseWriter, r *http.Request){
		n, ok := r.URL.Query()["pass"]
		if ok && n[0] == "admin"{
			leaderboard = make(map[float64][]string)
		}
	})
	
	// Exit allows for remote shutdown
	mux.HandleFunc("/exit/", func(w http.ResponseWriter, r *http.Request){
		n, ok := r.URL.Query()["pass"]
		if ok && n[0] == "admin"{
			os.Exit(0)
		}
	})
	
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		fmt.Println(err.Error())
	}
}