
package main

import s "strings"

import (
    "fmt"
    "os"
    "bufio"
    "strconv"    
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  arifmeticOperators := [4]string {"+","-","*","/"}  
   
  fmt.Println("Введите выражение")   
  expression, _ := reader.ReadString('\n') 
  expression = s.TrimSpace(expression)  
            
  arithmeticOperator := FindArithmeticOperator(expression, arifmeticOperators)
  resultArab, resultRome := SpliteToNumbers(expression, arithmeticOperator)
  if resultArab != 0 {
    fmt.Println(resultArab)
  } else {
    fmt.Println(resultRome)     
  }         
}

//Выявление арифмитического операнда из выражения
func FindArithmeticOperator(expression string, arifmetic [4]string) string{   
    
  var arithmeticOperator string 
  count := 0
  sumCount :=0
             
  for _, ch := range arifmetic { 
    count = s.Count(expression, ch)      
      
    if count > 0 {
      sumCount += count   
      arithmeticOperator = ch
    }               
   }
   
  if sumCount == 1{
    return arithmeticOperator
  } else {
    panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
  }    
}

//Разделение выражения на операнды и его вычесление
func SpliteToNumbers(expression string, arithmeticOperator string )(int, string) {
      
  numbers := s.Split(expression, arithmeticOperator)
 
  numberOneArab, _ := strconv.Atoi(numbers[0])
  numberTwoArab, _ := strconv.Atoi(numbers[1])
  
  if numberOneArab > 0 && numberTwoArab > 0 {
    if (numberOneArab > 10 || numberTwoArab > 10 ){
      panic("Некорректный ввод чисел")     
    } 
    sumArab := calculate (numberOneArab, numberTwoArab, arithmeticOperator)
    return sumArab, ""
     
  } else {
    romeToArab := MapRomeToArab()
    numberOneRome, _ := romeToArab[numbers[0]]
    numberTwoRome, _ := romeToArab[numbers[1]]
        
    if numberOneRome > 10 || numberTwoRome > 10 || numberOneRome == 0 || numberTwoRome == 0 {
      panic("Некорректный ввод чисел")     
    }    
    result := calculate (numberOneRome, numberTwoRome, arithmeticOperator)       
    if result < 1 {
      panic("Выдача паники, так как в римской системе нет отрицательных чисел. ")     
    }
    sumRome := ConverArabToRome(result) 
    return 0, sumRome
  }

}

//Вычесление выражения согласно оператору
func calculate(numberOne int, numberTwo int, arithmeticOperator string ) int{
  
  result := 0
  switch arithmeticOperator {
  case "+":
    result = numberOne + numberTwo
  case "-":
    result = numberOne - numberTwo
  case "*":
    result = numberOne * numberTwo      
  case "/":
    result = numberOne / numberTwo  
  default:
    panic("Некорректный оператор ввода")  
  }
  return result 
}

//Конвертация суммы выражения в римские числа
func ConverArabToRome(sumArab int) string{
  result := ""
  arabToRome := MapArabToRome()
  hundred := sumArab / 100
  hundredRemainder := sumArab % 100 
  if hundredRemainder >= 90{
    result += arabToRome[90] 
    hundredRemainder -= 90 
  }
 
  fifty := hundredRemainder / 50
  result += arabToRome[fifty * 50]
  fiftyRemainder := fifty % 50
  if hundredRemainder >= 40{
    result += arabToRome[40] 
    hundredRemainder -= 40 
  }
  
  decade := fiftyRemainder / 10
  remainder := sumArab % 10
  result += arabToRome[hundred * 100]
  result += arabToRome[fifty * 50]
  for decade > 0 {
    result += arabToRome[10]
    decade--    
  }
  
  result += arabToRome[remainder]
  return result
}

//Объявление карты римские цифры в арабские
func MapRomeToArab() map[string]int{
  romeToArab := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10} 
  return romeToArab
}
//Объявление карты арабские цифры в римские
func MapArabToRome() map[int]string{
  arabToRome := map[int]string{ 1:"I", 2:"II", 3:"III", 4:"IV", 5:"V",6:"VI", 7:"VII", 8:"VIII", 9:"IX", 10:"X",40:"XL", 50:"L",90:"XC", 100:"C"} 
  return arabToRome
}

