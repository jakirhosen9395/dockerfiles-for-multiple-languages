package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type CalculatorResult struct {
	Num1     string
	Num2     string
	Operator string
	Result   string
	Error    string
}

var (
	tpl *template.Template
)

// calculatorHandler handles GET (render form) and POST (compute) for /calculator
func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	// Always render the page with whatever data we have
	data := CalculatorResult{}

	switch r.Method {
	case http.MethodGet:
		// nothing to do; render empty form
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			data.Error = "Failed to parse form"
			break
		}

		data.Num1 = r.FormValue("num1")
		data.Num2 = r.FormValue("num2")
		data.Operator = r.FormValue("operator")

		if data.Num1 == "" || data.Num2 == "" || data.Operator == "" {
			data.Error = "Missing parameters"
			break
		}

		n1, err1 := strconv.ParseFloat(data.Num1, 64)
		n2, err2 := strconv.ParseFloat(data.Num2, 64)
		if err1 != nil || err2 != nil {
			data.Error = "Invalid numbers"
			break
		}

		switch data.Operator {
		case "add":
			data.Result = strconv.FormatFloat(n1+n2, 'f', -1, 64)
		case "subtract":
			data.Result = strconv.FormatFloat(n1-n2, 'f', -1, 64)
		case "multiply":
			data.Result = strconv.FormatFloat(n1*n2, 'f', -1, 64)
		case "divide":
			if n2 == 0 {
				data.Error = "Cannot divide by zero"
			} else {
				data.Result = strconv.FormatFloat(n1/n2, 'f', -1, 64)
			}
		default:
			data.Error = "Invalid operator (use: add, subtract, multiply, divide)"
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		data.Error = "Method not allowed"
	}

	if err := tpl.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Printf("template execute error: %v", err)
	}
}

func main() {
	// load templates once at startup
	var err error
	tpl, err = template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	// port from env, default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", calculatorHandler)
	mux.HandleFunc("/calculator", calculatorHandler)

	log.Printf("Starting server on :%s ...", port)
	if err := http.ListenAndServe(":"+port, securityHeaders(mux)); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// securityHeaders is a tiny middleware to add basic headers.
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		next.ServeHTTP(w, r)
	})
}
