package handlers

import (
	"fmt"
	"github.com/amanitaverna/go-mp3"
)

type GetAudioInfoResponse struct {
	Duration int64 `json:"duration"`
}

func GetAudioDuration(decoder *mp3.Decoder) (GetAudioInfoResponse, error) {

	bitRate := int64(decoder.SampleRate())
	fmt.Println("bitRate")
	fmt.Println(bitRate)

	size := decoder.Length()

	fmt.Println("size")
	fmt.Println(size)
	duration := size / bitRate

	return GetAudioInfoResponse{Duration: duration}, nil
}
