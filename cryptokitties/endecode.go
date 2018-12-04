// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptokitties

import (
	"math/big"
	// "strings"
	// "strconv"
	"math/rand"
	// ethereum "github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/event"
	// "fmt"
)
var (
	// ALPHABET = "123456789abcdefghijkmnopqrstuvwx"//此实现不是32进制表示，正确的应该是从0-w，在golang中需要将转换完的32进制
	// 			   0123456789abcdefghijklmnopqrstuv
	// BASE     = 32   //## 32 chars/letters/digits
	
)
func ConvertTo32(hex string)(ret string){
	m1 := make(map[string]string)
	m1["0"] = "1"
	m1["1"] = "2"
	m1["2"] = "3"
	m1["3"] = "4"
	m1["4"] = "5"
	m1["5"] = "6"
	m1["6"] = "7"
	m1["7"] = "8"
	m1["8"] = "9"
	m1["9"] = "a"
	m1["a"] = "b"
	m1["b"] = "c"
	m1["c"] = "d"
	m1["d"] = "e"
	m1["e"] = "f"
	m1["f"] = "g"
	m1["g"] = "h"
	m1["h"] = "i"
	m1["i"] = "j"
	m1["j"] = "k"
	m1["k"] = "m"
	m1["l"] = "n"
	m1["m"] = "o"
	m1["n"] = "p"
	m1["o"] = "q"
	m1["p"] = "r"
	m1["q"] = "s"
	m1["r"] = "t"
	m1["s"] = "u"
	m1["t"] = "v"
	m1["u"] = "w"
	m1["v"] = "x"

	g1:=big.NewInt(0)
	g1.SetString(hex,16)
	// fmt.Println(g1.Text(16))
	// fmt.Println(g1.Text(32))
	ori:=g1.Text(32)
	for _,v:=range ori{
		ret+=m1[string(v)]
	}
	return
}
func iseven(c rune)bool{
	switch c {
		case '2'://123456789abcdefghijkmnopqrstuvwx
		case '4':
		case '6':
		case '8':
		case 'a':
		case 'c':
		case 'e':
		case 'g':
		case 'i':
		case 'k':
		case 'n':
		case 'p':
		case 'r':
		case 't':
		case 'v':
		case 'x':			
			return true
		default:
			// freebsd, openbsd,
			// plan9, windows...
			// fmt.Printf("%s.", os)
			return false
	}
}
func MixGenes(mGenes,sGenes string)(babyGenes string){
	for i:=0;i<12;i++{
		index:=4*i
		for j:=3;j>0;j--{
			if rand.Float64()<0.25{
				chars := []rune(mGenes)
				chars[index+j], chars[index+j-1] = chars[index+j-1], chars[index+j]
				mGenes=string(chars)
			}
			if rand.Float64()<0.25{
				chars := []rune(sGenes)
				chars[index+j], chars[index+j-1] = chars[index+j-1], chars[index+j]
				sGenes=string(chars)
			}
		}
	}
	for i:=0;i<48;i++{
		mutation := 0
		if i%4==0{
			gene1 := mGene[i]
			gene2 := sGene[i]
			if gene1 > gene2{
				gene1, gene2 = gene2, gene1
			}
			
			if (gene2 - gene1) == 1 && iseven(gene1){
				probability := 0.25
				if gene1 > 23{
					probability /= 2
				}
				
				if rand.Float64() < probability{
					mutation = (gene1 / 2) + 16
				}	
			}
		}
		if mutation>0{
			babyGenes[i] = mutation
		}else{
			if rand.Float64() < 0.5{
				babyGenes[i] = mGene[i]
			}else{
				babyGenes[i] = sGene[i]
			}
		}
	} 
}
