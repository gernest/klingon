package main

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
	default:
		return 0
	}
}

func main() {

}
