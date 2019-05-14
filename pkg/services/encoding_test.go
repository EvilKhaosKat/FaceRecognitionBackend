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
		t.Error("want nil, got", err)
	}

	if len(encoding) != 6 {
		t.Error("want 6; got", len(encoding))
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
		t.Error("want nil; got", encoding)
	}

	if err == nil {
		t.Error("want error, got nil")
	}
}

func TestNewEncodingNotNumber(t *testing.T) {
	encodingStr := "[[-0.02517794  0.12061624  0.04272895 -0.0952117  -0.00082525  0.09786011 NOT_NUMBER_AT_ALL]]"

	encoding, err := NewEncoding(encodingStr)

	if encoding != nil {
		t.Error("want nil; got", encoding)
	}

	if err == nil {
		t.Error("want error, got nil")
	}
}

func TestGetDist(t *testing.T) {
	first, _ := NewEncoding(getFirstEncodingString())

	dist, err := first.GetDist(first)

	if err != nil {
		t.Fatal("want nil; got", err)
	}

	if dist > 0.00000000001 {
		t.Fatal("distance between same encodings is not close to zero")
	}
}

func TestIsSame(t *testing.T) {
	first, _ := NewEncoding(getFirstEncodingString())
	second, _ := NewEncoding(getSecondEncodingString())

	isSame, err := first.IsSame(second)

	if err != nil {
		t.Fatal("want nil; got", err)
	}

	if !isSame {
		t.Fatal("want same; got not same")
	}
}

func getFirstEncodingString() string {
	return `[[-0.02517794  0.12061624  0.04272895 -0.0952117  -0.00082525
   0.11663063  0.10064542 -0.12526244  0.05244305  0.09572414  0.13917078
   0.04654023 -0.07373067  0.0774107   0.08275025 -0.00226332  0.11106832
  -0.16339034  0.08947118 -0.06504941  0.05461143  0.06309925 -0.04119856
  -0.1263787   0.04763409  0.0664267  -0.14906143  0.07862936  0.03313464
   0.05804996  0.05032563 -0.09558917  0.03146945 -0.08778512  0.00032996
  -0.1676044   0.05428734  0.20914543  0.0087307  -0.02346108  0.10011243
   0.12281854 -0.02066345  0.08955534  0.12052457 -0.0492502  -0.05357183
  -0.11322106 -0.04778777 -0.14879316  0.09299239 -0.06210891 -0.0606531
   0.04096187 -0.07158172]]`
}

func getSecondEncodingString() string {
	return `[[-0.04781156  0.13086452  0.06519756 -0.06235839  0.01312022
   0.00484892  0.06017131 -0.07756685 -0.01801117  0.07118477  0.12083602
   0.15290402  0.03152245  0.01988193  0.10849094  0.0032022   0.11602321
  -0.16587712  0.03919415 -0.12275632 -0.04176315  0.14333847 -0.04478907
  -0.0008282   0.06668508 -0.06372692  0.10177787 -0.06872098 -0.08778807
   0.07202978  0.06168408 -0.17359063 -0.03584487 -0.07363114  0.00227653
  -0.14662246  0.04413157  0.20206797  0.03183846 -0.03969065 -0.07047424
   0.13661303 -0.06639413  0.1545957   0.11782251 -0.07934082 -0.01013817
  -0.13595062  0.03486867 -0.10575092  0.1386303  -0.07088415 -0.01846558
  -0.02831654 -0.09127404]]`
}
