import React, {component, useState} from "react";
import { naturalWidth, naturalHeight } from './PlayingCard/Card'
import { Game } from "./Game"
const util = require('util');
const WebSocket = require('ws');

const baseURL = "5ed886257c7c46e28cdf3b5d46f59003.vfs.cloud9.us-east-1.amazonaws.com:8081"

class App extends React.Component {
  constructor(props) {
    super(props)
    this.ws = new WebSocket(`ws://${baseURL}/test?clientId=1`)
    this.ws.binaryType = 'arraybuffer';
    
    this.ws.on('open', () => {
      let data = Buffer.from(JSON.parse({"actionType": "connect"}).data)
      console.log(data)
    })
    
    this.ws.on('message', (data) => {
      var buf = new Uint8Array(data).buffer;
      var dec = new util.TextDecoder("utf-8");
      console.log(dec.decode(buf));
    })
    
    this.game = React.createRef();
  
  // handleClick = () => {
  //   this.game.current.setGameState({
  //       hands: [
  //         [
  //           ["1h", "3s", "4s"],
  //           ["Qs", "Kc"],
  //           ["Td", "Qc"],
  //           ["Jh", "3c", "Ts"],
  //           ["8d", "1h"]
  //         ],
  //         [
  //           ["nocard", "7s", "2h"],
  //           ["7c", "Kd", "1c"],
  //           ["nocard", "7d", "Qh"],
  //           ["nocard", "Ts", "Jh"],
  //           ["1c", "Qh", "Tc"]
  //         ]
  //       ],
  //       isCurrentTurn: true,
  //       topCard: "Qs",
  //       iteration: 3,
  //       playerIndex: 0
      
  //     }
  //   )
  //   setTimeout(() => {
  //     console.log('slept')
  //   this.game.current.setGameState({
  //       hands: [
  //         [
  //           ["1h", "3s", "4s"],
  //           ["Qs", "Kc", ],
  //           ["Td", "Qc"],
  //           ["Jh", "3c", "Ts"],
  //           ["8d", "1h"]
  //         ],
  //         [
  //           ["Js", "7s", "2h"],
  //           ["7c", "Kd", "1c"],
  //           ["nocard", "7d", "Qh"],
  //           ["nocard", "Ts", "Jh"],
  //           ["1c", "Qh", "Tc"]
  //         ]
  //       ],
  //       isCurrentTurn: true,
  //       topCard: "Qs",
  //       iteration: 3,
  //       playerIndex: 0
      
  //     }
  //   )
  //   }, 3000)
  }
  render(){
    return(
      <div onClick={this.handleClick}>
        <Game ref={this.game}/>
      </div>
    )
  }
}


export default App;
