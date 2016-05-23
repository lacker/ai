print_string "Starting to reason....\n";

let rec length = fun(alist) => {
  switch alist {
    | [] => 0
    | [x, ...xs] => 1 + length(xs)
  };
};

let isempty = fun(alist) => {
  switch alist {
    | [] => true
    | [x, ...xs] => false
  }
};

let rec merge = fun(xlist, ylist) => {
  if (isempty(xlist)) {
    ylist;
  } else if (isempty(ylist)) {
    xlist;
  } else {
    let [headX, ...restX] = xlist;
    let [headY, ...restY] = ylist;
    if (headX < headY) {
      [headX, ...merge(restX, ylist)];
    } else {
      [headY, ...merge(restY, xlist)];
    }
  }
};

print_int(length(merge([1, 2], [3, 4])));