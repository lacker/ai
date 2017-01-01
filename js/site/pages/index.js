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
  messages: []
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
  }

  handleSubmit(value) {
    // TODO: use chat server instead of just locally doing stuff
    store.messages.push({
      id: makeID(),
      content: value,
    });
  }

  render() {
    return (
      <div>
      Welcome to chat.
        <ul>
          {this.props.messages.map(message => (
            <li key={message.id}>{message.content}</li>
          ))}
        </ul>
        <ChatInput onSubmit={this.handleClick} />
      </div>
    );
  }
}
ListView = observer(ListView);

export default ListView;
