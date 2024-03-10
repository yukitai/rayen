package transpiler

var chars []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type idgen struct {
	curr rune
	head *idgen
	len  int
	end  bool
}

type id struct {
	gen idgen
}

var Id = id{
	gen: idgen{
		curr: 0,
		head: nil,
		len:  1,
		end:  false,
	},
}

func (i *id) Next() string {
	if i.gen.end {
		expanded := idgen{
			curr: 0,
			head: nil,
			len: i.gen.len + 1,
			end: false,
		}
		i.gen = expanded
	}
	return string(i.gen.Next())
}

func (i *idgen) Next() []rune {
	if i.len == 1 {
		ret := chars[i.curr]
		i.curr += 1
		if int(i.curr) >= len(chars) {
			i.end = true
		}
		return []rune{ret}
	}
	if i.head == nil {
		head := idgen{
			curr: 0,
			head: nil,
			len:  i.len - 1,
			end:  false,
		}
		i.head = &head
	}
	ret := append(i.head.Next(), chars[i.curr])
	if i.head.end {
		i.head = nil
		i.curr += 1
		if int(i.curr) >= len(chars) {
			i.end = true
		}
	}
	return ret
}