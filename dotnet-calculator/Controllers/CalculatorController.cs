using Microsoft.AspNetCore.Mvc;
using dotnet_calculator.Models;

namespace dotnet_calculator.Controllers
{
    public class CalculatorController : Controller
    {
        // GET: /Calculator/
        [HttpGet]
        public IActionResult Index()
        {
            return View(new CalculatorResult());
        }

        // POST: /Calculator/
        [HttpPost]
        public IActionResult Index(CalculatorResult result)
        {
            if (string.IsNullOrEmpty(result.Num1) || string.IsNullOrEmpty(result.Num2) || string.IsNullOrEmpty(result.Operator))
            {
                result.Error = "Missing parameters";
            }
            else
            {
                double num1, num2;
                if (double.TryParse(result.Num1, out num1) && double.TryParse(result.Num2, out num2))
                {
                    switch (result.Operator.ToLower())
                    {
                        case "add":
                            result.Result = $"Result: {num1 + num2}";
                            break;
                        case "subtract":
                            result.Result = $"Result: {num1 - num2}";
                            break;
                        case "multiply":
                            result.Result = $"Result: {num1 * num2}";
                            break;
                        case "divide":
                            if (num2 == 0)
                            {
                                result.Error = "Cannot divide by zero";
                            }
                            else
                            {
                                result.Result = $"Result: {num1 / num2}";
                            }
                            break;
                        default:
                            result.Error = "Invalid operator";
                            break;
                    }
                }
                else
                {
                    result.Error = "Invalid numbers";
                }
            }

            return View(result);
        }
    }
}
