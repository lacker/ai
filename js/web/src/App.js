import React from 'react';

import ListView from './ListView';
import MessageClient from './MessageClient';

class App extends React.Component {
  constructor(props) {
    super(props);

    this.client = new MessageClient();

    this.state = {
      loading: true,
    };

    this.handleSubmit = this.handleSubmit.bind(this);

    // TODO: check if this works
    this.client.load().then(() => {
      this.setState({
        loading: false,
      });
    });
  }

  handleSubmit(content) {
    this.client.create(content);
  }

  render() {
    if (this.state.loading) {
      return <div>Loading...</div>;
    }

    return <ListView
             onSubmit={this.handleSubmit}
             messages={this.client.messages} />;
  }
}

export default App;
