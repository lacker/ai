import Exponent from 'exponent';
import React from 'react';
import {
  Dimensions,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import kenken from './kenken';

class Cell extends React.Component {
  render() {
    if (Math.random() < 0.5) {
      return <View style={{backgroundColor:'#ff0000'}}/>;
    } else {
      return <View style={{backgroundColor:'#0000ff'}}/>;
    }
  }
}

class App extends React.Component {
  constructor(props) {
    super(props);

    this.size = 6;
    let k = kenken(this.size);
    this.puzzle = k.puzzle;
    this.solution = k.solution;
    let answer = [];
    for (let i = 0; i < this.size; i++) {
      answer.push(Array(this.size).fill(null));
    }
    this.state = {
      answer: answer,
    };
  }

  renderRow(i) {
    let cells = [];
    for (let j = 0; j < this.size; j++) {
      cells.push(<Cell key={'cell' + i + '-' + j}/>);
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

    let rows = [];
    for (let i = 0; i < this.size; i++) {
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
    backgroundColor: '#ccc',
    alignItems: 'stretch',
  },
  row: {
    flex: 1,
    flexDirection: 'row',
  },
  cell: {

  },
});

Exponent.registerRootComponent(App);
