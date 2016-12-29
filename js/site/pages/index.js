import React from 'react';
import { extendObservable } from 'mobx';
import { observer } from 'mobx-react';
const rp = require('request-promise');

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

    const options = {
      uri: 'http://localhost:2428/messages',
      json: true,
    };

    return rp(options).then((messages) => {
      store.messages = messages;
      return {
        messages: store.messages,
      };
    });
  }

  constructor(props) {
    super(props);
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
      </div>
    );
  }
}
ListView = observer(ListView);

export default ListView;
