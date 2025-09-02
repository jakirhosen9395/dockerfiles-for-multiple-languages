package com.example.calculator;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;

@Controller
public class CalculatorController {

    @GetMapping("/calculator")
    public String calculator(Model model) {
        model.addAttribute("num1", "");
        model.addAttribute("num2", "");
        model.addAttribute("operator", "");
        model.addAttribute("result", "");
        model.addAttribute("error", "");
        return "index";
    }

    @PostMapping("/calculator")
    public String calculate(
            @RequestParam("num1") String num1,
            @RequestParam("num2") String num2,
            @RequestParam("operator") String operator,
            Model model) {

        double calcResult = 0;
        String error = "";

        try {
            if (num1.isEmpty() || num2.isEmpty() || operator.isEmpty()) {
                error = "Missing parameters";
            } else {
                double n1 = Double.parseDouble(num1);
                double n2 = Double.parseDouble(num2);

                switch (operator) {
                    case "add":
                        calcResult = n1 + n2;
                        break;
                    case "subtract":
                        calcResult = n1 - n2;
                        break;
                    case "multiply":
                        calcResult = n1 * n2;
                        break;
                    case "divide":
                        if (n2 == 0) {
                            error = "Cannot divide by zero";
                        } else {
                            calcResult = n1 / n2;
                        }
                        break;
                    default:
                        error = "Invalid operator";
                }
            }
        } catch (NumberFormatException e) {
            error = "Invalid numbers";
        }

        model.addAttribute("num1", num1);
        model.addAttribute("num2", num2);
        model.addAttribute("operator", operator);
        if (error.isEmpty()) {
            model.addAttribute("result", "Result: " + calcResult);
        } else {
            model.addAttribute("error", error);
        }

        return "index";
    }
}
