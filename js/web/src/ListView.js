import React from 'react';
import { observer } from 'mobx-react';

import ChatInput from './ChatInput';

function makeID() {
  let answer = '';
  for (let i = 0; i < 8; i++) {
    let index = Math.floor(Math.random() * 32);
    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index];
  }
  return answer;
}

class ListView extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    console.log(
      'rendering messages:',
      JSON.stringify(this.props.messages));

    return (
      <div>
      Welcome to chat.
        <ul>
          {this.props.messages.map(message => (
            <li key={message.id}>{message.content}</li>
          ))}
        </ul>
        <ChatInput onSubmit={this.props.onSubmit} />
      </div>
    );
  }
}
ListView = observer(ListView);

export default ListView;
