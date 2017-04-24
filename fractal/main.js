#!/usr/bin/env node

const Jimp = require('jimp');

jimp = new Jimp(500, 1500, 0x00000000, function (error, image) {
  console.log(error, canvas);
  image.scan(0, 0, image.bitmap.width, image.bitmap.height).then(
    (x, y, idx) => {
      // x, y is the position of this pixel on the image
      // rgba values run from 0 - 255

      let alpha = 255;
      let red = (x + y) & 1 ? 255 : 0; // checkerboard
      let green = 0;
      let blue = 0;

      this.bitmap.data[ idx + 0 ] = red;
      this.bitmap.data[ idx + 1 ] = green;
      this.bitmap.data[ idx + 2 ] = blue;
      this.bitmap.data[ idx + 3 ] = alpha;

      // TODO: write out
    });
});
