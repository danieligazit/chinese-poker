import React from "react";
import { Game } from "./ChinesePoker"
const TextEncoder = require('text-encoder-lite').TextEncoderLite;
const TextDecoder = require('text-encoder-lite').TextDecoderLite;
const WS = require('ws');

const baseURL = "34.204.197.224:8081"
// const ws = new WebSocket("ws://3.237.28.197:8081/chinese-poker/bs7e5g9o3vkf67aasg9g?clientId=1")

async function messageToJson(message){
  return message.data.arrayBuffer()
    .then(buf => {
      return JSON.parse(String.fromCharCode.apply(null, new Uint8Array(buf)))     
    })
}

function jsonToMessage(json){
  return new TextEncoder('utf-8').encode(JSON.stringify(json))
}

export class ChinesePokerGame extends React.Component {
  constructor(props) {
    super(props)
    
    this.url = `ws://${baseURL}${props.endpoint}?clientId=${window.prompt("insert id (integer)")}`
    this.state = {
      active: false
    }
    this.game = React.createRef();
  }
  
  componentDidMount(){
    this.ws = new WebSocket(this.url)
    this.ws.onopen = () => {
      this.ws.send(jsonToMessage({"actionType": "connect"}))
      
    }
    
    this.ws.onmessage = (message) => {
      messageToJson(message)
        .then((data) => {
          if (data.actionType === "setState") {
            this.game.current.setGameState(data.action)
          } else if (data.actionType === "startGame"){
            this.setState({active: true})
          }
        })
    }
  }
  
  makeMove(move){
    this.ws.send(jsonToMessage({"actionType": "makeMove", "action": move}))
  }
  
  render(){
    return(
      <div onClick={this.handleClick}>
        {this.data}
        <Game ref={this.game} makeMove={this.makeMove.bind(this)} active={this.state.active}/>
      </div>
    )
  }
}


