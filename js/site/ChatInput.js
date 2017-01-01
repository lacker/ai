import React from 'react';

export default class ChatInput extends React.Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.onSubmit = props.onSubmit;

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    alert('A chat was submitted: ' + this.state.value);
    event.preventDefault();
    this.onSubmit(this.state.value);
    this.setState({value: ''});
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Name:
          <input type="text" value={this.state.value} onChange={this.handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
    );
  }
}
