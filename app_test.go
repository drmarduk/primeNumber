package main

import(
  "testing"
  "reflect"
  "strconv"
)

func TestParseInput(t *testing.T) {
  goodInt, err := parseInput("10")
  expect(t, goodInt, 10)
  expect(t, err, nil)

  badInt, err := parseInput("bad int")
  _, expectedError := strconv.Atoi("bad int")
  expect(t, badInt, 0)
  expect(t, reflect.TypeOf(err), reflect.TypeOf(expectedError))
}

func TestWork(t *testing.T) {
  third_prim := work(3)
  expect(t, third_prim, 5)
  primeNumber = primeNumber[:0]
  hundredth_prim := work(100)
  expect(t, hundredth_prim, 541)
  primeNumber = primeNumber[:0]
}

func TestPrim(t *testing.T) {
  prim(547)
  expect(t, primeNumber[0], 547)
}

func expect(t *testing.T, a interface{}, b interface{}) {
  if a != b {
    t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
  }
}