package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// byte is a uint8
type WaveChuck struct {
	ChuckID       [4]byte // "RIFF"
	ChuckDataSize uint32
	RiffType      [4]byte // "WAVE"
}
type FormatChuck struct {
	ChuckID       [4]byte // "fmt "
	ChuckDataSize uint32
	FormatTag     uint16
	NumOfChannels uint16
	SampleRate    uint32
	BytePerSec    uint32
	BlockAlign    uint16
	BitDepth      uint16
	//	ExtraFormatByte uint16
}
type SoundDataChuck struct {
	ChuckID       [4]byte
	SoundDataSize []uint32 // not sure
}

type Wav struct {
	WaveChuck
	FormatChuck
	SoundDataChuck
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("wav file to parse???")
		os.Exit(1)
	}
	filename := os.Args[1]
	f, _ := os.Open(filename)
	Q := Wav{}
	binary.Read(f, binary.BigEndian, &Q.WaveChuck.ChuckID)
	binary.Read(f, binary.LittleEndian, &Q.WaveChuck.ChuckDataSize)
	binary.Read(f, binary.BigEndian, &Q.WaveChuck.RiffType)
	binary.Read(f, binary.BigEndian, &Q.FormatChuck.ChuckID)
	binary.Read(f, binary.LittleEndian, &Q.FormatChuck.ChuckDataSize)
	binary.Read(f, binary.LittleEndian, &Q.FormatChuck.FormatTag)
	binary.Read(f, binary.LittleEndian, &Q.FormatChuck.NumOfChannels)
	binary.Read(f, binary.LittleEndian, &Q.FormatChuck.SampleRate)
	binary.Read(f, binary.LittleEndian, &Q.FormatChuck.BytePerSec)
	binary.Read(f, binary.LittleEndian, &Q.FormatChuck.BlockAlign)
	binary.Read(f, binary.LittleEndian, &Q.FormatChuck.BitDepth)
	//binary.Read(f, binary.LittleEndian, &Q.FormatChuck.ExtraFormatByte)
	binary.Read(f, binary.BigEndian, &Q.SoundDataChuck.ChuckID)
	riff := ""
	for i := range Q.WaveChuck.ChuckID {
		riff += string(Q.WaveChuck.ChuckID[i])
	}
	wave := ""
	for i := range Q.WaveChuck.RiffType {
		wave += string(Q.WaveChuck.RiffType[i])
	}
	wfmt := ""
	for i := range Q.FormatChuck.ChuckID {
		wfmt += string(Q.FormatChuck.ChuckID[i])
	}

	data := ""
	for i := range Q.SoundDataChuck.ChuckID {
		data += string(Q.SoundDataChuck.ChuckID[i])
	}
	fmt.Printf("RIFF:\t\t%s\n", riff)

	fmt.Printf("WAVE:\t\t\t%s\n", wave)
	fmt.Printf("WFMT:\t\t\t%s\n", wfmt)
	fmt.Printf("DATA:\t\t\t%s\n", data)
	fmt.Printf("WaveChuckDataSize:\t%d\n", Q.WaveChuck.ChuckDataSize)
	fmt.Printf("BitDepth:\t\t%d\n", Q.FormatChuck.BitDepth)
	fmt.Printf("Sample Rate:\t%d\n", Q.FormatChuck.SampleRate)

	return
}
