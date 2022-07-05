package code

import (
	"encoding/binary"
	"fmt"
)

// Instructions themselves are a series of bytes and
// a single instruction consists of an opcode
// and an optional number of operands.
type Instructions []byte

// Opcode is exactly one byte wide, has an arbitrary but unique value
// and is the first byte in the instruction.
type Opcode byte

const (
	OpConstant Opcode = iota
)

// Definition for debugging and testing purposes, Name helps to make an Opcode readable
// and OperandWidths contains the number of bytes each operand takes up.
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}
