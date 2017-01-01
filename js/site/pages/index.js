import React from 'react';
import { extendObservable } from 'mobx';
import { observer } from 'mobx-react';
import 'isomorphic-fetch';

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

// TODO: try running this with the chat server going, see if it works
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

    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(e) {
    console.log('button pressed with event:', e);
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
        <button onClick={this.handleClick} title='Add Chat' />
      </div>
    );
  }
}
ListView = observer(ListView);

export default ListView;
