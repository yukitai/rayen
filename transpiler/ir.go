package transpiler

type Blocks struct {
	Blocks map[string]Block `json:"blocks"`
}

func NewBlocks() Blocks {
	return Blocks{
		Blocks: make(map[string]Block),
	}
}

type Block struct {
	Opcode    string           `json:"opcode"`
	Next     *string           `json:"next"`
	Parent   *string           `json:"parent"`
	Inputs    map[string]Input `json:"inputs"`
	Fields    map[string]Field `json:"fields"`
	Shadow    bool             `json:"shadow"`
	TopLevel  bool             `json:"topLevel"`
	Xloc      int              `json:"x"`
	Yloc      int              `json:"y"`
	Mutation *Mutation         `json:"mutation"`
}

func NewBlock() Block {
	return Block{
		Inputs: make(map[string]Input),
		Fields: make(map[string]Field),
	}
}

type InputType int

const (
	Number    InputType = 1
	Angle     InputType = 8
	Color     InputType = 9
	String    InputType = 10
	Broadcast InputType = 11
	Variable  InputType = 12
	List      InputType = 13
)

type InputValue interface {}

type NumberInput    float64
type AngleInput     float64
type ColorInput     string
type StringInput    string

type BroadcastInput []string

func NewBroadcastInput(value, id string) BroadcastInput {
	return BroadcastInput{
		value,
		id,
	}
}

type VariableInput []string

func NewVariableInput(value, id string) VariableInput {
	return VariableInput{
		value,
		id,
	}
}

type ListInput []string

func NewListInput(value, id string) ListInput {
	return ListInput{
		value,
		id,
	}
}

type ShadowType int

const (
	Shadow   ShadowType = 1
	NoShadow ShadowType = 2
)

type Input []interface{}

func NewInput(shadow ShadowType, value InputValue) Input {
	return Input{
		shadow,
		value,
	}
}

type Field []string

func NewField(value, id string) Field {
	return Field{
		value,
		id,
	}
}

type Mutation interface {}

type BasicMutation struct {
	TagName     string     `json:"tagName"`
	Children    []struct{} `json:"children"`
	Proccode    string     `json:"proccode"`
	Argumentids []string   `json:"argumentids"`
	Warp        bool       `json:"warp"`
}

func (args *Arguments) NewBasicMutation(proccode string, warp bool) BasicMutation {
	return BasicMutation{
		TagName: "mutation",
		Children: make([]struct{}, 0),
		Proccode: proccode,
		Argumentids: args.Argumentids,
		Warp: warp,
	}
}

type PrototypeMutation struct {
	TagName          string     `json:"tagName"`
	Children         []struct{} `json:"children"`
	Proccode         string     `json:"proccode"`
	Argumentids      []string   `json:"argumentids"`
	Warp             bool       `json:"warp"`
	Argumentnames    []string   `json:"argumentnames"`
	Argumentdefaults []string   `json:"argumentdefaults"`
}

func (args *Arguments) NewPrototypeMutation() PrototypeMutation {
	return PrototypeMutation{
		TagName: "mutation",
		Children: make([]struct{}, 0),
		Argumentids: args.Argumentids,
		Argumentnames: args.Argumentnames,
		Argumentdefaults: make([]string, args.Count),
	}
}

type StopBlockMutation struct {
	TagName     string     `json:"tagName"`
	Children    []struct{} `json:"children"`
	Hasnext     bool       `json:"hasnext"`
}

func NewStopBlockMutation(hasNext bool) StopBlockMutation {
	return StopBlockMutation{
		TagName: "mutation",
		Children: make([]struct{}, 0),
		Hasnext: hasNext,
	}
}

type Arguments struct {
	Count int
	Argumentids []string
	Argumentnames []string

}

