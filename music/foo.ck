// connect sine oscillator to D/A converter (sound card)
SinOsc s => dac;

// loop in time
while (true) {
  2::second => now;
}
