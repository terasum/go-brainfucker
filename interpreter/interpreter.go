//MIT License
//
//Copyright (c) 2018 ChenQuan
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package interpreter
/*
   BrainFuck the language, which created by Urban Mueller
   Turing complete language
   implements in go
   +-------------------------------------------------------+
   | > | pointer + 1                                       |
   | < | pointer -1                                        |
   | + | pointer's byte + 1                                |
   | - | pointer's byte - 1                                |
   | . | output the pointer's byte (ASCII code)            |
   | , | input byte into pointer's memory                  |
   | [ | if pointer's memory is zero, jump forward to `]`  |
   | ] | if pointer's memory is not zero, jump back to `[` |
   +-------------------------------------------------------+
*/

import (
	"bytes"
	"fmt"
)
// Interpreter read and run BrainFuck code
type Interpreter struct {
	code []byte
	// input/output stream bytes
	input  []byte
	output []byte
	// pointers
	pointer     int64
	codePointer int64
	// data sets
	dataSet []byte
	// bracket Stack
	bracketStack []int64

	// debug or not
	debug bool
}
// NewInterpreter return new interpreter instance
func NewInterpreter(code []byte, isdebug bool) *Interpreter {
	return &Interpreter{
		code:         code,
		input:        make([]byte, 0),
		output:       make([]byte, 0),
		dataSet:      make([]byte, 0),
		bracketStack: make([]int64, 0),
		debug:        isdebug,
	}
}

func (inter *Interpreter) operate(opcode byte) {
	inter.resizeMem(inter.pointer)
	if inter.debug {
		fmt.Printf("op: %s\tinter pt: %d\tds: %v\tout: %v\n", string(opcode), inter.pointer, inter.dataSet, inter.output)
	}
	switch opcode {
	case '+':
		inter.dataSet[inter.pointer] = byte(int(inter.dataSet[inter.pointer]) + 1)
	case '-':
		inter.dataSet[inter.pointer]--
	case '<':
		if inter.pointer > 0 {
			inter.pointer--
		}
	case '>':
		inter.pointer++
	case '.':
		inter.output = append(inter.output, inter.dataSet[inter.pointer])
	case ',':
		inter.dataSet[inter.pointer] = inter.input[len(inter.input)]
		inter.input = inter.input[:len(inter.input)-1]
	case '[':
		inter.leftBracket()
	case ']':
		inter.rightBracket()
	}
}

func (inter *Interpreter) leftBracket() {
	openBrackets := 1
	if inter.pointer < int64(len(inter.dataSet)) && inter.dataSet[inter.pointer] != 0 {
		inter.bracketStack = append(inter.bracketStack, inter.codePointer)
	} else {
		for openBrackets > 0 && ((inter.codePointer + 1) < int64(len(inter.code))) {
			inter.codePointer++
			if inter.code[inter.codePointer] == ']' {
				openBrackets--
			} else if inter.code[inter.codePointer] == '[' {
				openBrackets--
			}

		}
	}
}

func (inter *Interpreter) rightBracket() {
	inter.codePointer = inter.bracketStack[len(inter.bracketStack)-1] - 1
	inter.bracketStack = inter.bracketStack[:len(inter.bracketStack)-1]
}

func (inter *Interpreter) resizeMem(pointer int64) {
	if inter.dataSet == nil || len(inter.dataSet) <= 0 {
		inter.dataSet = make([]byte, 1)
	} else if int64(len(inter.dataSet)) <= pointer {
		inter.dataSet = append(inter.dataSet, make([]byte, pointer, pointer+1)...)
	}
	//else do nothing
}
// Run the BrainFuncker code
func (inter *Interpreter) Run() (string, error) {
	list := []byte{'+', '-', '<', '>', '.', ',', '[', ']'}
	for {
		c := inter.code[inter.codePointer]
		if bytes.Contains(list, []byte{c}) {
			inter.operate(c)
		}
		inter.codePointer++
		if inter.codePointer >= int64(len(inter.code)) {
			break
		}
	}
	return string(inter.output), nil
}
