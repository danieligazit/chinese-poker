import React from "react";
import {ChinesePokerGame} from "./ChinesePoker/Game"

export default class Game extends React.Component{
  constructor(props){
    super(props)
    this.gameSelection = props.gameSelection
  }
  render() {
    switch (this.gameSelection){
      case 'chinese-poker':
        return (
          <ChinesePokerGame />
        )
      default:
        return <h1>oops</h1>
    }
  }
}
