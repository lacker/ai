import React from 'react';
import { extendObservable } from 'mobx';
import { observer } from 'mobx-react';
import 'isomorphic-fetch';

import ChatInput from '../ChatInput';

function makeID() {
  let answer = '';
  for (let i = 0; i < 8; i++) {
    let index = Math.floor(Math.random() * 32);
    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index];
  }
  return answer;
}

let store = {};
extendObservable(store, {
  messages: [],
});

class ListView extends React.Component {
  static async getInitialProps({req}) {

    const res = await fetch('http://localhost:2428/messages');
    const data = await res.json();
    store.messages = data;
    return {
      messages: store.messages,
    };
  }

  constructor(props) {
    super(props);

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(value) {
    console.log(
      'at the start of handleSubmit messages is:',
      JSON.stringify(this.props.messages));

    // TODO: use chat server instead of just locally doing stuff
    // TODO: figure out how to make this rerender the ListView
    this.props.messages.push({
      id: makeID(),
      content: value,
    });
    
    console.log(
      'at the end of handleSubmit messages is:',
      JSON.stringify(this.props.messages));
  }

  render() {
    console.log('rendering messages:', JSON.stringify(this.props.messages));
    return (
      <div>
      Welcome to chat.
        <ul>
          {this.props.messages.map(message => (
            <li key={message.id}>{message.content}</li>
          ))}
        </ul>
        <ChatInput onSubmit={this.handleSubmit} />
      </div>
    );
  }
}
ListView = observer(ListView);

export default ListView;
