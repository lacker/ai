print_string "Starting to reason....\n";

let rec length = fun(alist) => {
  switch alist {
    | [] => 0
    | [x, ...xs] => 1 + length(xs)
  };
};

let rec union = fun(xset, yset) => {
  switch xset {
    | [] => yset
    | [x, ...xs] => switch yset {
      | [] => xset
      | [y, ...ys] => {
        if (x == y) {
          union(xs, yset);
        } else if (x < y) {
          [x, ...union(xs, yset)];
        } else {
          [y, ...union(ys, xset)];
        }
      }
    }
  }
};

let rec map = fun(f, alist) => {
  switch alist {
    | [] => []
    | [x, ...xs] => [f(x), ...map(f, xs)]
  }
};

let rec join = fun(alist : list string, sep : string) => {
  switch alist {
    | [] => ""
    | [x] => x
    | [x, ...rest] => x ^ sep ^ join(rest, sep)
  }
};

let print_list_int = fun(alist) => {
  print_string("[" ^ join(map(string_of_int, alist), ", ") ^ "]\n");
};

let rec mult = fun(k, alist) => {
  switch alist {
    | [] => []
    | [x, ...xs] => [k * x, ...mult(k, xs)]
  }
};

let rec take = fun(n, alist) => {
  if (n <= 0) {
    []
  } else switch alist {
    | [] => []
    | [x, ...xs] => [x, ...take(n - 1, xs)]
  }
};

let rec undupe = fun(alist) => {
  switch alist {
    | [] => []
    | [first, ...pastfirst] => {
      switch pastfirst {
        | [] => [first]
        | [second, ...rest] => {
          if (first == second) {
            undupe(pastfirst);
          } else {
            [first, ...undupe(pastfirst)];
          }
        }
      }
    }
  }
};

let expand = fun(alist) => {
  [1, ...union(union(mult(2, alist), mult(3, alist)), mult(5, alist))];
};

/*
print_list_int(undupe([1, 2, 2, 3, 4, 4, 5]));
*/

print_list_int(expand(expand(expand([1]))));
