import React from 'react'
import database from '../database'

function makeID() {
  let answer = '';
  for (let i = 0; i < 8; i++) {
    let index = Math.floor(Math.random() * 32)
    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index]
  }
  return answer;
}

export default class extends React.Component {
  static async getInitialProps({req}) {
    let id = makeID();
    database.push('messages', {
      id,
      content: 'this is message ' + id,
    })

    return {
      messages: database.all('messages')
    }
  }

  constructor(props) {
    super(props)
  }

  render() {
    this.props.messages.map(message => {
      console.log(message)
    })
    return (
      <div>
      Welcome to chat.
        <ul>
          {this.props.messages.map(message => (
            <li key={message.id}>{message.content}</li>
          ))}
        </ul>
      </div>
    )
  }
}
