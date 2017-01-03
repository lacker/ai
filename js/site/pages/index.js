import React from 'react';
import { observable } from 'mobx';
import 'isomorphic-fetch';

import ListView from '../ListView';

class App extends React.Component {
  static async getInitialProps({req}) {

    const res = await fetch('http://localhost:2428/messages');
    const data = await res.json();

    let messages = observable(data);
    return {
      messages: messages,
    };
  }

  render() {
    return <ListView messages={this.props.messages} />;
  }
}

export default App;
