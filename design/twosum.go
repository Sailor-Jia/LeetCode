package design

type dataStruct struct {
	number map[int][]int
}

func NewData() dataStruct{
	return dataStruct{
		number: make(map[int][]int,100),
	}
}
func (s *dataStruct) Add(value int){
	if val,ok:=s.number[value];ok{
		val=append(val,len(s.number))
		s.number[value]=val
	}else {
		s.number[value]=[]int{len(s.number)}
	}
}

func (s *dataStruct) Find(target int) bool{
	for num,_:=range s.number{
		if i,ok:=s.number[target-num];ok{
			if (num==target-num && len(i)>1) ||num!=target-num{
				return true
			}
		}
	}
	return false
}


