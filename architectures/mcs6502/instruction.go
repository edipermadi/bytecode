package mcs6502

type Instruction struct {
	Address        uint16
	Opcode         Opcode
	Operand        uint16
	AddressingMode AddressingMode
}
