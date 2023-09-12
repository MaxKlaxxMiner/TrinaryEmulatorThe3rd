#include "go_asm.h"
#include "asm_amd64.h"
#include "textflag.h"

// RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11
// R12, R13

// func Gp3(u Uint27) (hi Uint9, mid Uint9, low Uint9) {
TEXT ·Gp3<ABIInternal>(SB),NOSPLIT,$0
    // AX = u
    MOVQ	AX, CX
	MOVQ	$-5665385581486077344, AX
	MULQ	CX
	SHRQ	$28, DX
	IMUL3Q	$387420489, DX, SI
	SUBQ	SI, CX
	MOVL	$3575102585, BX
	IMULQ	CX, BX
	SHRQ	$46, BX
	IMUL3Q	$19683, BX, SI
	SUBQ	SI, CX
    MOVQ	DX, AX

    // AX = hi
    // BX = mid
    // CX = low

    RET
