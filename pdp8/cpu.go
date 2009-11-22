package pdp8

type Word uint16

const (
	opAND = iota;
	opTAD;
	opISZ;
	opDCA;
	opJMS;
	opJMP;
	opIOT;
	opOP;
)

type registers struct {
	AC int;
	memoryAddress int;
	memoryBuffer int;
	PC int;
	effectiveAddress int;
	switchRegister int;
	instructionRegister int;
}

type CPU struct {
	reg registers;
	mem [4096]int;
}

func (m *CPU) Step () {
	inst := m.mem[m.reg.PC];
	m.PC++;
	m.run(inst);
}

func (m *CPU) run (inst Word) {
	op := inst >> 9;
	offset := inst & 0x7f;
	switch op {
	case opAND:
	case opTAD:
	case opISZ:
	case opJMS:
	case opJMP:
	case opIOT:
	case opOP:}
}