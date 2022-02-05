package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

//var tpl *template.Template

//DB
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

//start date and end date
type Dates struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}

// export fields to templates
// fields changed to uppercase
type OrderInfo struct {
	Order_name         string `json:"order_name"`
	Product            string `json:"product"`
	Company_name       string `json:"company_name"`
	Name               string `json:"name"`
	Created_at         string `json:"created_at"`
	Delivered_quantity string `json:"delivered_quantity"`
}

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	//tpl = template.Must(template.ParseGlob("C:/Users/Usman/Desktop/Full-Stack Engineer Test/Web Application/Frontend/*.gohtml"))
	http.HandleFunc("/", orders)
	http.HandleFunc("/orders", orders)
	http.HandleFunc("/ordersSearch", ordersSearch)
	http.ListenAndServe(":8090", nil)
}

func orders(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	rows, err := db.Query("Select public.orders.order_name,public.order_items.product,public.customer_companies.company_name,public.customers.name,public.orders.created_at,public.deliveries.delivered_quantity\n" +
		"FROM public.customer_companies\n" +
		"INNER JOIN public.customers\n" +
		"ON public.customer_companies.company_Id=public.customers.company_Id\n" +
		"INNER JOIN public.orders\n" +
		"ON public.customers.user_Id=public.orders.customer_id\n" +
		"INNER JOIN public.order_items\n" +
		"ON public.orders.id=public.order_items.order_id\n" +
		"INNER JOIN public.deliveries\n" +
		"ON public.order_items.id=public.deliveries.order_item_id\n")

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	//decoder.Decode(&rows)
	defer rows.Close()
	bks := make([]OrderInfo, 0)
	for rows.Next() {
		bk := OrderInfo{}
		err := rows.Scan(&bk.Order_name, &bk.Product, &bk.Company_name, &bk.Name, &bk.Created_at, &bk.Delivered_quantity) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	fmt.Print("Data Retrieved\n")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bks); err != nil {
		panic(err)
	}
	//tpl.ExecuteTemplate(w, "orders.gohtml", bks)
}
func ordersSearch(w http.ResponseWriter, r *http.Request) {

	fmt.Print("OrderSearch is checked\n")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	var DateSearch Dates
	decoder.Decode(&DateSearch)
	fmt.Print("StartDate is: " + DateSearch.StartDate + "\n")
	fmt.Print("EndDate is:" + DateSearch.EndDate + "\n")
	rows, err := db.Query("Select public.orders.order_name,public.order_items.product,public.customer_companies.company_name,public.customers.name,public.orders.created_at,public.deliveries.delivered_quantity\n" +
		"FROM public.customer_companies\n" +
		"INNER JOIN public.customers\n" +
		"ON public.customer_companies.company_Id=public.customers.company_Id\n" +
		"INNER JOIN public.orders\n" +
		"ON public.customers.user_Id=public.orders.customer_id\n" +
		"INNER JOIN public.order_items\n" +
		"ON public.orders.id=public.order_items.order_id\n" +
		"INNER JOIN public.deliveries\n" +
		"ON public.order_items.id=public.deliveries.order_item_id\n" +
		"WHERE public.orders.created_at between '" + DateSearch.StartDate + "' and '" + DateSearch.EndDate + "' ")

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	//decoder.Decode(&rows)
	defer rows.Close()
	bks := make([]OrderInfo, 0)
	for rows.Next() {
		bk := OrderInfo{}
		err := rows.Scan(&bk.Order_name, &bk.Product, &bk.Company_name, &bk.Name, &bk.Created_at, &bk.Delivered_quantity) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}

	fmt.Print(bks)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bks); err != nil {
		panic(err)
	}
	//tpl.ExecuteTemplate(w, "orders.gohtml", bks)
}
