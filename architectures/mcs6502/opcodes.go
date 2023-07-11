package mcs6502

type Opcode int

const (
	OpcodeAdcImmediate   Opcode = 0x69
	OpcodeAdcZeroPage    Opcode = 0x65
	OpcodeAdcZeroPageX   Opcode = 0x75
	OpcodeAdcAbsolute    Opcode = 0x6D
	OpcodeAdcAbsoluteX   Opcode = 0x7D
	OpcodeAdcAbsoluteY   Opcode = 0x79
	OpcodeAdcIndirectX   Opcode = 0x61
	OpcodeAdcIndirectY   Opcode = 0x71
	OpcodeAndImmediate   Opcode = 0x29
	OpcodeAndZeroPage    Opcode = 0x25
	OpcodeAndZeroPageX   Opcode = 0x35
	OpcodeAndAbsolute    Opcode = 0x2D
	OpcodeAndAbsoluteX   Opcode = 0x3D
	OpcodeAndAbsoluteY   Opcode = 0x39
	OpcodeAndIndirectX   Opcode = 0x21
	OpcodeAndIndirectY   Opcode = 0x31
	OpcodeAslAccumulator Opcode = 0x0A
	OpcodeAslZeroPage    Opcode = 0x06
	OpcodeAslZeroPageX   Opcode = 0x16
	OpcodeAslAbsolute    Opcode = 0x0E
	OpcodeAslAbsoluteX   Opcode = 0x1E
	OpcodeBitZeroPage    Opcode = 0x24
	OpcodeBitAbsolute    Opcode = 0x2C
	OpcodeBpl            Opcode = 0x10
	OpcodeBmi            Opcode = 0x30
	OpcodeBvc            Opcode = 0x50
	OpcodeBvs            Opcode = 0x70
	OpcodeBcc            Opcode = 0x90
	OpcodeBcs            Opcode = 0xB0
	OpcodeBne            Opcode = 0xD0
	OpcodeBeq            Opcode = 0xF0
	OpcodeBrk            Opcode = 0x00
	OpcodeCmpImmediate   Opcode = 0xC9
	OpcodeCmpZeroPage    Opcode = 0xC5
	OpcodeCmpZeroPageX   Opcode = 0xD5
	OpcodeCmpAbsolute    Opcode = 0xCD
	OpcodeCmpAbsoluteX   Opcode = 0xDD
	OpcodeCmpAbsoluteY   Opcode = 0xD9
	OpcodeCmpIndirectX   Opcode = 0xC1
	OpcodeCmpIndirectY   Opcode = 0xD1
	OpcodeCpxImmediate   Opcode = 0xE0
	OpcodeCpxZeroPage    Opcode = 0xE4
	OpcodeCpxAbsolute    Opcode = 0xEC
	OpcodeCpyImmediate   Opcode = 0xC0
	OpcodeCpyZeroPage    Opcode = 0xC4
	OpcodeCpyAbsolute    Opcode = 0xCC
	OpcodeDecZeroPage    Opcode = 0xC6
	OpcodeDecZeroPageX   Opcode = 0xD6
	OpcodeDecAbsolute    Opcode = 0xCE
	OpcodeDecAbsoluteX   Opcode = 0xDE
	OpcodeEorImmediate   Opcode = 0x49
	OpcodeEorZeroPage    Opcode = 0x45
	OpcodeEorZeroPageX   Opcode = 0x55
	OpcodeEorAbsolute    Opcode = 0x4D
	OpcodeEorAbsoluteX   Opcode = 0x5D
	OpcodeEorAbsoluteY   Opcode = 0x59
	OpcodeEorIndirectX   Opcode = 0x41
	OpcodeEorIndirectY   Opcode = 0x51
	OpcodeClc            Opcode = 0x18
	OpcodeSec            Opcode = 0x38
	OpcodeCli            Opcode = 0x58
	OpcodeSei            Opcode = 0x78
	OpcodeClv            Opcode = 0xB8
	OpcodeCld            Opcode = 0xD8
	OpcodeSed            Opcode = 0xF8
	OpcodeIncZeroPage    Opcode = 0xE6
	OpcodeIncZeroPageX   Opcode = 0xF6
	OpcodeIncAbsolute    Opcode = 0xEE
	OpcodeIncAbsoluteX   Opcode = 0xFE
	OpcodeJmpAbsolute    Opcode = 0x4C
	OpcodeJmpIndirect    Opcode = 0x6C
	OpcodeJsr            Opcode = 0x20
	OpcodeLdaImmediate   Opcode = 0xA9
	OpcodeLdaZeroPage    Opcode = 0xA5
	OpcodeLdaZeroPageX   Opcode = 0xB5
	OpcodeLdaAbsolute    Opcode = 0xAD
	OpcodeLdaAbsoluteX   Opcode = 0xBD
	OpcodeLdaAbsoluteY   Opcode = 0xB9
	OpcodeLdaIndirectX   Opcode = 0xA1
	OpcodeLdaIndirectY   Opcode = 0xB1
	OpcodeLdxImmediate   Opcode = 0xA2
	OpcodeLdxZeroPage    Opcode = 0xA6
	OpcodeLdxZeroPageY   Opcode = 0xB6
	OpcodeLdxAbsolute    Opcode = 0xAE
	OpcodeLdxAbsoluteY   Opcode = 0xBE
	OpcodeLdyImmediate   Opcode = 0xA0
	OpcodeLdyZeroPage    Opcode = 0xA4
	OpcodeLdyZeroPageX   Opcode = 0xB4
	OpcodeLdyAbsolute    Opcode = 0xAC
	OpcodeLdyAbsoluteX   Opcode = 0xBC
	OpcodeLsrAccumulator Opcode = 0x4A
	OpcodeLsrZeroPage    Opcode = 0x46
	OpcodeLsrZeroPageX   Opcode = 0x56
	OpcodeLsrAbsolute    Opcode = 0x4E
	OpcodeLsrAbsoluteX   Opcode = 0x5E
	OpcodeNop            Opcode = 0xEA
	OpcodeOraImmediate   Opcode = 0x09
	OpcodeOraZeroPage    Opcode = 0x05
	OpcodeOraZeroPageX   Opcode = 0x15
	OpcodeOraAbsolute    Opcode = 0x0D
	OpcodeOraAbsoluteX   Opcode = 0x1D
	OpcodeOraAbsoluteY   Opcode = 0x19
	OpcodeOraIndirectX   Opcode = 0x01
	OpcodeOraIndirectY   Opcode = 0x11
	OpcodeTax            Opcode = 0xAA
	OpcodeTxa            Opcode = 0x8A
	OpcodeDex            Opcode = 0xCA
	OpcodeInx            Opcode = 0xE8
	OpcodeTay            Opcode = 0xA8
	OpcodeTya            Opcode = 0x98
	OpcodeDey            Opcode = 0x88
	OpcodeIny            Opcode = 0xC8
	OpcodeRolAccumulator Opcode = 0x2A
	OpcodeRolZeroPage    Opcode = 0x26
	OpcodeRolZeroPageX   Opcode = 0x36
	OpcodeRolAbsolute    Opcode = 0x2E
	OpcodeRolAbsoluteX   Opcode = 0x3E
	OpcodeRorAccumulator Opcode = 0x6A
	OpcodeRorZeroPage    Opcode = 0x66
	OpcodeRorZeroPageX   Opcode = 0x76
	OpcodeRorAbsolute    Opcode = 0x6E
	OpcodeRorAbsoluteX   Opcode = 0x7E
	OpcodeRti            Opcode = 0x40
	OpcodeRts            Opcode = 0x60
	OpcodeSbcImmediate   Opcode = 0xE9
	OpcodeSbcZeroPage    Opcode = 0xE5
	OpcodeSbcZeroPageX   Opcode = 0xF5
	OpcodeSbcAbsolute    Opcode = 0xED
	OpcodeSbcAbsoluteX   Opcode = 0xFD
	OpcodeSbcAbsoluteY   Opcode = 0xF9
	OpcodeSbcIndirectX   Opcode = 0xE1
	OpcodeSbcIndirectY   Opcode = 0xF1
	OpcodeStaZeroPage    Opcode = 0x85
	OpcodeStaZeroPageX   Opcode = 0x95
	OpcodeStaAbsolute    Opcode = 0x8D
	OpcodeStaAbsoluteX   Opcode = 0x9D
	OpcodeStaAbsoluteY   Opcode = 0x99
	OpcodeStaIndirectX   Opcode = 0x81
	OpcodeStaIndirectY   Opcode = 0x91
	OpcodeTxs            Opcode = 0x9A
	OpcodeTsx            Opcode = 0xBA
	OpcodePha            Opcode = 0x48
	OpcodePla            Opcode = 0x68
	OpcodePhp            Opcode = 0x08
	OpcodePlp            Opcode = 0x28
	OpcodeStxZeroPage    Opcode = 0x86
	OpcodeStxZeroPageY   Opcode = 0x96
	OpcodeStxAbsolute    Opcode = 0x8E
	OpcodeStyZeroPage    Opcode = 0x84
	OpcodeStyZeroPageX   Opcode = 0x94
	OpcodeStyAbsolute    Opcode = 0x8C
	OpcodeInvalid        Opcode = -1
)

var opcodeNames = map[Opcode]string{
	OpcodeAdcImmediate:   "OpcodeAdcImmediate",
	OpcodeAdcZeroPage:    "OpcodeAdcZeroPage",
	OpcodeAdcZeroPageX:   "OpcodeAdcZeroPageX",
	OpcodeAdcAbsolute:    "OpcodeAdcAbsolute",
	OpcodeAdcAbsoluteX:   "OpcodeAdcAbsoluteX",
	OpcodeAdcAbsoluteY:   "OpcodeAdcAbsoluteY",
	OpcodeAdcIndirectX:   "OpcodeAdcIndirectX",
	OpcodeAdcIndirectY:   "OpcodeAdcIndirectY",
	OpcodeAndImmediate:   "OpcodeAndImmediate",
	OpcodeAndZeroPage:    "OpcodeAndZeroPage",
	OpcodeAndZeroPageX:   "OpcodeAndZeroPageX",
	OpcodeAndAbsolute:    "OpcodeAndAbsolute",
	OpcodeAndAbsoluteX:   "OpcodeAndAbsoluteX",
	OpcodeAndAbsoluteY:   "OpcodeAndAbsoluteY",
	OpcodeAndIndirectX:   "OpcodeAndIndirectX",
	OpcodeAndIndirectY:   "OpcodeAndIndirectY",
	OpcodeAslAccumulator: "OpcodeAslAccumulator",
	OpcodeAslZeroPage:    "OpcodeAslZeroPage",
	OpcodeAslZeroPageX:   "OpcodeAslZeroPageX",
	OpcodeAslAbsolute:    "OpcodeAslAbsolute",
	OpcodeAslAbsoluteX:   "OpcodeAslAbsoluteX",
	OpcodeBitZeroPage:    "OpcodeBitZeroPage",
	OpcodeBitAbsolute:    "OpcodeBitAbsolute",
	OpcodeBpl:            "OpcodeBpl",
	OpcodeBmi:            "OpcodeBmi",
	OpcodeBvc:            "OpcodeBvc",
	OpcodeBvs:            "OpcodeBvs",
	OpcodeBcc:            "OpcodeBcc",
	OpcodeBcs:            "OpcodeBcs",
	OpcodeBne:            "OpcodeBne",
	OpcodeBeq:            "OpcodeBeq",
	OpcodeBrk:            "OpcodeBrk",
	OpcodeCmpImmediate:   "OpcodeCmpImmediate",
	OpcodeCmpZeroPage:    "OpcodeCmpZeroPage",
	OpcodeCmpZeroPageX:   "OpcodeCmpZeroPageX",
	OpcodeCmpAbsolute:    "OpcodeCmpAbsolute",
	OpcodeCmpAbsoluteX:   "OpcodeCmpAbsoluteX",
	OpcodeCmpAbsoluteY:   "OpcodeCmpAbsoluteY",
	OpcodeCmpIndirectX:   "OpcodeCmpIndirectX",
	OpcodeCmpIndirectY:   "OpcodeCmpIndirectY",
	OpcodeCpxImmediate:   "OpcodeCpxImmediate",
	OpcodeCpxZeroPage:    "OpcodeCpxZeroPage",
	OpcodeCpxAbsolute:    "OpcodeCpxAbsolute",
	OpcodeCpyImmediate:   "OpcodeCpyImmediate",
	OpcodeCpyZeroPage:    "OpcodeCpyZeroPage",
	OpcodeCpyAbsolute:    "OpcodeCpyAbsolute",
	OpcodeDecZeroPage:    "OpcodeDecZeroPage",
	OpcodeDecZeroPageX:   "OpcodeDecZeroPageX",
	OpcodeDecAbsolute:    "OpcodeDecAbsolute",
	OpcodeDecAbsoluteX:   "OpcodeDecAbsoluteX",
	OpcodeEorImmediate:   "OpcodeEorImmediate",
	OpcodeEorZeroPage:    "OpcodeEorZeroPage",
	OpcodeEorZeroPageX:   "OpcodeEorZeroPageX",
	OpcodeEorAbsolute:    "OpcodeEorAbsolute",
	OpcodeEorAbsoluteX:   "OpcodeEorAbsoluteX",
	OpcodeEorAbsoluteY:   "OpcodeEorAbsoluteY",
	OpcodeEorIndirectX:   "OpcodeEorIndirectX",
	OpcodeEorIndirectY:   "OpcodeEorIndirectY",
	OpcodeClc:            "OpcodeClc",
	OpcodeSec:            "OpcodeSec",
	OpcodeCli:            "OpcodeCli",
	OpcodeSei:            "OpcodeSei",
	OpcodeClv:            "OpcodeClv",
	OpcodeCld:            "OpcodeCld",
	OpcodeSed:            "OpcodeSed",
	OpcodeIncZeroPage:    "OpcodeIncZeroPage",
	OpcodeIncZeroPageX:   "OpcodeIncZeroPageX",
	OpcodeIncAbsolute:    "OpcodeIncAbsolute",
	OpcodeIncAbsoluteX:   "OpcodeIncAbsoluteX",
	OpcodeJmpAbsolute:    "OpcodeJmpAbsolute",
	OpcodeJmpIndirect:    "OpcodeJmpIndirect",
	OpcodeJsr:            "OpcodeJsr",
	OpcodeLdaImmediate:   "OpcodeLdaImmediate",
	OpcodeLdaZeroPage:    "OpcodeLdaZeroPage",
	OpcodeLdaZeroPageX:   "OpcodeLdaZeroPageX",
	OpcodeLdaAbsolute:    "OpcodeLdaAbsolute",
	OpcodeLdaAbsoluteX:   "OpcodeLdaAbsoluteX",
	OpcodeLdaAbsoluteY:   "OpcodeLdaAbsoluteY",
	OpcodeLdaIndirectX:   "OpcodeLdaIndirectX",
	OpcodeLdaIndirectY:   "OpcodeLdaIndirectY",
	OpcodeLdxImmediate:   "OpcodeLdxImmediate",
	OpcodeLdxZeroPage:    "OpcodeLdxZeroPage",
	OpcodeLdxZeroPageY:   "OpcodeLdxZeroPageY",
	OpcodeLdxAbsolute:    "OpcodeLdxAbsolute",
	OpcodeLdxAbsoluteY:   "OpcodeLdxAbsoluteY",
	OpcodeLdyImmediate:   "OpcodeLdyImmediate",
	OpcodeLdyZeroPage:    "OpcodeLdyZeroPage",
	OpcodeLdyZeroPageX:   "OpcodeLdyZeroPageX",
	OpcodeLdyAbsolute:    "OpcodeLdyAbsolute",
	OpcodeLdyAbsoluteX:   "OpcodeLdyAbsoluteX",
	OpcodeLsrAccumulator: "OpcodeLsrAccumulator",
	OpcodeLsrZeroPage:    "OpcodeLsrZeroPage",
	OpcodeLsrZeroPageX:   "OpcodeLsrZeroPageX",
	OpcodeLsrAbsolute:    "OpcodeLsrAbsolute",
	OpcodeLsrAbsoluteX:   "OpcodeLsrAbsoluteX",
	OpcodeNop:            "OpcodeNop",
	OpcodeOraImmediate:   "OpcodeOraImmediate",
	OpcodeOraZeroPage:    "OpcodeOraZeroPage",
	OpcodeOraZeroPageX:   "OpcodeOraZeroPageX",
	OpcodeOraAbsolute:    "OpcodeOraAbsolute",
	OpcodeOraAbsoluteX:   "OpcodeOraAbsoluteX",
	OpcodeOraAbsoluteY:   "OpcodeOraAbsoluteY",
	OpcodeOraIndirectX:   "OpcodeOraIndirectX",
	OpcodeOraIndirectY:   "OpcodeOraIndirectY",
	OpcodeTax:            "OpcodeTax",
	OpcodeTxa:            "OpcodeTxa",
	OpcodeDex:            "OpcodeDex",
	OpcodeInx:            "OpcodeInx",
	OpcodeTay:            "OpcodeTay",
	OpcodeTya:            "OpcodeTya",
	OpcodeDey:            "OpcodeDey",
	OpcodeIny:            "OpcodeIny",
	OpcodeRolAccumulator: "OpcodeRolAccumulator",
	OpcodeRolZeroPage:    "OpcodeRolZeroPage",
	OpcodeRolZeroPageX:   "OpcodeRolZeroPageX",
	OpcodeRolAbsolute:    "OpcodeRolAbsolute",
	OpcodeRolAbsoluteX:   "OpcodeRolAbsoluteX",
	OpcodeRorAccumulator: "OpcodeRorAccumulator",
	OpcodeRorZeroPage:    "OpcodeRorZeroPage",
	OpcodeRorZeroPageX:   "OpcodeRorZeroPageX",
	OpcodeRorAbsolute:    "OpcodeRorAbsolute",
	OpcodeRorAbsoluteX:   "OpcodeRorAbsoluteX",
	OpcodeRti:            "OpcodeRti",
	OpcodeRts:            "OpcodeRts",
	OpcodeSbcImmediate:   "OpcodeSbcImmediate",
	OpcodeSbcZeroPage:    "OpcodeSbcZeroPage",
	OpcodeSbcZeroPageX:   "OpcodeSbcZeroPageX",
	OpcodeSbcAbsolute:    "OpcodeSbcAbsolute",
	OpcodeSbcAbsoluteX:   "OpcodeSbcAbsoluteX",
	OpcodeSbcAbsoluteY:   "OpcodeSbcAbsoluteY",
	OpcodeSbcIndirectX:   "OpcodeSbcIndirectX",
	OpcodeSbcIndirectY:   "OpcodeSbcIndirectY",
	OpcodeStaZeroPage:    "OpcodeStaZeroPage",
	OpcodeStaZeroPageX:   "OpcodeStaZeroPageX",
	OpcodeStaAbsolute:    "OpcodeStaAbsolute",
	OpcodeStaAbsoluteX:   "OpcodeStaAbsoluteX",
	OpcodeStaAbsoluteY:   "OpcodeStaAbsoluteY",
	OpcodeStaIndirectX:   "OpcodeStaIndirectX",
	OpcodeStaIndirectY:   "OpcodeStaIndirectY",
	OpcodeTxs:            "OpcodeTxs",
	OpcodeTsx:            "OpcodeTsx",
	OpcodePha:            "OpcodePha",
	OpcodePla:            "OpcodePla",
	OpcodePhp:            "OpcodePhp",
	OpcodePlp:            "OpcodePlp",
	OpcodeStxZeroPage:    "OpcodeStxZeroPage",
	OpcodeStxZeroPageY:   "OpcodeStxZeroPageY",
	OpcodeStxAbsolute:    "OpcodeStxAbsolute",
	OpcodeStyZeroPage:    "OpcodeStyZeroPage",
	OpcodeStyZeroPageX:   "OpcodeStyZeroPageX",
	OpcodeStyAbsolute:    "OpcodeStyAbsolute",
	OpcodeInvalid:        "OpcodeInvalid",
}

func (o Opcode) String() string {
	return opcodeNames[o]
}

type AddressingMode int

const (
	AddressingModeNone      AddressingMode = iota
	AddressingModeImmediate AddressingMode = iota
	AddressingModeZeroPage  AddressingMode = iota
	AddressingModeZeroPageX AddressingMode = iota
	AddressingModeZeroPageY AddressingMode = iota
	AddressingModeAbsolute  AddressingMode = iota
	AddressingModeAbsoluteX AddressingMode = iota
	AddressingModeAbsoluteY AddressingMode = iota
	AddressingModeIndirectX AddressingMode = iota
	AddressingModeIndirectY AddressingMode = iota
)

var addressingModeNames = map[AddressingMode]string{
	AddressingModeNone:      "AddressingModeNone",
	AddressingModeImmediate: "AddressingModeImmediate",
	AddressingModeZeroPage:  "AddressingModeZeroPage",
	AddressingModeZeroPageX: "AddressingModeZeroPageX",
	AddressingModeZeroPageY: "AddressingModeZeroPageY",
	AddressingModeAbsolute:  "AddressingModeAbsolute",
	AddressingModeAbsoluteX: "AddressingModeAbsoluteX",
	AddressingModeAbsoluteY: "AddressingModeAbsoluteY",
	AddressingModeIndirectX: "AddressingModeIndirectX",
	AddressingModeIndirectY: "AddressingModeIndirectY",
}

func (a AddressingMode) String() string {
	return addressingModeNames[a]
}
