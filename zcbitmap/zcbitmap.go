package zcbitmap

import (
	"fmt"
	"strconv"
)

// --------------------------------------------------------------------------

type BitSet8 uint8

// SetBit 将位图中第i位(从右开始)设置为 1
//  i 一般取值范围: [1, 8], 超过该范围取对8的求余
func (b *BitSet8) SetBit(i uint) {
	mask := uint8(1) << ((i - 1) % 8)
	*b |= BitSet8(mask)
}

// ClearBit 将位图中第i位(从右开始)设置为 0
//  i 一般取值范围: [1, 8], 超过该范围取对8的求余
func (b *BitSet8) ClearBit(i uint) {
	mask := uint8(1) << ((i - 1) % 8)
	*b &= BitSet8(^mask)
}

// CheckBit 检查位图中第 i 位是否为 1
//  i 一般取值范围: [1, 8], 超过该范围取对8的求余
func (b *BitSet8) CheckBit(i uint) bool {
	mask := uint8(1) << ((i - 1) % 8)
	return (*b & BitSet8(mask)) != 0
}

// ToInt 将位图转换为一个整数
func (b *BitSet8) ToInt() uint8 {
	return uint8(*b)
}

// ConvBs8FromUInt8 将一个整数转换为位图
func ConvBs8FromUInt8(i uint8) BitSet8 {
	return BitSet8(i)
}

// ToBinaryStr 将位图转换为一个二进制字符串
func (b *BitSet8) ToBinaryStr(paddingZero bool) string {
	if paddingZero {
		return fmt.Sprintf("%08b", b.ToInt())
	}
	return strconv.FormatUint(uint64(*b), 2)
}

// ConvBs8FromBinaryStr 将一个二进制字符串转换为位图
func ConvBs8FromBinaryStr(s string) BitSet8 {
	if len(s) > 8 {
		s = s[len(s)-8:]
	}
	i, err := strconv.ParseUint(s, 2, 0)
	if err != nil {
		return 0
	}
	return BitSet8(i)
}

// MatchAll 检查当前位图是否完全匹配目标位图的所有非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果仍等于bs则表示满足MatchAll
func (b *BitSet8) MatchAll(bs BitSet8) bool {
	newBs := *b & bs
	return newBs == bs
}

// MatchAny 检查当前位图是否有任意一位匹配目标位图对应的非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果大于0则表示满足MatchAny
func (b *BitSet8) MatchAny(bs BitSet8) bool {
	newBs := *b & bs
	return newBs > 0
}

// --------------------------------------------------------------------------

// BitSet32 32位的位图
type BitSet32 uint32

// SetBit 将位图中第i位(从右开始)设置为 1
//  i 一般取值范围: [1, 32], 超过该范围取对32的求余
func (b *BitSet32) SetBit(i uint) {
	mask := uint32(1) << ((i - 1) % 32)
	*b |= BitSet32(mask)
}

// ClearBit 将位图中第i位(从右开始)设置为 0
//  i 一般取值范围: [1, 32], 超过该范围取对32的求余
func (b *BitSet32) ClearBit(i uint) {
	mask := uint32(1) << ((i - 1) % 32)
	*b &= BitSet32(^mask)
}

// CheckBit 检查位图中第 i 位是否为 1
func (b *BitSet32) CheckBit(i uint) bool {
	mask := uint32(1) << ((i - 1) % 32)
	return (*b & BitSet32(mask)) != 0
}

// ToInt 将位图转换为一个整数
func (b *BitSet32) ToInt() int {
	return int(*b)
}

// ConvBs32FromUInt32 将一个整数转换为位图
func ConvBs32FromUInt32(i uint32) BitSet32 {
	return BitSet32(i)
}

// ToBinaryStr 将位图转换为一个二进制字符串
func (b *BitSet32) ToBinaryStr(paddingZero bool) string {
	if paddingZero {
		return fmt.Sprintf("%032b", b.ToInt())
	}
	return strconv.FormatUint(uint64(*b), 2)
}

// ConvBs32FromBinaryStr 将一个二进制字符串转换为位图
func ConvBs32FromBinaryStr(s string) BitSet32 {
	if len(s) > 32 {
		s = s[len(s)-32:]
	}
	i, err := strconv.ParseUint(s, 2, 0)
	if err != nil {
		return 0
	}
	return BitSet32(i)
}

// MatchAll 检查当前位图是否完全匹配目标位图的所有非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果仍等于bs则表示满足MatchAll
func (b *BitSet32) MatchAll(bs BitSet32) bool {
	newBs := *b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs == bs
}

// MatchAny 检查当前位图是否有任意一位匹配目标位图对应的非零位。
//  所谓匹配，即目标位图第i位非0的话，当前位图的第i位也非零。
//  当前位图与目标位图bs按位与,结果大于0则表示满足MatchAny
func (b *BitSet32) MatchAny(bs BitSet32) bool {
	newBs := *b & bs
	fmt.Printf("newBs: %s\n", newBs.ToBinaryStr(false))
	return newBs > 0
}
