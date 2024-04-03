package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const BinaryType uint8 = iota + 1
// constanta , uint8 ->unsign integer 8 bits , iota-> dimulai dari 0 (karena 0+1 jd dimulai dari 1)

type Payload interface {
// interface itu tipe data tp dipakai contract method
// kl mw jd payload harus memenuhi 3 syarat di bwh itu
	io.WriterTo
	io.ReaderFrom
	Bytes() Binary  //[]byte --> Binary
	// method bytes tp returnnya binary
}

type Binary []byte
// nama var binary tp data typenya byte

// buat return binarnya
// kalo di kasih kata hallo akan jwb hallo
func (m Binary) Bytes() Binary{
	return m
}


func(m Binary) WriteTo(w io.Writer)(int64,error){

	//bakal di write dalam bentuk apa? dan size nya 1
	err := binary.Write(w, binary.BigEndian,BinaryType)
// buat ubah messagenya jdi format Binary

	if err != nil{
		fmt.Println(err)
		return 0,err
	}

	// bakal write sebanyak apa? dan menghasilkan size 4
	err = binary.Write(w, binary.BigEndian,int32(len(m)))

	if err != nil{
		fmt.Println(err)
		return 0,err
	}

		
	n,err := w.Write(m)
	return int64(n+5),err
}


func (m *Binary) ReadFrom(r io.Reader)(int64,error){

	var typ uint8
	err:= binary.Read(r,binary.BigEndian,&typ)

	if err != nil{
		fmt.Println(err)
		return 0,err
	}

	var size int32
	err = binary.Read(r,binary.BigEndian,&size)
	if err != nil{
		fmt.Println(err)
		return 0,err
	}

	*m = make(Binary, size)
	n,err1 := r.Read(*m)

	return int64(n + 5),err1
}

func Decode(r io.Reader)(Payload, error){

	var typ uint8
	err:= binary.Read(r,binary.BigEndian,&typ)

	if err != nil{
		fmt.Println(err)
		return nil,err
	}

	payload := new(Binary)

	_,err = payload.ReadFrom(io.MultiReader(bytes.NewReader(Binary{typ}),r))

	if err != nil{
		fmt.Println(err)
		return nil,err
	}

	return payload,err


}