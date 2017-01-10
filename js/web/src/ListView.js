import React from 'react';
import { observer } from 'mobx-react';

import ChatInput from './ChatInput';

class ListView extends React.Component {
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
