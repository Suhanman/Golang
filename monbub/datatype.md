## Go 데이터타입

- 정수형  
int, int8, int32, int64
uint, uint8, uint32, uint64

- 실수형  
float32, float64
complex64, complex128

- 문자열  
string

- 불린  
bool

- 기타  
byte
rune

문자열은 "" 로 선언  
복수 라인으로 선언할 수 없음  
특수 문자는 이스케이프 문자를 이용해서 처리함  

str := "AA"
str2 := "B\nB"

''로 둘러쌓인 문장은 이스케이프 문자열 해석하지않고 처리
아래의 str2는 출력시 개행하지 않는다.

str := 'AA'
str2 := 'B\nB'