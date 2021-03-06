package vm

import (
	"fmt"
	"math/big"

	"github.com/yido/yido-chain/common"
)

func Disassemble(script []byte) (asm []string) {
	pc := new(big.Int)
	for {
		if pc.Cmp(big.NewInt(int64(len(script)))) >= 0 {
			return
		}

		val := script[pc.Int64()]

		op := OpCode(val)

		asm = append(asm, fmt.Sprintf("%v", op))

		switch op {
		case PUSH1, PUSH2, PUSH3, PUSH4, PUSH5, PUSH6, PUSH7, PUSH8, PUSH9, PUSH10, PUSH11, PUSH12, PUSH13, PUSH14, PUSH15, PUSH16, PUSH17, PUSH18, PUSH19, PUSH20, PUSH21, PUSH22, PUSH23, PUSH24, PUSH25, PUSH26, PUSH27, PUSH28, PUSH29, PUSH30, PUSH31, PUSH32:
			pc.Add(pc, common.Big1)
			a := int64(op) - int64(PUSH1) + 1
			if int(pc.Int64()+a) > len(script) {
				return nil
			}

			data := script[pc.Int64() : pc.Int64()+a]
			if len(data) == 0 {
				data = []byte{0}
			}
			asm = append(asm, fmt.Sprintf("0x%x", data))

			pc.Add(pc, big.NewInt(a-1))
		}

		pc.Add(pc, common.Big1)
	}

	return
}
