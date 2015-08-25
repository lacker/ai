// Overall params
130 => float bpm;
1 :: minute / bpm => dur quarter;
quarter / 2 => dur eighth;

// This loop plays the melody
BeatBox bb;
while (true) {
    bb.chuh(eighth);
    bb.chuh(eighth);
    bb.chuh(eighth);
    bb.khuh(eighth);
}