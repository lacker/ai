// connect sine oscillator to D/A converter (sound card)
SinOsc s => dac;

// loop in time
while (true) {
    100::ms => now;
    Std.rand2f(30.0, 1000.0) => s.freq;
}

