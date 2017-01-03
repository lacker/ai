import React from 'react';
import { observable } from 'mobx';
import 'isomorphic-fetch';

import ListView from './ListView';

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      loading: true,
    };

    fetch('http://localhost:2428/messages').then(res => {
      return res.json();
    }).then(data => {
      let messages = observable(data);
      this.setState({
        loading: false,
        messages: messages,
      });
    });
  }

  render() {
    if (this.state.loading) {
      return <div>Loading...</div>;
    }

    return <ListView messages={this.state.messages} />;
  }
}

export default App;
