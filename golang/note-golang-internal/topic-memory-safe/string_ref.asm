"".main STEXT size=144 args=0x0 locals=0x38 funcid=0x0
	0x0000 00000 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	TEXT	"".main(SB), ABIInternal, $56-0
	0x0000 00000 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	PCDATA	$0, $-2
	0x000d 00013 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	JLS	134
	0x000f 00015 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	PCDATA	$0, $-1
	0x000f 00015 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	SUBQ	$56, SP
	0x0013 00019 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	MOVQ	BP, 48(SP)
	0x0018 00024 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	LEAQ	48(SP), BP
	0x001d 00029 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	FUNCDATA	$1, gclocals·d6c87e1e613370763384ad234357cff1(SB)
																															; 获取 RODATA 段中的字符串地址
																															; string_ref.go:4  a := "a"
	0x001d 00029 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:4)	LEAQ	go.string."a"(SB), AX ; "a" string 地址传到 AX
																															; type StringHeader struct {
																															;			Data uintptr      // uintptr is an integer type 
																															;			Len  int
																															; }
																															;
	0x0024 00036 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:4)	MOVQ	AX, "".a+32(SP) ;  AX 的内容给 a.Data (a.Data指向 字符串 go.string."a" 的地址)
	0x0029 00041 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:4)	MOVQ	$1, "".a+40(SP) ;  a.Len = 1 (a 的长度是 1) PS: $1 表示常数 1 
																															; string_ref.go:5  b := a
	0x0032 00050 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:5)	MOVQ	AX, "".b+16(SP) ;  AX 的内容给 b.Data (b.Data指向 字符串 go.string."a" 的地址)
	0x0037 00055 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:5)	MOVQ	$1, "".b+24(SP) ;  b.Len = 1 (b 的长度是 1)

																															; string_ref.go:6  a = "xor"
	0x0040 00064 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:6)	LEAQ	go.string."xor"(SB), AX ;; 这三个指令表示 xor 字符串字面量赋值给 a
	0x0047 00071 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:6)	MOVQ	AX, "".a+32(SP)
	0x004c 00076 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:6)	MOVQ	$3, "".a+40(SP)

	0x0055 00085 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	PCDATA	$1, $1   ; GC 相关
	0x0055 00085 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	CALL	runtime.printlock(SB)
	0x005a 00090 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	MOVQ	"".b+16(SP), AX
	0x005f 00095 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	MOVQ	"".b+24(SP), CX
	0x0064 00100 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	MOVQ	AX, (SP)
	0x0068 00104 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	MOVQ	CX, 8(SP)
	0x006d 00109 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	PCDATA	$1, $0
	0x006d 00109 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	CALL	runtime.printstring(SB)
	0x0072 00114 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	CALL	runtime.printnl(SB)
	0x0077 00119 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:7)	CALL	runtime.printunlock(SB)
	0x007c 00124 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:8)	MOVQ	48(SP), BP
	0x0081 00129 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:8)	ADDQ	$56, SP
	0x0085 00133 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:8)	RET
	0x0086 00134 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:8)	NOP
	0x0086 00134 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	PCDATA	$1, $-1
	0x0086 00134 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	PCDATA	$0, $-2
	0x0086 00134 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x008b 00139 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	PCDATA	$0, $-1
	0x008b 00139 (/Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-memory-safe/string_ref.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 77 48  eH..%....H;a.vwH
	0x0010 83 ec 38 48 89 6c 24 30 48 8d 6c 24 30 48 8d 05  ..8H.l$0H.l$0H..
	0x0020 00 00 00 00 48 89 44 24 20 48 c7 44 24 28 01 00  ....H.D$ H.D$(..
	0x0030 00 00 48 89 44 24 10 48 c7 44 24 18 01 00 00 00  ..H.D$.H.D$.....
	0x0040 48 8d 05 00 00 00 00 48 89 44 24 20 48 c7 44 24  H......H.D$ H.D$
	0x0050 28 03 00 00 00 e8 00 00 00 00 48 8b 44 24 10 48  (.........H.D$.H
	0x0060 8b 4c 24 18 48 89 04 24 48 89 4c 24 08 e8 00 00  .L$.H..$H.L$....
	0x0070 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8b 6c 24  ............H.l$
	0x0080 30 48 83 c4 38 c3 e8 00 00 00 00 e9 70 ff ff ff  0H..8.......p...
	rel 5+4 t=17 TLS+0
	rel 32+4 t=16 go.string."a"+0
	rel 67+4 t=16 go.string."xor"+0
	rel 86+4 t=8 runtime.printlock+0
	rel 110+4 t=8 runtime.printstring+0
	rel 115+4 t=8 runtime.printnl+0
	rel 120+4 t=8 runtime.printunlock+0
	rel 135+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFCUINFO dupok size=0
	0x0000 2d 4e 20 2d 6c                                   -N -l
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tcommand-line-arguments\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2" SRODATA dupok size=96
	0x0000 30 77 af 0c 92 74 08 02 41 e1 c1 07 e6 d6 18 e6  0w...t..A.......
	0x0010 70 61 74 68 09 63 6f 6d 6d 61 6e 64 2d 6c 69 6e  path.command-lin
	0x0020 65 2d 61 72 67 75 6d 65 6e 74 73 0a 6d 6f 64 09  e-arguments.mod.
	0x0030 63 6f 6d 6d 61 6e 64 2d 6c 69 6e 65 2d 61 72 67  command-line-arg
	0x0040 75 6d 65 6e 74 73 09 28 64 65 76 65 6c 29 09 0a  uments.(devel)..
	0x0050 f9 32 43 31 86 18 20 72 00 82 42 10 41 16 d8 f2  .2C1.. r..B.A...
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
go.string."a" SRODATA dupok size=1
	0x0000 61                                               a
go.string."xor" SRODATA dupok size=3
	0x0000 78 6f 72                                         xor
runtime.modinfo SDATA size=16
	0x0000 00 00 00 00 00 00 00 00 60 00 00 00 00 00 00 00  ........`.......
	rel 0+8 t=1 go.string."0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tcommand-line-arguments\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"+0
type..importpath.unsafe. SRODATA dupok size=9
	0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·d6c87e1e613370763384ad234357cff1 SRODATA dupok size=10
	0x0000 02 00 00 00 04 00 00 00 00 01