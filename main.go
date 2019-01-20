package main
import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"math/rand"
	"strconv"
	"fmt"
)
type Capacity struct {
	Type string   `json:"type"`
	Icon string   `json:"icon"`
	Title  string   `json:"title"`
	Value  string   `json:"value"`
	FooterText  string   `json:"footerText"`
	FooterIcon  string   `json:"footerIcon"`
}
 type User struct {
 	Email  string `json:"email"`
 	Password string `json:"password"`
 }
 type Userchart struct{
	Time string   `json:"time"`
	ActionType string   `json:"actiontype"`
	Value  int   `json:"value"`

}

type Salesdata struct{
	Model string   `json:"model"`
	Month string   `json:"month"`
	SalesAmount  int   `json:"salesAmount"`
}
 type Emaildata struct{
	Status string   `json:"status"`
	PercentageValue int   `json:"percentageValue"`
}
type Req struct {
	FromDate  string `json:"fromDate"`
	ToDate string `json:"toDate"`
}
//get stat cards
func Getstats(w http.ResponseWriter, r *http.Request) {
	var statscard [] Capacity

	var req Req
	_ = json.NewDecoder(r.Body).Decode(&req)

	fmt.Println(req.FromDate)
	fmt.Println(req.ToDate)
	var capValue = rand.Intn(256)
	var revValue = rand.Intn(10000)
	var errValue = rand.Intn(50)
	var followValue = rand.Intn(100)

	//capacity
	statscard = append(statscard,Capacity{Type: "warning", Icon: "ti-server", Title: "Capacity", Value: strconv.Itoa(capValue) + "GB",FooterText:"Updated now",FooterIcon:"ti-reload"})
	//
	statscard = append(statscard,Capacity{Type: "success", Icon: "ti-wallet", Title: "Revenue", Value: "$"+strconv.Itoa(revValue) ,FooterText:"Last day",FooterIcon:"ti-calendar"})
	//
	statscard = append(statscard,Capacity{Type: "danger", Icon: "ti-pulse", Title: "Errors", Value: strconv.Itoa(errValue) ,FooterText:"In the last hour",FooterIcon:"ti-timer"})
	//
	statscard = append(statscard,Capacity{Type: "info", Icon: "ti-twitter-alt", Title: "Followers", Value: "+"+strconv.Itoa(followValue) ,FooterText:"Updated now",FooterIcon:"ti-reload"})
   w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statscard)
}
//user login function
func Userlogin(w http.ResponseWriter, r *http.Request) {
 	var login User
	_ = json.NewDecoder(r.Body).Decode(&login)

	if (login.Email == "anjali@email.com" && login.Password == "test1234" ){

			fmt.Println("Login Sucess")
			json.NewEncoder(w).Encode(true)

	} else {
			fmt.Println("Login Failed")
			json.NewEncoder(w).Encode(false)
	}
}
//get user chart data
func Getuserchart(w http.ResponseWriter, r *http.Request){
	var userchart [] Userchart
	var req Req
	_ = json.NewDecoder(r.Body).Decode(&req)

	fmt.Println(req.FromDate)
	fmt.Println(req.ToDate)
	var time = []string{"9.00AM",
	"12:00AM",
	"3:00PM",
	"6:00PM",
	"9:00PM",
	"12:00PM",
	"3:00AM",
	"6:00AM"}
	 var openValue = rand.Intn(300)
	var singleclickValue = rand.Intn(150)
	 var doubleclickValue = rand.Intn(200)
	for i := 0; i < len(time); i++ {
		// var actionData [] Data
		// actionData = append(actionData,Data{ActionType:"open",Value:openValue + (i+1)*75})
		// actionData = append(actionData,Data{ActionType:"click",Value:openValue + (i+1)*60})
		// actionData = append(actionData,Data{ActionType:"secondclick",Value:openValue + (i+1)*50})
		// userchart = append(userchart,Userchart{Time:time[i],Data: actionData})
		userchart = append(userchart,Userchart{Time:time[i],ActionType: "open", Value: openValue + i*30})
	}
	for i := 0; i < len(time); i++ {
		userchart = append(userchart,Userchart{Time:time[i],ActionType: "click", Value: singleclickValue + i*30})
	}
	for i := 0; i < len(time); i++ {
		userchart = append(userchart,Userchart{Time:time[i],ActionType: "secondclick", Value: doubleclickValue + i*40})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userchart)
	fmt.Println(userchart)

}
//get sales data
func Getsalesdata(w http.ResponseWriter, r *http.Request){
	var sale [] Salesdata
	var req Req
	_ = json.NewDecoder(r.Body).Decode(&req)

	fmt.Println(req.FromDate)
	fmt.Println(req.ToDate)
	var month = []string{"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec"}


	for i := 0; i < len(month); i++ {
		var openValue = rand.Intn(900)
      sale = append(sale,Salesdata{Model:"Tesla Model S",Month:month[i],SalesAmount:openValue })
	}
	for i := 0; i < len(month); i++ {
		var openValue = rand.Intn(900)
      sale = append(sale,Salesdata{Model:"BMW 5 Series",Month:month[i],SalesAmount:openValue })
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sale)
	fmt.Println(sale)

}
// get email data
func Getemaildata(w http.ResponseWriter, r *http.Request) {
	var email [] Emaildata
	var req Req
	_ = json.NewDecoder(r.Body).Decode(&req)

	fmt.Println(req.FromDate)
	fmt.Println(req.ToDate)
	var openValue = rand.Intn(100)
	var bounceValue = (100-openValue)/2
	var unsubValue = 100-(openValue + bounceValue)


	//capacity
	email = append(email,Emaildata{Status: "Open", PercentageValue: openValue})
	//
	email = append(email,Emaildata{Status: "Bounce", PercentageValue: bounceValue})
	//
	email = append(email,Emaildata{Status: "Unsubscribe", PercentageValue: unsubValue})

   w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(email)
	fmt.Println(email)
}


func main() {

    allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080","http://127.0.0.1:8080","http://192.168.1.155:8080"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	router := mux.NewRouter()
	router.HandleFunc("/statscards", Getstats).Methods("POST")
	router.HandleFunc("/login", Userlogin).Methods("POST")
	router.HandleFunc("/userchart", Getuserchart).Methods("POST")
	router.HandleFunc("/salesdata", Getsalesdata).Methods("POST")
	router.HandleFunc("/emaildata", Getemaildata).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS( allowedOrigins,allowedMethods)(router)))
}
