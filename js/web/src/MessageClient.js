import { observable } from 'mobx';
import 'isomorphic-fetch';

export default class MessageClient {
  constructor() {
    // Whether the client is in the initial pre-data loading state
    this.loading = true;

    this.messages = observable([]);

    this.load();
  }

  // TODO: make this multi-entrant
  // TODO: make the react stuff use MessageClient
  load() {
    return fetch('http://localhost:2428/messages').then(res => {
      return res.json();
    }).then(data => {
      this.loading = false;
      this.messages.replace(data);
    });
  }
}
