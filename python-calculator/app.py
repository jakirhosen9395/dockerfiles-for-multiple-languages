from flask import Flask, render_template, request

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def calculator():
    result = ""
    error = ""
    
    if request.method == 'POST':
        try:
            num1 = float(request.form['num1'])
            num2 = float(request.form['num2'])
            operator = request.form['operator']

            if operator == 'add':
                result = num1 + num2
            elif operator == 'subtract':
                result = num1 - num2
            elif operator == 'multiply':
                result = num1 * num2
            elif operator == 'divide':
                if num2 == 0:
                    error = "Cannot divide by zero"
                else:
                    result = num1 / num2
            else:
                error = "Invalid operator"
        except ValueError:
            error = "Invalid input. Please enter valid numbers."

    return render_template('index.html', result=result, error=error)

if __name__ == "__main__":
    # Run the Flask app on port 8080
    app.run(debug=True, host='0.0.0.0', port=8080)
