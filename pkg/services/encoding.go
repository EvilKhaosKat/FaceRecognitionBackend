package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Encoding []float64

//TODO custom errors?

func NewEncoding(encoding string) (Encoding, error) {
	encodingStringLen := utf8.RuneCountInString(encoding)

	if encodingStringLen < 42 { //usually encoding is 128 float numbers, so it's pretty long in text form
		return nil, errors.New(fmt.Sprint("Too short encoding, probably incorrect data:", encoding))
	}

	numbers := strings.Split(encoding[2:encodingStringLen-2], " ")

	var result Encoding

	for _, numberStr := range numbers {
		if utf8.RuneCountInString(numberStr) == 0 {
			continue
		}

		numberStr = strings.Trim(numberStr, "\n")
		number, err := strconv.ParseFloat(numberStr, 16) //technically we get 16 bit length from ML model
		if err != nil {
			return nil, err
		}

		result = append(result, number)
	}

	return Encoding(result), nil
}
