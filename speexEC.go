package speexEC

/*
#cgo LDFLAGS: -lspeexdsp
#include <speex/speex_echo.h>
#include <stdlib.h>

typedef struct {
    SpeexEchoState *echo_state;
} EchoCanceller;

EchoCanceller* speex_aec_init(int frame_size, int filter_length, int sample_rate) {
    EchoCanceller* ec = (EchoCanceller*)malloc(sizeof(EchoCanceller));
    ec->echo_state = speex_echo_state_init(frame_size, filter_length); // Fixed typo: frames_size -> frame_size
    speex_echo_ctl(ec->echo_state, SPEEX_ECHO_SET_SAMPLING_RATE, &sample_rate);
    return ec;
}

void speex_aec_process(EchoCanceller *ec, short *mic, short *ref, short *out) {
    spx_int32_t Yout; // Add this variable for the output energy
    speex_echo_cancel(ec->echo_state, mic, ref, out, &Yout); // Added Yout as the 5th argument
}

void speex_aec_destroy(EchoCanceller *ec) {
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

// Creates new AEC module
func NewEchoCanceller(frameSize, filterLen int, sampleRate int) *EchoCanceller {
    ec := C.speex_aec_init(C.int(frameSize), C.int(filterLen), C.int(sampleRate))
    return &EchoCanceller{echoState: ec}
}

// Processes audio frame with AEC
func (ec *EchoCanceller) Process(mic, ref []int16) []int16 {
    frameSize := len(mic)
    out := make([]int16, frameSize)

    C.speex_aec_process(
        ec.echoState,
        (*C.short)(unsafe.Pointer(&mic[0])),
        (*C.short)(unsafe.Pointer(&ref[0])),
        (*C.short)(unsafe.Pointer(&out[0])),
    )
    return out
}

// Destroys AEC module
func (ec *EchoCanceller) Destroy() {
    C.speex_aec_destroy(ec.echoState)
}
