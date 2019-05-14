package services

import (
	"math/big"
	"testing"
)

func TestNewEncoding(t *testing.T) {
	encodingStr := "[[-0.02517794  0.12061624  0.04272895 -0.0952117  -0.00082525  0.09786011]]"

	encoding, err := NewEncoding(encodingStr)

	if encoding == nil {
		t.Error("want encoding; got nil")
	}

	if err != nil {
		t.Error("want nil, got ", err)
	}

	if len(encoding) != 6 {
		t.Error("want 6; got ", len(encoding))
	}

	//lets assume that's enough for sanity check
	firstNum := big.NewFloat(encoding[0])
	firstNumTest := big.NewFloat(-0.02517794)

	if firstNum.Cmp(firstNumTest) != 0 {
		t.Errorf("want %f; got %f", firstNumTest, firstNum)
	}
}

func TestNewEncodingShortData(t *testing.T) {
	encodingStr := "[[short input]]"

	encoding, err := NewEncoding(encodingStr)

	if encoding != nil {
		t.Error("want nil; got ", encoding)
	}

	if err == nil {
		t.Error("want error, got nil")
	}
}

func TestNewEncodingNotNumber(t *testing.T) {
	encodingStr := "[[-0.02517794  0.12061624  0.04272895 -0.0952117  -0.00082525  0.09786011 NOT_NUMBER_AT_ALL]]"

	encoding, err := NewEncoding(encodingStr)

	if encoding != nil {
		t.Error("want nil; got ", encoding)
	}

	if err == nil {
		t.Error("want error, got nil")
	}
}
