// Please use library.
package main

import (
	"fmt"
	"os"

	"github.com/seastart/go-fdkaac/fdkaac"
)

func main() {

	bytes, e := os.ReadFile("/root/tmp/flv/9.aac")
	if e != nil {
		panic(e)
	}

	var err error
	d := fdkaac.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitRaw(asc); err != nil {
		fmt.Println("init decoder failed, err is", err)
		return
	}
	defer d.Close()

	pcm, err := d.Decode(bytes[2:])
	if err != nil {
		panic(err)
	}
	fmt.Println("SampleRate:", d.SampleRate())
	fmt.Println("FrameSize:", d.FrameSize())
	fmt.Println("NumChannels:", d.NumChannels())
	fmt.Println("AacSampleRate:", d.AacSampleRate())
	fmt.Println("Profile:", d.Profile())
	fmt.Println("AudioObjectType:", d.AudioObjectType())
	fmt.Println("ChannelConfig:", d.ChannelConfig())
	fmt.Println("Bitrate:", d.Bitrate())
	fmt.Println("AacSamplesPerFrame:", d.AacSamplesPerFrame())
	fmt.Println("AacNumChannels:", d.AacNumChannels())
	fmt.Println("ExtensionAudioObjectType:", d.ExtensionAudioObjectType())
	fmt.Println("ExtensionSamplingRate:", d.ExtensionSamplingRate())
	fmt.Println("NumLostAccessUnits:", d.NumLostAccessUnits())
	fmt.Println("NumTotalBytes:", d.NumTotalBytes())
	fmt.Println("NumBadBytes:", d.NumBadBytes())
	fmt.Println("NumTotalAccessUnits:", d.NumTotalAccessUnits())
	fmt.Println("NumBadAccessUnits:", d.NumBadAccessUnits())
	fmt.Println("SampleBits:", d.SampleBits())
	fmt.Println("PCM:", len(pcm))
	err = os.WriteFile("/root/tmp/autio.pcm", pcm, 0644)
	if err != nil {
		panic(err)
	}

	return
}
