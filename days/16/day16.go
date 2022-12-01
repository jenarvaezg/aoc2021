package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"aoc2021/utils/intMath"
	"fmt"
	"strconv"
	"strings"
)

type packet struct {
	version    int
	id         int
	value      int
	subPackets []packet
}

func hexToBin(input string) string {
	var sb strings.Builder
	for _, c := range input {
		ui, err := strconv.ParseUint(string(c), 16, 64)
		if err != nil {
			panic(err)
		}
		sb.WriteString(fmt.Sprintf("%04b", ui))
	}

	return sb.String()

}

func (p packet) versionSum() int {
	total := p.version
	for _, subP := range p.subPackets {
		total += subP.versionSum()
	}

	return total

}

func (p packet) Value() int {
	values := make([]int, len(p.subPackets))
	for i, subP := range p.subPackets {
		values[i] = subP.Value()
	}
	switch p.id {
	case 0:
		return intMath.IntSum(values...)
	case 1:
		return intMath.IntProduct(values...)
	case 2:
		return intMath.IntMin(values...)
	case 3:
		return intMath.IntMax(values...)
	case 4:
		return p.value
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0

	default:
		panic(fmt.Sprintf("%v", p))
	}
}

func parsePacket(input string) (packet, int) {
	version := conversions.MustAtobin(input[:3])
	id := conversions.MustAtobin(input[3:6])
	pointer := 6
	var value int
	var subPackets []packet
	if id == 4 { //literal
		var valueBuilder strings.Builder
		for i := pointer; ; i += 5 {
			// todo multiple of 4 shit ?
			pointer += 5
			valueBuilder.WriteString(input[i+1 : i+5])
			if input[i] == '0' {
				break
			}
		}
		value = conversions.MustAtobin(valueBuilder.String())
	} else { //operator
		pointer += 1
		if input[6] == '0' { // bit length
			length := conversions.MustAtobin(input[7:22])
			for l := 0; l < length; {
				subPacket, read := parsePacket(input[22+l:])
				l += read
				subPackets = append(subPackets, subPacket)
			}
			pointer += length + 15
		} else { //subPackets count
			n := conversions.MustAtobin(input[7:18])
			var nestedRead int
			for i := 0; i < n; i++ {
				subPacket, read := parsePacket(input[18+nestedRead:])
				nestedRead += read
				subPackets = append(subPackets, subPacket)
			}
			pointer += nestedRead + 11
		}

	}

	return packet{version: version, id: id, subPackets: subPackets, value: value}, pointer
}

func main() {
	puzzleInput := files.ReadInput()
	binInput := hexToBin(puzzleInput)

	packet, _ := parsePacket(binInput)

	fmt.Println(packet.versionSum())
	fmt.Println(packet.Value())

}
