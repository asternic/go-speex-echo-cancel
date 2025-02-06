package speexEC

/*
#cgo LDFLAGS: -lspeexdsp
#include <speex/speex_echo.h>
#include <speex/speex_preprocess.h>
#include <stdlib.h>

SpeexEchoState* create_echo_state(int frame_size, int filter_length) {
    return speex_echo_state_init(frame_size, filter_length);
}

void destroy_echo_state(SpeexEchoState* st) {
    speex_echo_state_destroy(st);
}

SpeexPreprocessState* create_preprocess_state(int frame_size) {
    return speex_preprocess_state_init(frame_size, 16000);
}

void destroy_preprocess_state(SpeexPreprocessState* st) {
    speex_preprocess_state_destroy(st);
}

void process_echo(SpeexEchoState* echo_state, short* mic, short* ref, short* output) {
    speex_echo_cancel(echo_state, mic, ref, output, NULL);
}

void preprocess(SpeexPreprocessState* preproc, short* input) {
    speex_preprocess_run(preproc, input);
}
*/
import "C"
import (
	"unsafe"
)

type EchoCanceller struct {
	echoState  *C.SpeexEchoState
	preprocess *C.SpeexPreprocessState
	frameSize  int
}

func NewEchoCanceller(frameSize, filterLen int) *EchoCanceller {
	return &EchoCanceller{
		echoState:  C.create_echo_state(C.int(frameSize), C.int(filterLen)),
		preprocess: C.create_preprocess_state(C.int(frameSize)),
		frameSize:  frameSize,
	}
}

func (ec *EchoCanceller) Process(mic, ref []int16) []int16 {
	output := make([]int16, ec.frameSize)

	C.process_echo(ec.echoState,
		(*C.short)(unsafe.Pointer(&mic[0])),
		(*C.short)(unsafe.Pointer(&ref[0])),
		(*C.short)(unsafe.Pointer(&output[0])),
	)

	C.preprocess(ec.preprocess,
		(*C.short)(unsafe.Pointer(&output[0])),
	)

	return output
}

func (ec *EchoCanceller) Close() {
	C.destroy_echo_state(ec.echoState)
	C.destroy_preprocess_state(ec.preprocess)
}
