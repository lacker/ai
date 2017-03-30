import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);

    this.size = 19;
    let board = [];
    for (let i = 0; i < this.size; ++i) {
      let row = [];
      for (let j = 0; j < this.size; ++j) {
        row.push(0);
      }
      board.push(row);
    }
    this.state = {
      board
    };
  }

  render() {
    return (
      <div>TODO: display the board</div>
    );
  }
}

export default App;
