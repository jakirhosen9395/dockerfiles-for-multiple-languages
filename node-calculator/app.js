const express = require('express');
const app = express();

// Set the view engine to EJS
app.set('view engine', 'ejs');

// Serve static files (e.g., CSS) from the "public" directory
app.use(express.static('public'));

// Parse URL-encoded bodies (form submissions)
app.use(express.urlencoded({ extended: true }));

// Handle GET requests (render form)
app.get('/', (req, res) => {
    res.render('index', {
        result: "",
        error: "",
        num1: "",
        num2: "",
        operator: ""
    });
});

// Handle POST requests (process form)
app.post('/', (req, res) => {
    const { num1, num2, operator } = req.body;
    let result = "";
    let error = "";

    // Validate inputs and perform calculations
    if (!num1 || !num2 || !operator) {
        error = "Please fill in all fields.";
    } else {
        try {
            let n1 = parseFloat(num1);
            let n2 = parseFloat(num2);

            switch (operator) {
                case 'add':
                    result = n1 + n2;
                    break;
                case 'subtract':
                    result = n1 - n2;
                    break;
                case 'multiply':
                    result = n1 * n2;
                    break;
                case 'divide':
                    if (n2 === 0) {
                        error = "Cannot divide by zero";
                    } else {
                        result = n1 / n2;
                    }
                    break;
                default:
                    error = "Invalid operator.";
            }
        } catch (e) {
            error = "Invalid input. Please enter valid numbers.";
        }
    }

    res.render('index', {
        result,
        error,
        num1,
        num2,
        operator
    });
});

// Start the server
app.listen(8080, () => {
    console.log("Server running at http://localhost:8080/");
});
