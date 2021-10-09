package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"unicode"
)

var counter int
var mySlice []byte

const LETTER int = 0
const DIGIT int = 1
const UNKNOWN int = 99

const INT_LIT int = 10
const IDENT int = 11
const ADD_OP int = 21
const SUB_OP int = 22
const MULT_OP int = 23
const DIV_OP int = 24
const LEFT_PAREN int = 25
const RIGHT_PAREN int = 26
//const EOF int = -1

var charClass int
var lexeme string

var nextChar byte
var token int
var nextToken int
var turnOff bool

func main() {
  data, err := ioutil.ReadFile("/home/cameron/Downloads/front.in")
  if err == io.EOF {
    fmt.Println("Reading file finished...")
    return
  }

  mySlice = data[:]
  my_getChar()
  for !turnOff {
    lex()
  }

  
}

func my_addChar() {
  tmpString := string(nextChar)
  lexeme += tmpString
}

func my_getChar() {
  nextChar = mySlice[counter]

  if unicode.IsLetter(rune(nextChar)) {
    charClass = LETTER
  } else if unicode.IsDigit(rune(nextChar)){
    charClass = DIGIT
  } else {
    charClass = UNKNOWN
  }
  
  // Increment global counter 
  counter++
}

func getNonBlank() {
  for unicode.IsSpace(rune(nextChar)) {
    my_getChar()
  }
}

func lex() {
  lexeme = " "
  getNonBlank()
  switch charClass {
  case LETTER:
    my_addChar()
    my_getChar()
    for charClass == LETTER || charClass == DIGIT {
      my_addChar()
      my_getChar()
    }
    nextToken = IDENT
  case DIGIT:
    my_addChar()
    my_getChar()
    for charClass == DIGIT {
      my_addChar()
      my_getChar()
    }
    nextToken = INT_LIT
  case UNKNOWN:
    lookup(nextChar)
    my_getChar()

  }

  fmt.Println("Next token is: ", nextToken, ", Next lexeme is: ", string(lexeme))
  
  if counter == len(mySlice) {
    turnOff = true
    fmt.Println("Next token is:  -1 , Next lexeme is:   EOF")
  }
  
}

func lookup(nextChar byte) {
  switch nextChar {
  case '(':
    my_addChar()
    nextToken = LEFT_PAREN
  case ')':
    my_addChar()
    nextToken = RIGHT_PAREN
  case '+':
    my_addChar()
    nextToken = ADD_OP
  case '-':
    my_addChar()
    nextToken = SUB_OP
  case '*':
    my_addChar()
    nextToken = MULT_OP
  case '/':
    my_addChar()
    nextToken = DIV_OP
  }
}
 