import React from "react";
import { Game } from "./Game"
const util = require('util');
const WebSocket = require('ws');

const baseURL = "localhost:8081"

class App extends React.Component {
  constructor(props) {
    super(props)
    const ws = new WebSocket(`ws://${baseURL}/test?clientId=1`)
    ws.binaryType = 'arraybuffer';
    
    ws.on('open', () => {
      this.data = Buffer.from(JSON.parse({"actionType": "connect"}).data)
      
    })
    
    ws.on('message', (data) => {
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
        {this.data}
        <Game ref={this.game}/>
      </div>
    )
  }
}


export default App;
