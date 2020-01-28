package main

import "fmt"

func breakPalindrome(p string) string {
  n := len(p)
  if n <= 1 {
    return ""
  }
  
  chars := []rune(p)
  l, r := 0, n - 1
  for r > l {
    if chars[l] > 'a' {
      return string(chars[:l]) + "a" + string(chars[l + 1:])
    }

    l++
    r--
  }
  
  chars[n-1] = 'b'
  return string(chars)
}


func main(){
  fmt.Println(breakPalindrome("abccba"))
  fmt.Println(breakPalindrome("acca"))
  fmt.Println(breakPalindrome(""))
  fmt.Println(breakPalindrome("aaa"))
}
