import React from "react";
import {ChinesePokerGame} from "./ChinesePoker/Game"

export default class Game extends React.Component{
  constructor(props){
    super(props)
  }
  render() {
    switch (this.props.match.params.game){
      case 'chinese-poker':
        return (
          <ChinesePokerGame endpoint={"/"+this.props.match.params.game+"/"+this.props.match.params.lobbyId}/>
        )
      default:
        return <h1>oops</h1>
    }
  }
}
