package mcs6502

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type rule struct {
	Opcode         Opcode
	Length         int
	AddressingMode AddressingMode
}

var allRules = []rule{
	{OpcodeBrk, 1, AddressingModeNone},
	{OpcodeOraIndirectX, 2, AddressingModeIndirectX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeOraZeroPage, 2, AddressingModeZeroPage},
	{OpcodeAslZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodePhp, 1, AddressingModeNone},
	{OpcodeOraImmediate, 2, AddressingModeImmediate},
	{OpcodeAslAccumulator, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeOraAbsolute, 3, AddressingModeAbsolute},
	{OpcodeAslAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBpl, 3, AddressingModeNone},
	{OpcodeOraIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeOraZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeAslZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeClc, 1, AddressingModeNone},
	{OpcodeOraAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeOraAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeAslAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeJsr, 3, AddressingModeNone},
	{OpcodeAndIndirectX, 2, AddressingModeIndirectX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBitZeroPage, 2, AddressingModeZeroPage},
	{OpcodeAndZeroPage, 2, AddressingModeZeroPage},
	{OpcodeRolZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodePlp, 1, AddressingModeNone},
	{OpcodeAndImmediate, 2, AddressingModeImmediate},
	{OpcodeRolAccumulator, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBitAbsolute, 3, AddressingModeAbsolute},
	{OpcodeAndAbsolute, 3, AddressingModeAbsolute},
	{OpcodeRolAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBmi, 3, AddressingModeNone},
	{OpcodeAndIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeAndZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeRolZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeSec, 1, AddressingModeNone},
	{OpcodeAndAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeAndAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeRolAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeRti, 1, AddressingModeNone},
	{OpcodeEorIndirectX, 2, AddressingModeIndirectX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeEorZeroPage, 2, AddressingModeZeroPage},
	{OpcodeLsrZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodePha, 1, AddressingModeNone},
	{OpcodeEorImmediate, 2, AddressingModeImmediate},
	{OpcodeLsrAccumulator, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeJmpAbsolute, 3, AddressingModeNone},
	{OpcodeEorAbsolute, 3, AddressingModeAbsolute},
	{OpcodeLsrAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBvc, 3, AddressingModeNone},
	{OpcodeEorIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeEorZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeLsrZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCli, 1, AddressingModeNone},
	{OpcodeEorAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeEorAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeLsrAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeRts, 1, AddressingModeNone},
	{OpcodeAdcIndirectX, 2, AddressingModeIndirectX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeAdcZeroPage, 2, AddressingModeZeroPage},
	{OpcodeRorZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodePla, 1, AddressingModeNone},
	{OpcodeAdcImmediate, 2, AddressingModeImmediate},
	{OpcodeRorAccumulator, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeJmpIndirect, 3, AddressingModeNone},
	{OpcodeAdcAbsolute, 3, AddressingModeAbsolute},
	{OpcodeRorAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBvs, 3, AddressingModeNone},
	{OpcodeAdcIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeAdcZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeRorZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeSei, 1, AddressingModeNone},
	{OpcodeAdcAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeAdcAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeRorAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeStaIndirectX, 2, AddressingModeIndirectX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeStyZeroPage, 2, AddressingModeNone},
	{OpcodeStaZeroPage, 2, AddressingModeZeroPage},
	{OpcodeStxZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeDey, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeTxa, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeStyAbsolute, 3, AddressingModeNone},
	{OpcodeStaAbsolute, 3, AddressingModeAbsolute},
	{OpcodeStxAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBcc, 3, AddressingModeNone},
	{OpcodeStaIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeStyZeroPageX, 2, AddressingModeNone},
	{OpcodeStaZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeStxZeroPageY, 2, AddressingModeZeroPageY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeTya, 1, AddressingModeNone},
	{OpcodeStaAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeTxs, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeStaAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeLdyImmediate, 2, AddressingModeImmediate},
	{OpcodeLdaIndirectX, 2, AddressingModeIndirectX},
	{OpcodeLdxImmediate, 2, AddressingModeImmediate},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeLdyZeroPage, 2, AddressingModeZeroPage},
	{OpcodeLdaZeroPage, 2, AddressingModeZeroPage},
	{OpcodeLdxZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeTay, 1, AddressingModeNone},
	{OpcodeLdaImmediate, 2, AddressingModeImmediate},
	{OpcodeTax, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeLdyAbsolute, 3, AddressingModeAbsolute},
	{OpcodeLdaAbsolute, 3, AddressingModeAbsolute},
	{OpcodeLdxAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBcs, 3, AddressingModeNone},
	{OpcodeLdaIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeLdyZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeLdaZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeLdxZeroPageY, 2, AddressingModeZeroPageY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeClv, 1, AddressingModeNone},
	{OpcodeLdaAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeTsx, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeLdyAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeLdaAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeLdxAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCpyImmediate, 2, AddressingModeImmediate},
	{OpcodeCmpIndirectX, 2, AddressingModeIndirectX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCpyZeroPage, 2, AddressingModeZeroPage},
	{OpcodeCmpZeroPage, 2, AddressingModeZeroPage},
	{OpcodeDecZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeIny, 1, AddressingModeNone},
	{OpcodeCmpImmediate, 2, AddressingModeImmediate},
	{OpcodeDex, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCpyAbsolute, 3, AddressingModeAbsolute},
	{OpcodeCmpAbsolute, 3, AddressingModeAbsolute},
	{OpcodeDecAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBne, 3, AddressingModeNone},
	{OpcodeCmpIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCmpZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeDecZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCld, 1, AddressingModeNone},
	{OpcodeCmpAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCmpAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeDecAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCpxImmediate, 2, AddressingModeImmediate},
	{OpcodeSbcIndirectX, 2, AddressingModeIndirectX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCpxZeroPage, 2, AddressingModeZeroPage},
	{OpcodeSbcZeroPage, 2, AddressingModeZeroPage},
	{OpcodeIncZeroPage, 2, AddressingModeZeroPage},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInx, 1, AddressingModeNone},
	{OpcodeSbcImmediate, 2, AddressingModeNone},
	{OpcodeNop, 1, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeCpxAbsolute, 3, AddressingModeAbsolute},
	{OpcodeSbcAbsolute, 3, AddressingModeAbsolute},
	{OpcodeIncAbsolute, 3, AddressingModeAbsolute},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeBeq, 3, AddressingModeNone},
	{OpcodeSbcIndirectY, 2, AddressingModeIndirectY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeSbcZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeIncZeroPageX, 2, AddressingModeZeroPageX},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeSed, 1, AddressingModeNone},
	{OpcodeSbcAbsoluteY, 3, AddressingModeAbsoluteY},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeInvalid, 0, AddressingModeNone},
	{OpcodeSbcAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeIncAbsoluteX, 3, AddressingModeAbsoluteX},
	{OpcodeInvalid, 0, AddressingModeNone},
}

func Parse(r io.Reader) ([]Instruction, error) {
	reader := bufio.NewReader(r)
	instructions := make([]Instruction, 0)
	address := 0

	for {
		opcode, err := reader.ReadByte()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		matchingRule := allRules[opcode]
		if matchingRule.Opcode == OpcodeInvalid {
			return nil, fmt.Errorf("invalid opcode %02x", opcode)
		}

		instruction := Instruction{
			Address:        uint16(address),
			Opcode:         matchingRule.Opcode,
			AddressingMode: matchingRule.AddressingMode,
		}

		if matchingRule.Length == 1 {
			// nothing to do
		} else if matchingRule.Length == 2 {
			if operand, err := reader.ReadByte(); err == nil {
				instruction.Operand = uint16(operand)
			}
		} else if matchingRule.Length == 3 {
			var operand uint16
			if err := binary.Read(reader, binary.LittleEndian, &operand); err == nil {
				instruction.Operand = operand
			}
		} else {
			return nil, fmt.Errorf("invalid instruction length for opcode %02x", opcode)
		}

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction)
		address += matchingRule.Length
	}

	return instructions, nil
}
