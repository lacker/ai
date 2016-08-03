/*
Implement Insert and Delete for a singly-linked linked list
*/

class LinkedList {
  constructor(first, rest) {
    this.first = first;
    this.rest = rest;
  }

  // Returns a new list. Doesn't mutate
  insert(value) {
    return new LinkedList(value, this);
  }

  // Deletes the provided value. If it occurs multiple times, only
  // delete one. Doesn't mutate
  destroy(value) {
    if (this.first == value) {
      return this.rest;
    }
    if (!this.rest) {
      throw new Error('value ' + value + ' not found');
    }
    return new LinkedList(this.first, this.rest.destroy(value));
  }
}
