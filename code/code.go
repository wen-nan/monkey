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

	// find out how long the resulting instruction is going to be.
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	// the Opcode as its first byte.
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	// we iterate over the defined OperandWidths,
	// take the matching element from operands and put it in the instruction.
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

// nicely-formatted multi-line output.
// There’s a counter at the start of each line,
// telling us which bytes we’re looking at,
// there are the opcodes in their human-readable form,
// and then there are the decoded operands.
func (ins *Instructions) String() string {
	return ""
}

// ReadOperands should then return the decoded operands
// and tell us how many bytes it read to do that.
// ReadOperands is supposed to be Make’s counterpart.
func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
