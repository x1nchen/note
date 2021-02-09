
"".add STEXT nosplit size=60 args=0x18 locals=0x10
	0x0000 00000 (main.go:3)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $16-24
	0x0000 00000 (main.go:3)	SUBQ	$16, SP
	0x0004 00004 (main.go:3)	MOVQ	BP, 8(SP)
	0x0009 00009 (main.go:3)	LEAQ	8(SP), BP
	0x000e 00014 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:3)	MOVQ	$0, "".~r2+40(SP)
	0x0017 00023 (main.go:4)	MOVQ	$0, "".sum(SP)
	0x001f 00031 (main.go:5)	MOVQ	"".a+24(SP), AX
	0x0024 00036 (main.go:5)	ADDQ	"".b+32(SP), AX
	0x0029 00041 (main.go:5)	MOVQ	AX, "".sum(SP)
	0x002d 00045 (main.go:6)	MOVQ	AX, "".~r2+40(SP)
	0x0032 00050 (main.go:6)	MOVQ	8(SP), BP
	0x0037 00055 (main.go:6)	ADDQ	$16, SP
	0x003b 00059 (main.go:6)	RET
	0x0000 48 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 48 c7  H...H.l$.H.l$.H.
	0x0010 44 24 28 00 00 00 00 48 c7 04 24 00 00 00 00 48  D$(....H..$....H
	0x0020 8b 44 24 18 48 03 44 24 20 48 89 04 24 48 89 44  .D$.H.D$ H..$H.D
	0x0030 24 28 48 8b 6c 24 08 48 83 c4 10 c3              $(H.l$.H....
"".main STEXT size=110 args=0x0 locals=0x28
	0x0000 00000 (main.go:9)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (main.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:9)	PCDATA	$0, $-2
	0x000d 00013 (main.go:9)	JLS	103
	0x000f 00015 (main.go:9)	PCDATA	$0, $-1
	0x000f 00015 (main.go:9)	SUBQ	$40, SP
	0x0013 00019 (main.go:9)	MOVQ	BP, 32(SP)
	0x0018 00024 (main.go:9)	LEAQ	32(SP), BP
	0x001d 00029 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:10)	MOVQ	$1, (SP)
	0x0025 00037 (main.go:10)	MOVQ	$2, 8(SP)
	0x002e 00046 (main.go:10)	PCDATA	$1, $0
	0x002e 00046 (main.go:10)	CALL	"".add(SB)
	0x0033 00051 (main.go:10)	MOVQ	16(SP), AX
	0x0038 00056 (main.go:10)	MOVQ	AX, ""..autotmp_0+24(SP)
	0x003d 00061 (main.go:10)	NOP
	0x0040 00064 (main.go:10)	CALL	runtime.printlock(SB)
	0x0045 00069 (main.go:10)	MOVQ	""..autotmp_0+24(SP), AX
	0x004a 00074 (main.go:10)	MOVQ	AX, (SP)
	0x004e 00078 (main.go:10)	CALL	runtime.printint(SB)
	0x0053 00083 (main.go:10)	CALL	runtime.printnl(SB)
	0x0058 00088 (main.go:10)	CALL	runtime.printunlock(SB)
	0x005d 00093 (main.go:11)	MOVQ	32(SP), BP
	0x0062 00098 (main.go:11)	ADDQ	$40, SP
	0x0066 00102 (main.go:11)	RET
	0x0067 00103 (main.go:11)	NOP
	0x0067 00103 (main.go:9)	PCDATA	$1, $-1
	0x0067 00103 (main.go:9)	PCDATA	$0, $-2
	0x0067 00103 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x006c 00108 (main.go:9)	PCDATA	$0, $-1
	0x006c 00108 (main.go:9)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 58 48  eH..%....H;a.vXH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 04  ..(H.l$ H.l$ H..
	0x0020 24 01 00 00 00 48 c7 44 24 08 02 00 00 00 e8 00  $....H.D$.......
	0x0030 00 00 00 48 8b 44 24 10 48 89 44 24 18 0f 1f 00  ...H.D$.H.D$....
	0x0040 e8 00 00 00 00 48 8b 44 24 18 48 89 04 24 e8 00  .....H.D$.H..$..
	0x0050 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8b 6c  .............H.l
	0x0060 24 20 48 83 c4 28 c3 e8 00 00 00 00 eb 92        $ H..(........
	rel 5+4 t=17 TLS+0
	rel 47+4 t=8 "".add+0
	rel 65+4 t=8 runtime.printlock+0
	rel 79+4 t=8 runtime.printint+0
	rel 84+4 t=8 runtime.printnl+0
	rel 89+4 t=8 runtime.printunlock+0
	rel 104+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
