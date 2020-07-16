import React from "react";
import fetch from 'isomorphic-fetch';
import ProgressButton from 'react-progress-button'
import "../../node_modules/react-progress-button/react-progress-button.css";
import Game from './../Games'
import {Redirect} from "react-router-dom"

const baseURL = "3.235.139.208:8081"


export default class CreateLobby extends React.Component {
  
  constructor(props) {
    super(props)
    this.state = {
      buttonState: '',
      text: 'create lobby',
      lobbyCreated: false,
      lobbyURL: '',
      redirect: false,
      gameSelection: props.gameSelection
    }
  }
  
  handleLobbyCreated() {
    this.setState({redirect: true})
  }
  
  handleClick() {
    if (this.state.lobbyCreated){ return this.handleLobbyCreated() }
    
    this.setState({buttonState: 'loading'})
  
    fetch(`http://${baseURL}/new/${this.state.gameSelection}`)
      .then(response => response.json())
      .then(data => this.setState({buttonState: 'success', text: 'go to lobby', lobbyURL: data.url, lobbyCreated: true}))
      .catch(() => this.setState({buttonState: 'error', }))
  }
  
  render () {
    if (this.state.redirect){
      return (

          
        <Redirect to={this.state.lobbyURL}/>

      )
    }
    return (

        
      <ProgressButton onClick={this.handleClick.bind(this)} state={this.state.buttonState}>
        {this.state.text}
      </ProgressButton>
    
    )
  }
}