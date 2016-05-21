print_string "Starting to reason....\n";

let merge = fun x y => switch {x, y} {
  | {[], y} => y
  | {x, []} => x
  | {[headX, ...xs], [headY, ...ys]} => {
    if (headX < headY) {
      [headX, ...(merge xs y)]
    } else {
      [headY, ...(merge x ys)]
    }
  }
};

merge [1, 3, 5] [2, 4, 6];