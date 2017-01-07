import React from 'react';
import { observable } from 'mobx';
import 'isomorphic-fetch';

import ListView from './ListView';
import MessageClient from './MessageClient';

class App extends React.Component {
  constructor(props) {
    super(props);

    this.client = new MessageClient();

    this.state = {
      loading: true,
    };

    // TODO: check if this works
    this.client.load().then(() => {
      this.setState({
        messages: this.client.messages,
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
