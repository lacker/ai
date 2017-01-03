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
        <ChatInput onSubmit={this.handleSubmit} />
      </div>
    );
  }
}
ListView = observer(ListView);

export default ListView;
