import React from "react";
import fetch from 'isomorphic-fetch';
import ProgressButton from 'react-progress-button'

const baseURL = "localhost:8081"


class CreateLobby extends React.Component {
  getInitialState() {
    return {
      buttonState: '',
      lobbyCreated: false
    }
  }
  
  constructor(props) {
    super(props)
    this.state = this.getInitialState()
  }
  
  handleLobbyCreated() {
    console.log('yay')  
  }
  
  handleClick() {
    if (this.state.lobbyCreated){ return this.handleLobbyCreated() }
    
    this.setState({buttonState: 'loading'})
    
    fetch(`https://${baseURL}/new/${this.gameSelection}`)
      .then(response => response.json())
      .then(data => this.setState({buttonState: 'success', lobbyCreated: true}))
      .catch(() => this.setState({buttonState: 'error'}))
  }
  
  render () {
    return (
      <ProgressButton onClick={this.handleClick.bind(this)} state={this.state.buttonState}>
        CreateLobby
      </ProgressButton>
    )
  }
}

class App extends React.Component {
  
  getInitialState() {
    return {
      gameSelection: "poker"
    }
  }
  
  handleClick(){
    
  }
  
  render() {
    return (
      <div>
        <CreateLobby />
      </div>
    )
  }
}


export default App;
