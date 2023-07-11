package gifmodule

import (
	"bytes"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

type Gif struct {
	Header        *Gif_Header
	LogicalScreen *Gif_LogicalScreen
	_io           *kaitai.Stream
	_root         *Gif
	_parent       interface{}
}

func NewGif() *Gif {
	return &Gif{}
}

func (this *Gif) Read(io *kaitai.Stream, parent interface{}, root *Gif) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := NewGif_Header()
	err = tmp1.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Header = tmp1
	tmp2 := NewGif_LogicalScreen()
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.LogicalScreen = tmp2
	return err
}

type Gif_Header struct {
	Magic   []byte
	Version []byte
	_io     *kaitai.Stream
	_root   *Gif
	_parent *Gif
}

func NewGif_Header() *Gif_Header {
	return &Gif_Header{}
}

func (this *Gif_Header) Read(io *kaitai.Stream, parent *Gif, root *Gif) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadBytes(int(3))
	if err != nil {
		return err
	}
	tmp3 = tmp3
	this.Magic = tmp3
	if !(bytes.Equal(this.Magic, []uint8{71, 73, 70})) {
		return kaitai.NewValidationNotEqualError([]uint8{71, 73, 70}, this.Magic, this._io, "/types/header/seq/0")
	}
	tmp4, err := this._io.ReadBytes(int(3))
	if err != nil {
		return err
	}
	tmp4 = tmp4
	this.Version = tmp4
	return err
}

type Gif_LogicalScreen struct {
	ImageWidth       uint16
	ImageHeight      uint16
	Flags            uint8
	BgColorIndex     uint8
	PixelAspectRatio uint8
	_io              *kaitai.Stream
	_root            *Gif
	_parent          *Gif
}

func NewGif_LogicalScreen() *Gif_LogicalScreen {
	return &Gif_LogicalScreen{}
}

func (this *Gif_LogicalScreen) Read(io *kaitai.Stream, parent *Gif, root *Gif) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp5, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.ImageWidth = uint16(tmp5)
	tmp6, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.ImageHeight = uint16(tmp6)
	tmp7, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Flags = tmp7
	tmp8, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.BgColorIndex = tmp8
	tmp9, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.PixelAspectRatio = tmp9
	return err
}
