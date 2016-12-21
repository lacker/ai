import React from 'react'
import { extendObservable } from 'mobx'
import { observer } from 'mobx-react'

function makeID() {
  let answer = ''
  for (let i = 0; i < 8; i++) {
    let index = Math.floor(Math.random() * 32)
    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index]
  }
  return answer
}

let store = {}

class ListView extends React.Component {
  static async getInitialProps({req}) {
    let id = makeID()
    let store = {}
    extendObservable(store, {
      messages: []
    })

    // The odd thing is that this adds a new message on each load
    store.messages.push({
      id,
      content: 'this is message ' + id,
    })

    return {
      messages: store.messages
    }
  }

  constructor(props) {
    super(props)
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
    )
  }
}
ListView = observer(ListView)

export default ListView
