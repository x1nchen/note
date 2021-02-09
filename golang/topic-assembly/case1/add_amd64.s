TEXT ·add(SB), $0-24 // add栈空间为0，入参+返回值大小=24字节
MOVQ x+0(FP), AX // 从main中取参数：2
ADDQ y+8(FP), AX // 从main中取参数：3
MOVQ AX, ret+16(FP) // 保存结果到返回值
RET
// 最后一行的空行是必须的，否则可能报 unexpected EOF