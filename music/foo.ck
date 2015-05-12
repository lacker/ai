#!/usr/local/bin/chuck

// connect sine oscillator to D/A convertor (sound card)
SinOsc s => dac;

// loop in time
while (true) {
  2::second => now;
}
