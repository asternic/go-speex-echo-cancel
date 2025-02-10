package speexEC

/*
#cgo LDFLAGS: -lspeexdsp
#include <speex/speex_echo.h>
#include <stdlib.h>

typedef struct {
	SpeexEchoState *echo_state;
} EchoCanceller;

EchoCanceller* speex_aec_init(int frame_size, int filter_length, int sample_rate){
	EchoCanceller* ec = (EchoCanceller*)malloc(sizeof(EchoCanceller));
	ec->echo_state = speex_echo_state_init(frames_size, filter_length);
	speex_echo_ctl(ec->echo_state, SPEEX_ECHO_SET_SAMPLING_RATE, &sample_rate);
	return ec;
}

void speex_aec_process(EchoCanceller *ec, short *mic, short *ref, short *out){
	speex_echo_cancel(ec->echo_state, mic, ref, out)
}

SpeexEchoState* create_echo_state(int frame_size, int filter_length) {
    return speex_echo_state_init(frame_size, filter_length);
}

void speex_aec_destroy(EchoCanceller *ec){
	speex_echo_state_destroy(ec->echo_state);
	free(ec);
}
*/
import "C"
import (
	"unsafe"
)

type EchoCanceller struct {
	echoState *C.EchoCanceller
}

func NewEchoCanceller(frameSize, filterLen int, sample_rate int) *EchoCanceller {

	ec := C.speex_aec_init(C.int(frameSize), C.int(filterLen), C.int(sample_rate))
	return &EchoCanceller{echoState: ec}

}

func (ec *EchoCanceller) Process(mic, ref []int16) []int16 {
	frameSize := len(mic)
	out := make([]int16, frameSize)

	C.speex_aec_process(
		ec.state,
		(*C.short)(unsafe.Pointer(&mic[0])),
		(*C.short)(unsafe.Pointer(&ref[0])),
		(*C.short)(unsafe.Pointer(&out[0])),
		C.int(frameSize),
	)
	return out
}

func (ec *EchoCanceller) Destroy() {
	C.speex_aec_destroy(ec.state)
}
