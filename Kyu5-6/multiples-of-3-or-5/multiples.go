package kata

func Multiple3And5(number int) int {
  result := 0
  
  for i := 1; i < number; i++ {
    if i % 3 == 0 && i % 5 == 0 {
      result += i
      continue
    } else if i % 3 == 0 || i % 5 == 0 {
      result += i
    }
  }
  
  return result
}
