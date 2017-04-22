#!/usr/bin/env node

const Jimp = require('jimp');

jimp = new Jimp(500, 1500, 0xffff0000, function (error, canvas) {
   console.log(error, canvas);
});
