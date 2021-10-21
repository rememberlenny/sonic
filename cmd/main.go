package cmd

/*
 #include <stdio.h>
 #include <stdlib.h>

 static void myprint(char* s) {
 	printf("%s\n", s);
 }

long charAt(const char *sp, size_t nb, char c)
{
    long ret = -1;
    for (long i = 0; i < nb; i++)
    {
        if (sp[i] == c)
        {
            ret = i;
            break;
        }
    }
    return ret;
}

*/
import "C"
import (
	"strconv"

	"github.com/bytedance/sonic/internal/native/avx2"
)

func GenS(n int) string {
	if n == 1 {
		return "]"
	}
	var s = "["
	for i := 0; i < n; i++ {
		if i > 0 {
			s += ","
		}
		s += strconv.Itoa(i)
	}
	s += "]"
	return s
}

func RunGo(s string) int {
	ret := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			ret = i
			break
		}
	}
	return ret
}

func RunC(s string) {
	cs := C.CString(s)
	nb := C.ulong(len(s))
	_ = C.charAt(cs, nb, ']')
}

func RunNative(s string) {
	_ = avx2.CharAt(s, ']')
}
