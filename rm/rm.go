// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package rm

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


type Rm_Color int
const (
	Rm_Color__Black Rm_Color = 0
	Rm_Color__Gray Rm_Color = 1
	Rm_Color__White Rm_Color = 2
)

type Rm_Brush int
const (
	Rm_Brush__Eraser Rm_Brush = 6
	Rm_Brush__Eraserselect Rm_Brush = 8
	Rm_Brush__Paintbrush Rm_Brush = 12
	Rm_Brush__Mechanical Rm_Brush = 13
	Rm_Brush__Pencil Rm_Brush = 14
	Rm_Brush__Ballpoint Rm_Brush = 15
	Rm_Brush__Marker Rm_Brush = 16
	Rm_Brush__Fineliner Rm_Brush = 17
	Rm_Brush__Highlighter Rm_Brush = 18
)
type Rm struct {
	Header *Rm_Header
	NumLayers int32
	Layers []*Rm_Layer
	_io *kaitai.Stream
	_root *Rm
	_parent interface{}
}

func (this *Rm) Read(io *kaitai.Stream, parent interface{}, root *Rm) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := new(Rm_Header)
	err = tmp1.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Header = tmp1
	tmp2, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.NumLayers = int32(tmp2)
	this.Layers = make([]*Rm_Layer, this.NumLayers)
	for i := range this.Layers {
		tmp3 := new(Rm_Layer)
		err = tmp3.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Layers[i] = tmp3
	}
	return err
}
type Rm_Header struct {
	Magic []byte
	Version string
	Spaces []byte
	_io *kaitai.Stream
	_root *Rm
	_parent *Rm
}

func (this *Rm_Header) Read(io *kaitai.Stream, parent *Rm, root *Rm) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp4, err := this._io.ReadBytes(int(32))
	if err != nil {
		return err
	}
	tmp4 = tmp4
	this.Magic = tmp4
	tmp5, err := this._io.ReadBytes(int(1))
	if err != nil {
		return err
	}
	tmp5 = tmp5
	this.Version = string(tmp5)
	tmp6, err := this._io.ReadBytes(int(10))
	if err != nil {
		return err
	}
	tmp6 = tmp6
	this.Spaces = tmp6
	return err
}
type Rm_Point struct {
	X float32
	Y float32
	Speed float32
	Direction float32
	Width float32
	Pressure float32
	_io *kaitai.Stream
	_root *Rm
	_parent *Rm_Line
}

func (this *Rm_Point) Read(io *kaitai.Stream, parent *Rm_Line, root *Rm) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp7, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.X = float32(tmp7)
	tmp8, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Y = float32(tmp8)
	tmp9, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Speed = float32(tmp9)
	tmp10, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Direction = float32(tmp10)
	tmp11, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Width = float32(tmp11)
	tmp12, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Pressure = float32(tmp12)
	return err
}
type Rm_Line struct {
	BrushType Rm_Brush
	Color Rm_Color
	Padding []byte
	BrushSize float32
	Unknown int32
	NumPoints int32
	Points []*Rm_Point
	_io *kaitai.Stream
	_root *Rm
	_parent *Rm_Layer
}

func (this *Rm_Line) Read(io *kaitai.Stream, parent *Rm_Layer, root *Rm) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp13, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.BrushType = Rm_Brush(tmp13)
	tmp14, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.Color = Rm_Color(tmp14)
	tmp15, err := this._io.ReadBytes(int(4))
	if err != nil {
		return err
	}
	tmp15 = tmp15
	this.Padding = tmp15
	tmp16, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.BrushSize = float32(tmp16)
	tmp17, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.Unknown = int32(tmp17)
	tmp18, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.NumPoints = int32(tmp18)
	this.Points = make([]*Rm_Point, this.NumPoints)
	for i := range this.Points {
		tmp19 := new(Rm_Point)
		err = tmp19.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Points[i] = tmp19
	}
	return err
}
type Rm_Layer struct {
	NumLines int32
	Lines []*Rm_Line
	_io *kaitai.Stream
	_root *Rm
	_parent *Rm
}

func (this *Rm_Layer) Read(io *kaitai.Stream, parent *Rm, root *Rm) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp20, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.NumLines = int32(tmp20)
	this.Lines = make([]*Rm_Line, this.NumLines)
	for i := range this.Lines {
		tmp21 := new(Rm_Line)
		err = tmp21.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Lines[i] = tmp21
	}
	return err
}
