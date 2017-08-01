package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
)

// KAlpha token for klingon alphabet
type KAlpha uint

// klingon alphabets
const (
	KA     KAlpha = iota //a
	KB                   //b
	KCH                  //ch
	KD                   //D
	KE                   //e
	KGH                  //gh
	KH                   //h
	KI                   //I
	KJ                   //j
	KL                   //l
	KM                   //m
	KN                   //n
	KNG                  //ng
	KO                   //o
	KP                   //p
	KQ                   //q
	KQCap                //Q
	KR                   //r
	KS                   //s
	KT                   //t
	KTLH                 //tlh
	KU                   //u
	KV                   //v
	KW                   //w
	KY                   //y
	KQUOTE               //'
	Space
)

func (k KAlpha) Hex() int {
	switch k {
	case KA:
		return 0xF8D0
	case KB:
		return 0xF8D1
	case KCH:
		return 0xF8D2
	case KD:
		return 0xF8D3
	case KE:
		return 0xF8D4
	case KGH:
		return 0xF8D5
	case KH:
		return 0xF8D6
	case KI:
		return 0xF8D7
	case KJ:
		return 0xF8D8
	case KL:
		return 0xF8D9
	case KM:
		return 0xF8DA
	case KN:
		return 0xF8DB
	case KNG:
		return 0xF8DC
	case KO:
		return 0xF8DD
	case KP:
		return 0xF8DE
	case KQ:
		return 0xF8DF
	case KQCap:
		return 0xF8E0
	case KR:
		return 0xF8E1
	case KS:
		return 0xF8E2
	case KT:
		return 0xF8E3
	case KTLH:
		return 0xF8E4
	case KU:
		return 0xF8E5
	case KV:
		return 0xF8E6
	case KW:
		return 0xF8E7
	case KY:
		return 0xF8E8
	case KQUOTE:
		return 0xF8E9
	case Space:
		return 0x0020
	default:
		return 0
	}
}

var errIgnore = errors.New("ignoring input")

func parse(src string) ([]KAlpha, error) {
	r := strings.NewReader(src)
	var tokens []KAlpha
	for {
		n, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		switch n {
		case 'a', 'A':
			tokens = append(tokens, KA)
		case 'b', 'B':
			tokens = append(tokens, KB)
		case 'c':
			next, _, err := r.ReadRune()
			if err != nil || next != 'h' {
				return nil, errIgnore
			}
			tokens = append(tokens, KCH)
		case 'D', 'd':
			tokens = append(tokens, KD)
		case 'e', 'E':
			tokens = append(tokens, KE)
		case 'g', 'G':
			next, _, err := r.ReadRune()
			if err != nil || next != 'h' {
				return nil, errIgnore
			}
			tokens = append(tokens, KGH)
		case 'H', 'h':
			tokens = append(tokens, KH)
		case 'I', 'i':
			tokens = append(tokens, KI)
		case 'j', 'J':
			tokens = append(tokens, KJ)
		case 'l', 'L':
			tokens = append(tokens, KL)
		case 'm', 'M':
			tokens = append(tokens, KM)
		case 'n':
			next, _, err := r.ReadRune()
			if err != nil {
				r.UnreadRune()
				tokens = append(tokens, KN)
				continue
			}
			if next == 'g' {
				tokens = append(tokens, KNG)
			} else {
				tokens = append(tokens, KN)
			}
		case 'o', 'O':
			tokens = append(tokens, KO)
		case 'p', 'P':
			tokens = append(tokens, KP)
		case 'q':
			tokens = append(tokens, KQ)
		case 'Q':
			tokens = append(tokens, KQCap)
		case 'r', 'R':
			tokens = append(tokens, KR)
		case 'S', 's':
			tokens = append(tokens, KS)
		case 't':
			next, _, err := r.ReadRune()
			if err != nil {
				r.UnreadRune()
				tokens = append(tokens, KT)
				continue
			}
			if next == 'l' {
				next, _, err = r.ReadRune()
				if err != nil || next != 'h' {
					return nil, errIgnore
				}
				tokens = append(tokens, KTLH)
			} else {
				tokens = append(tokens, KT)
			}
		case 'u', 'U':
			tokens = append(tokens, KU)
		case 'v', 'V':
			tokens = append(tokens, KV)
		case 'w', 'W':
			tokens = append(tokens, KW)
		case 'y', 'Y':
			tokens = append(tokens, KY)
		case '\'':
			tokens = append(tokens, KQUOTE)
		case ' ', '\t':
			tokens = append(tokens, Space)
		default:
			return nil, errIgnore
		}
	}
	return tokens, nil
}

func printHex(src string) {
	tks, err := parse(src)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range tks {
		fmt.Printf("0x%X ", v.Hex())
	}
	fmt.Println("")
}

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		x := strings.Join(flag.Args(), " ")
		printHex(x)
	} else {
		flag.PrintDefaults()
	}
}
