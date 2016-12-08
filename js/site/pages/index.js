import React from 'react'
import database from './database'

function makeID() {
  let answer = '';
  for (let i = 0; i < 8; i++) {
    let index = Math.floor(Math.random() * 32
    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index]
  }
  return answer;
}

export default class extends React.Component {
  static async getInitialProps({req}) {
    let id = makeID();
    database.push({
      id,
      content: 'this is message ' + id,
    })

    return {
      messages: database.all('messages')
    }
  }

  render() {
    return (
      <ul>
        {this.props.messages.map(message => {
          <li key={message.id}>{message.content}</li>
        })}
      </ul>
    )
  }
}
