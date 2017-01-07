import { observable } from 'mobx';
import 'isomorphic-fetch';

export default class MessageClient {
  constructor() {
    this.messages = observable([]);
  }

  // TODO: make this multi-entrant
  load() {
    return fetch('http://localhost:2428/messages').then(res => {
      return res.json();
    }).then(data => {
      this.messages.replace(data);
    });
  }
}
