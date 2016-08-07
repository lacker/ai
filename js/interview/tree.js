/*
Insert into a binary tree
*/

class Tree {
  constructor(value, left, right) {
    this.value = value;
    this.left = left;
    this.right = right;
  }
}

// Does not mutate
function insert(tree, value) {
  if (!tree) {
    return new Tree(value);
  }
  if (value < tree.value) {
    return new Tree(tree.value, insert(tree.left, value), tree.right);
  }
  return new Tree(tree.value, tree.left, insert(tree.right, value));
}
