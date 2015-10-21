package cwsharp

type Kind int

const (
	PUNCT   Kind = 1 //标号
	NUMERAL Kind = 2 //数字
	LETTER  Kind = 4 //字母
	CJK     Kind = 8 //中文
)

func (k Kind) String() string {
	t := "UNKNOW"
	if k&NUMERAL==NUMERAL && k&LETTER==LETTER{
		t="ALPHANUM"
	}else{
	switch k {
		case PUNCT:
			{
				t = "PUNCT"
			}
		case NUMERAL:
			{
				t = "NUMERAL"
			}
		case LETTER:
			{
				t = "LETTER"
			}
		case CJK:
			{
				t = "CJK"
			}
		}
	}

	return t
}
