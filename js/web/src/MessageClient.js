import { observable } from 'mobx';
import 'isomorphic-fetch';

function makeID() {
  let answer = '';
  for (let i = 0; i < 8; i++) {
    let index = Math.floor(Math.random() * 32);
    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index];
  }
  return answer;
}

export default class MessageClient {
  // TODO: also kick off a websocket listen
  constructor() {
    this.messages = observable([]);
  }

  // TODO: see if this works
  load() {
    return fetch('http://localhost:2428/messages').then(res => {
      return res.json();
    }).then(data => {
      let newMessages = {};
      for (let message of this.messages) {
        newMessages[message.id] = message;
      }
      for (let message of data) {
        newMessages[message.id] = message;
      }
      newMessages.sort((m1, m2) => (m1.timestamp - m2.timestamp));
      this.messages.replace(newMessages);
    });
  }

  // Creates a new message
  // TODO: make this post to the chat server
  create(content) {
    this.messages.push({
      id: makeID(),
      content: content,
      timestamp: (new Date()).getTime(),
    });
  }
}
