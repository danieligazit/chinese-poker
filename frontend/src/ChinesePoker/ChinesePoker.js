import React  from "react";
import { Column } from './Column';
import { naturalWidth, naturalHeight, PlayingCard } from './../PlayingCard/Card'
import { DndProvider } from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend'

const cardImageRatio = naturalHeight / naturalWidth

export class Game extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      gameState: {},
    }
  }  
  
  setGameState(newState){
    let playerIndex = (newState.playerIndex + 1) % 2
    
    newState.hands[playerIndex].map((hand, handIndex) => {
      while (hand.length < newState.iteration) {
        hand.push("nocard")
      }
      hand = hand.reverse()
    })
    
    this.setState({
      gameState: newState,
      originCardState: {
        ...this.state.gameState.originCardState,
        value: newState.top ? newState.top : "nocard",
        columnIndex: -1,
      }
    })
  }
  
  
  renderMy(){
    if (!this.state.gameState.hands){
      return
    }
    console.log(this.state.gameState.hands)
    const hands = this.state.gameState.hands[this.state.gameState.playerIndex]
    console.log("active", this.props.active)
    return hands.map((hand, index) => {
      return (  
        <div style = {{width: "25%", maxWidth:naturalWidth}}>
          <Column 
            values={hand} 
            key={"column-"+index} 
            index={index} 
            section="current"
            addable={this.props.active && hand.length === this.state.gameState.iteration -1 }
            originCardSetter={this.setOriginCard.bind(this)}
          />
        </div>
      ) 
    })
  }

  renderOponenet() {
    if (!this.state.gameState.hands) {return}
    
    return this.state.gameState.hands[(this.state.gameState.playerIndex + 1) % 2].map((hand, index) => {
      return (
        <div style = {{width: "25%", maxWidth:naturalWidth}}>
          <Column 
            values={hand}
            key={"openent-column-"+index} 
            index={index}
            addable={false} 
            section="oponent"
          />
        </div>
      )
    })
  }

  setOriginCard(newColumnIndex){
    if (!this.props.active){return}
    if( this.state.gameState.columnIndex === "nocard"){return}
    
    
    this.setState(prevState => {
      const prev = prevState.gameState
      let newHands = prev.hands
      newHands[prev.playerIndex][newColumnIndex] = [...newHands[prev.playerIndex][newColumnIndex], prev.top] 
      return {
        gameState: {
          ...prev,
          hands: newHands,
          top: "nocard"
        }
      }
    })
    
    this.props.makeMove({"handIndex": newColumnIndex})
  }
  
  render() {return (
    <div style = {{minWidth:naturalWidth/2 * 5}}>
      <DndProvider backend={HTML5Backend}>
        <div  style={{
          display: "flex", 
          justifyContent: "center", 
        }}>
          {this.renderOponenet()}
        </div>
        <div  style={{
          display: "flex", 
          justifyContent: "center", 
        }}>
          {this.renderMy()}
        </div>
        <div  style={{
          display: "flex", 
          justifyContent: "center", 
          height: "100%",
          backgroundColor: "green"
        }}>
          
          {<PlayingCard key={"origin"} value={this.state.gameState.top}/>}
        </div>
        
      </DndProvider>
    </div>
  )}
}
