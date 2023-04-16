package validation

import (
	"fmt"
	"regexp"
	"strconv"
)

// validateINN cheks lehght and checksum of legal persons inn number
func ValidateINN(inn string) error {
	weights := [9]int{2, 4, 10, 3, 5, 9, 4, 6, 8}

	if !regexp.MustCompile("^[0-9]{10}$").MatchString(inn) {
	  return fmt.Errorf("%s is not a legal person inn", inn)
	}
  
	array, err := intStringToIntArray(inn)
	if err != nil {
	  return fmt.Errorf("%s is not a legal person inn", inn)
	}
  
	var result int
	for index, weight := range weights {
	  result += array[index] * weight
	}
	checksum := result % 11 % 10
	arrayLastElement := array[9]
  
	if checksum != arrayLastElement {
	  return fmt.Errorf("%s is not a legal person inn", inn)
	}

	return nil
}

func intStringToIntArray(inputString string) ([]int, error) {
  if inputString == "" {
    return nil, fmt.Errorf("%s is empty", inputString)
  }

  res := make([]int, 10)

  for i, value := range inputString {
    r := string(value)

    // strconv.Atoi return integer and nil error
    // if all chars in string is integer
    if digit, err := strconv.Atoi(r); err != nil {
      return nil, fmt.Errorf("%s in %s is not an int", r, inputString)
    }else{
		res[i] = digit
	}   
  }

  return res, nil
}