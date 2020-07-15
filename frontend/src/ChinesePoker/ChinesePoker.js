import React  from "react";
import { Column } from './Column';
import { naturalWidth, naturalHeight, PlayingCard } from './../PlayingCard/Card'
import { DndProvider } from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend'

const cardImageRatio = naturalHeight / naturalWidth



function renderOponenet(gameState) {
  if (!gameState.hands) {
    return
  }
  
  return gameState.hands[(gameState.playerIndex + 1) % 2].map((hand, index) => {
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

function renderMy(gameState, originCardState, setOriginCardState){
  if (!gameState.hands){
    return
  }
  return gameState.hands[gameState.playerIndex].map((hand, index) => {
    return (  
      <div style = {{width: "25%", maxWidth:naturalWidth}}>
        <Column 
          values={hand} 
          key={"column-"+index} 
          index={index} 
          section="current"
          addable={gameState.iteration > gameState.hands[gameState.playerIndex][index].length} 
          originCardState={originCardState} 
          originCardSetter={setOriginCardState}
        />
      </div>
    ) 
  })
}


export class Game extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      gameState: {},
      originCardState: {columnIndex: -1, value: "nocard"}
    }
  }  
  
  setGameState(newState){
    this.setState({
      ...this.state,
      gameState: newState,
      originCardState: {
        ...this.state.originCardState,
        value: newState.top ? newState.top : "nocard"
      }
    })
  }
  
  setOriginCardState(newColumnIndex){
    this.setState({
      ...this.state,
      originCardState: {
        ...this.state.originCardState,
        columnIndex: newColumnIndex
      }
    })
  }
  
  setCardColumn(columnIndex) {
    this.state.originCardState.columnIndex = columnIndex
  }
  
  render() {return (
    <div style = {{minWidth:naturalWidth/2 * 5}}>
      <DndProvider backend={HTML5Backend}>
        <div  style={{
          display: "flex", 
          justifyContent: "center", 
          height: "100%",
        }}>
          {renderOponenet(this.state.gameState)}
        </div>
        <div  style={{
          display: "flex", 
          justifyContent: "center", 
          height: "100%"
        }}>
          {renderMy(this.state.gameState, this.state.originCardState, this.setOriginCardState)}
        </div>
        <div  style={{
          display: "flex", 
          justifyContent: "center", 
          height: "100%",
          backgroundColor: "green"
        }}>
          
          {this.state.originCardState.columnIndex == -1 && <PlayingCard key={"origin"} state={this.state.originCardState}/>}
          {this.state.originCardState.columnIndex != -1 && <PlayingCard key={"origin"} state={{cardPosition: -1, value: 'nocard'}}/>}
        </div>
        
      </DndProvider>
    </div>
  )}
}
