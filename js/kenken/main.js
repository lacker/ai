import Exponent from 'exponent';
import React from 'react';
import {
  Dimensions,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import kenken from './kenken';

const SIZE = 6;

class Cell extends React.Component {
  cageBorder(delta) {
    return (
      this.props.cageForIndex[this.props.index] !==
      this.props.cageForIndex[this.props.index + delta]);
  }

  render() {
    // Customizing the style
    let cageNum = this.props.cageForIndex[this.props.index];
    let custom = {};
    let border = '#000';

    if (this.props.index < SIZE || this.cageBorder(-SIZE)) {
      custom.borderTopColor = border;
    }
    if (this.props.index >= SIZE * SIZE - SIZE || this.cageBorder(SIZE)) {
      custom.borderBottomColor = border;
    }
    if (this.props.index % SIZE === 0 || this.cageBorder(-1)) {
      custom.borderLeftColor = border;
    }
    if (this.props.index % SIZE === SIZE - 1 || this.cageBorder(1)) {
      custom.borderRightColor = border;
    }

    return (
      <View style={[styles.cell, custom]} />
    );
  }
}

class App extends React.Component {
  constructor(props) {
    super(props);

    let k = kenken(SIZE);
    this.puzzle = k.puzzle;
    this.solution = k.solution;
    let answer = [];
    for (let i = 0; i < SIZE; i++) {
      answer.push(Array(SIZE).fill(null));
    }
    this.state = {
      answer: answer,
    };
  }

  renderRow(i) {
    let cells = [];
    for (let j = 0; j < SIZE; j++) {
      let index = SIZE * i + j;
      cells.push(
        <Cell
          key={'cell' + index}
          index={index}
          cageForIndex={this.puzzle.cageForIndex}
          />
      );
    }
    return (
      <View style={styles.row} key={'row' + i}>
        {cells}
      </View>
    );
  }

  render() {
    let {height, width} = Dimensions.get('window');
    let dim = Math.min(height, width);
    while (dim % 6 !== 2) {
      dim--;
    }

    let rows = [];
    for (let i = 0; i < SIZE; i++) {
      rows.push(this.renderRow(i));
    }
    return (
      <View style={styles.container}>
        <View style={{flex: 1}} />
        <View style={[styles.board, {height: dim, width: dim}]}>
          {rows}
        </View>
        <View style={{flex: 1}} />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
  board: {
    backgroundColor: '#fff',
    alignItems: 'stretch',
    borderWidth: 1,
    borderColor: '#000',
  },
  row: {
    flex: 1,
    flexDirection: 'row',
  },
  cell: {
    flex: 1,
    borderWidth: 1,
    borderColor: '#fafafa',
  },
});

Exponent.registerRootComponent(App);
