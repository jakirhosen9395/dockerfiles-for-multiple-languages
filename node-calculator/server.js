const express = require('express');
const path = require('path');
const app = express();

// View engine setup
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));

// Middleware
app.use(express.static(path.join(__dirname, 'public')));
app.use(express.urlencoded({ extended: true }));

// GET form
app.get('/', (req, res) => {
  res.render('index', { result: "", error: "", num1: "", num2: "", operator: "" });
});

// POST form
app.post('/', (req, res) => {
  const { num1, num2, operator } = req.body;
  let result = "";
  let error = "";

  if (!num1 || !num2 || !operator) {
    error = "Please fill in all fields.";
  } else {
    const n1 = parseFloat(num1);
    const n2 = parseFloat(num2);

    if (isNaN(n1) || isNaN(n2)) {
      error = "Invalid numbers.";
    } else {
      switch (operator) {
        case 'add': result = n1 + n2; break;
        case 'subtract': result = n1 - n2; break;
        case 'multiply': result = n1 * n2; break;
        case 'divide':
          if (n2 === 0) error = "Cannot divide by zero";
          else result = n1 / n2;
          break;
        default: error = "Invalid operator.";
      }
    }
  }

  res.render('index', { result, error, num1, num2, operator });
});

// Start server
const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}/`);
});
