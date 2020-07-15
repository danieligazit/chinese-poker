import React from "react";
import fetch from 'isomorphic-fetch';
import ProgressButton from 'react-progress-button'
import "../../node_modules/react-progress-button/react-progress-button.css";
import Game from './../Games'
import {BrowserRouter, Router, Route, Redirect} from "react-router-dom"

const baseURL = "localhost:8081"


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
    console.log('yay') 
  }
  
  handleClick() {
    if (this.state.lobbyCreated){ return this.handleLobbyCreated() }
    
    this.setState({buttonState: 'loading'})
  
    fetch(`https://${baseURL}/new/${this.gameSelection}`)
      .then(response => response.json())
      .then(data => this.setState({buttonState: 'success', text: 'go to lobby', lobbyURL: data.url, lobbyCreated: true}))
      .catch(() => this.setState({buttonState: 'error', }))
  }
  
  render () {
    if (this.state.redirect){
      return (
        <BrowserRouter>
          <Route exact path={this.state.lobbyURL} component={() => <Game gameSelection={this.state.gameSelection} lobby={this.state.lobbyURL}/>} />
          <Redirect to={this.state.lobbyURL}/>
        </BrowserRouter>
      )
    }
    return (

        
      <ProgressButton onClick={this.handleClick.bind(this)} state={this.state.buttonState}>
        {this.state.text}
      </ProgressButton>
    
    )
  }
}