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
	SoundDataSize uint32 // not sure
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
	Q := new(Wav)
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
	binary.Read(f, binary.BigEndian, &Q.SoundDataChuck.ChuckID)
	binary.Read(f, binary.LittleEndian, &Q.SoundDataChuck.SoundDataSize)
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
	fmt.Printf("RIFF:\t\t\t\t%s\n", riff)
	fmt.Printf("WAVE:\t\t\t\t%s\n", wave)
	fmt.Printf("WFMT:\t\t\t\t%s\n", wfmt)
	fmt.Printf("DATA:\t\t\t\t%s\n", data)
	fmt.Printf("WaveChuckDataSize:\t\t%d\n", Q.WaveChuck.ChuckDataSize)
	fmt.Printf("SoundDataChuckDataSize:\t\t%d\n", Q.SoundDataChuck.SoundDataSize)
	fmt.Printf("BitDepth:\t\t\t%d\n", Q.FormatChuck.BitDepth)
	fmt.Printf("Sample Rate:\t\t\t%d\n", Q.FormatChuck.SampleRate)
	fmt.Printf("WaveChuckDataSize:\t\t%d\n", Q.WaveChuck.ChuckDataSize)
	fmt.Printf("FormatChuckDataSize:\t\t%d\n", Q.FormatChuck.ChuckDataSize)
	fmt.Printf("FormatTag:\t\t\t%d\n", Q.FormatChuck.FormatTag)
	fmt.Printf("NumOfChannels:\t\t\t%d\n", Q.FormatChuck.NumOfChannels)
	fmt.Printf("SampleRate:\t\t\t%d\n", Q.FormatChuck.SampleRate)
	fmt.Printf("BytePerSec:\t\t\t%d\n", Q.FormatChuck.BytePerSec)
	fmt.Printf("BlockAlign:\t\t\t%d\n", Q.FormatChuck.BlockAlign)
	fmt.Printf("BitDepth:\t\t\t%d\n", Q.FormatChuck.BitDepth)
	fmt.Printf("SoundDataChuckSize:\t\t%d\n", Q.SoundDataChuck.SoundDataSize)
	fmt.Printf("%d 個 sample\n", Q.SoundDataChuck.SoundDataSize/(uint32(Q.FormatChuck.NumOfChannels)*uint32(Q.FormatChuck.BitDepth/8)))

	return
}
