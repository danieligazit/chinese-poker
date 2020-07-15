import React from "react";
import { DropTarget } from "react-dnd";
import {PlayingCard} from "./../PlayingCard/Card"

export const Column = ({values, index, addable, originCardIndex, originCardState, originCardSetter, section}) => {
  return (
    <div style= {{
      flexDirection: "column",
      alignItems: "center",
      display: "inline-block", 
      zoom: "1", 
      verticalAlign: "top",
      backgroundColor: "black",
      padding: 4,
      
    }}
    
    onClick = {() => {
        if (originCardSetter){
          originCardSetter({
            ...originCardState,
            columnIndex: index
          })
        }
      }
    }
    >
      {values.map((value, cardIndex) => {
        return (
          <PlayingCard state={{columnIndex: index, value: value}} key={section+"-column-"+index+"-card-"+cardIndex}/>
        )
      })}
      {addable && originCardState.columnIndex == index && <PlayingCard state={originCardState} key={"origin-"+index} />}
    </div>
  )
} 


export default DropTarget(
  "PlayingCard",
  {
    canDrop: (props) => true,
    drop: (props) => {
    }
  },
  (connect, monitor) => {
    return {
      connectDropTarget: connect.dropTarget(),
      isOver: !!monitor.isOver(),
      canDrop: !!monitor.canDrop()
    };
  }
)(Column)