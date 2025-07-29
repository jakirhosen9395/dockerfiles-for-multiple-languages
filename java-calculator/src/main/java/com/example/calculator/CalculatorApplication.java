package com.example.calculator;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

@SpringBootApplication
@Controller
public class CalculatorApplication {

    public static void main(String[] args) {
        SpringApplication.run(CalculatorApplication.class, args);
    }

    @RequestMapping("/calculator")
    public String calculator(Model model) {
        model.addAttribute("num1", "");
        model.addAttribute("num2", "");
        model.addAttribute("operator", "");
        model.addAttribute("result", "");
        model.addAttribute("error", "");
        return "index";
    }

    @PostMapping("/calculator")
    public String calculate(@RequestParam("num1") String num1,
                            @RequestParam("num2") String num2,
                            @RequestParam("operator") String operator,
                            Model model) {
        double calcResult = 0;
        String error = "";

        try {
            if (num1.isEmpty() || num2.isEmpty() || operator.isEmpty()) {
                error = "Missing parameters";
            } else {
                double num1Float = Double.parseDouble(num1);
                double num2Float = Double.parseDouble(num2);

                switch (operator) {
                    case "add":
                        calcResult = num1Float + num2Float;
                        break;
                    case "subtract":
                        calcResult = num1Float - num2Float;
                        break;
                    case "multiply":
                        calcResult = num1Float * num2Float;
                        break;
                    case "divide":
                        if (num2Float == 0) {
                            error = "Cannot divide by zero";
                        } else {
                            calcResult = num1Float / num2Float;
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
