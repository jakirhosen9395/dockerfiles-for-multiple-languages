package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a struct to store the result and input values for rendering
type CalculatorResult struct {
	Num1     string
	Num2     string
	Operator string
	Result   string
	Error    string
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct to store the input values and result for rendering
	var result CalculatorResult

	// Handle POST requests
	if r.Method == "POST" {
		// Parse form data
		r.ParseForm()
		num1 := r.FormValue("num1")
		num2 := r.FormValue("num2")
		operator := r.FormValue("operator")

		// Store the inputs in the struct
		result.Num1 = num1
		result.Num2 = num2
		result.Operator = operator

		var calcResult float64
		var err string

		// Validate input values
		if num1 == "" || num2 == "" || operator == "" {
			err = "Missing parameters"
		} else {
			// Convert string inputs to float64
			num1Float, err1 := strconv.ParseFloat(num1, 64)
			num2Float, err2 := strconv.ParseFloat(num2, 64)

			if err1 != nil || err2 != nil {
				err = "Invalid numbers"
			} else {
				// Perform the calculation based on the operator
				switch operator {
				case "add":
					calcResult = num1Float + num2Float
				case "subtract":
					calcResult = num1Float - num2Float
				case "multiply":
					calcResult = num1Float * num2Float
				case "divide":
					if num2Float == 0 {
						err = "Cannot divide by zero"
					} else {
						calcResult = num1Float / num2Float
					}
				default:
					err = "Invalid operator"
				}
			}
		}

		// Set the result or error message in the struct
		if err != "" {
			result.Error = err
		} else {
			result.Result = fmt.Sprintf("Result: %f", calcResult)
		}
	}

	// Render the HTML template with the struct data
	tmpl, _ := template.New("result").Parse(`
                <html>
                <head><title>Simple Calculator</title></head>
                <body>
                        <h1>Simple Calculator</h1>


                        <form action="/calculator" method="POST">
                                <label for="num1">Number 1:</label>
                                <input type="text" id="num1" name="num1" value="{{.Num1}}"><br><br>


                                <label for="num2">Number 2:</label>
                                <input type="text" id="num2" name="num2" value="{{.Num2}}"><br><br>


                                <label for="operator">Operator (add, subtract, multiply, divide):</label>
                                <input type="text" id="operator" name="operator" value="{{.Operator}}"><br><br>


                                <input type="submit" value="Calculate">
                                <button type="button" onclick="clearResult()">Clear Result</button>
                        </form>


                        <p id="result">{{.Result}}</p>
                        <p style="color:red;">{{.Error}}</p>


                        <script>
                                function clearResult() {
                                        document.getElementById("result").innerHTML = "";
                                }
                        </script>
                </body>
                </html>
        `)

	tmpl.Execute(w, result)
}

func main() {
	// Log that the server is starting
	fmt.Println("Starting server on :8080...")

	// Register the handler function for the /calculator endpoint
	http.HandleFunc("/calculator", calculatorHandler)

	// Start the web server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
