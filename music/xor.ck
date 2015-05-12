// connect sine oscillator to D/A converter (sound card)
SinOsc s => dac;

// loop in time
0 => int n;
while (true) {
    1 + n => n;

    // find the xor of n
    0 => int xor;
    n => int ncopy;
    while (ncopy > 0) {
        xor ^ (ncopy & 1) => xor;
        ncopy >> 1 => ncopy;
    }

    // A and B
    440 + 53.88 * xor => s.freq;

    100::ms => now;
}

