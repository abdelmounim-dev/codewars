package kata

func Solution(str string) []string {
  if len(str) % 2 == 1 {
    str += "_"
  }
  strArr := []string {}
  for i := 0; i < len(str); i += 2 {
    strArr = append(strArr, str[i:i+2])
  }

  return strArr
}
