import React  from "react";
import { Column } from './Column';
import { naturalWidth, naturalHeight, PlayingCard } from './../PlayingCard/Card'
import { motion } from "framer-motion"

export class Game extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      gameState: {},
      winners: [],
      evaluations: [],
      gameOver: false,
    }
  }  
  
  setGameState(newState){
    newState.hands.map((playerHands, playerIndex) => {
      playerHands.map((hand, handIndex) => {
        while (hand.length < newState.iteration) {
          hand.push("nocard")
        }
        if (playerIndex !== newState.playerIndex) {hand.reverse()}
      })
    })
    
    this.setState({
      ...this.state,
      gameState: newState,
    })
  }
  
  setGameResult(gameResult){
    this.setState({
      ...this.state,
      evaluations: gameResult.evaluations,
      winners: gameResult.winners,
      gameOver: true
    })
  }
  
  renderMy(){
    if (!this.state.gameState.hands){
      return
    }
    const myIndex = this.state.gameState.playerIndex
    const hands = this.state.gameState.hands[myIndex]
    return hands.map((hand, index) => {
      return (
        <div style = {{width: "20%", maxWidth:naturalWidth}} key={"column-"+index} >
          {this.state.gameOver &&
            <div style = {{fontSize: 22, backgroundColor: 'black', textAlign: 'center', color: this.state.winners[index].includes(myIndex) ? 'green' : 'red'}}>
              {this.state.evaluations[myIndex][index].rankStr}
            </div>
          }
          <Column 
            values={hand}
            index={index} 
            section="current"
            addable={this.props.active && hand.filter(x => x !== "nocard").length === this.state.gameState.iteration -1 && this.state.gameState.top !== "nocard"}
            originCardSetter={this.setOriginCard.bind(this)}
            iteration = {this.state.gameState.iteration}
          />
         
        </div>
      ) 
    })
  }

  renderOponenet() {
    if (!this.state.gameState.hands) {return}
    
    const oponentIndex = (this.state.gameState.playerIndex + 1) % 2
    return this.state.gameState.hands[oponentIndex].map((hand, index) => {
      return (
        <div style = {{width: "20%", maxWidth:naturalWidth}}  key={"oponent-column-"+index} >
          
          <Column 
            values={hand}
            index={index}
            addable={false} 
            section="oponent"
            iteration = {this.state.gameState.iteration}
          />
          {this.state.gameOver &&
            <div style = {{fontSize: 22, backgroundColor: 'black', textAlign: 'center', color: this.state.winners[index].includes(oponentIndex) ? 'green' : 'red'}}>
              {this.state.evaluations[oponentIndex][index].rankStr}
            </div>
          }
        </div>
      )
    })
  }

  setOriginCard(newColumnIndex){
     this.props.makeMove({"handIndex": newColumnIndex})
  }
  
  render() {
    return (
    
    <div style = {{minWidth:naturalWidth/2 * 5}}>
      <div  style={{
        display: "flex", 
        height: "100%",
        justifyContent: "center", 
      }}>
        {this.renderOponenet()}
      </div>
      <div  style={{
        display: "flex", 
        height: "100%",
        justifyContent: "center", 
      }}>
        {this.renderMy()}
      </div>
      <div  style={{
        display: "flex", 
        justifyContent: "center", 
        height: "800px",
        backgroundColor: "green",
        overflow: "hidden"
      }}>
        {<PlayingCard style={{position: "absolute"}} key={"origin"} value={this.state.gameState.top}/>}
      </div>
        
    </div>
  )}
}
