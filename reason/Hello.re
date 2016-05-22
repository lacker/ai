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

print_int(length([6, 7, 8]));