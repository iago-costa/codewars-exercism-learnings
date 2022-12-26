package reverse_words

import "strings"
import "fmt"

func ReverseWords(str string) (result string) {
  list_of_words := strings.Split(str, " ")
  func reverse(word string) (result string) {
    for _, caracteres := range str {
      result = string(caracteres) + result
    }
    return
  }
  for _, word := range list_of_words {
    result += reverse(word) + " "
  }
}

