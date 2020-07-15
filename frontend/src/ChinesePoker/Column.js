import React from "react";
import { DropTarget } from "react-dnd";
import {PlayingCard} from "./../PlayingCard/Card"
import Grid from '@material-ui/core/Grid';

export const Column = ({values, index, addable, originCardSetter, section}) => {
  return (
    <div style= {{
      flexDirection: "column",
      alignItems: "center",
      zoom: "1", 
      verticalAlign: "top",
      backgroundColor: "black",
      padding: 4,
      
    }}
    
    onClick = {() => {
      console.log(addable)
      // if (!addable) {return}
      originCardSetter(index)
    }}
    
    >
      <Grid container justify="center" spacing={2}>
      </Grid>
      {values.map((value, cardIndex) => {
        return (
          <PlayingCard value={value} index={index} key={section+"-column-"+index+"-card-"+cardIndex}/>
        )
      })}
      
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