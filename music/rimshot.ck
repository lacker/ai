// A cymbaly sound
public class RimShot {
    Impulse i => LPF f1 => DelayA body => Gain ampenv => SinOsc drive => Gain output;
    body => Gain body_fb => body;
    i => Gain g1 => Gain g1_fb => g1 => ampenv;
    Noise n => LPF n_f => Gain n_ampenv => output;
    i => Gain g2 => Gain g2_fb => g2 => LPF g2_f => n_ampenv;   
   
    3 => ampenv.op => n_ampenv.op;
    1 => drive.sync;
   
    f1.set(1000, 1);
    second => body.max;
    second / 600 => body.delay;
    .9 => body_fb.gain;   
   
    1.0 - 1.0/800 => g1_fb.gain;
   
    1.2 => ampenv.gain;
    0.3 => n_ampenv.gain;
   
    n_f.set(5000, 1.5);
    1.0 - 1.0 / 1000 => g2_fb.gain;
    g2_f.set(50, 1);
   
    fun void hit(float v) {
        v => i.next;
    }

    fun void rhythm(dur quarter) {
        quarter / 2 => dur eighth;
        eighth / 2 => dur sixteenth;

        hit(2);
        quarter => now;
        hit(1);
        eighth => now;
        hit(1);
        sixteenth => now;
        hit(1);
        sixteenth => now;
    }
}

RimShot r;
r.output => dac;

for (int i; i < 100; i++) {
    r.rhythm(1 :: minute / 260);
}
