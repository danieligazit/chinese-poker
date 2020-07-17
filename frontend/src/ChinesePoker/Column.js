import React from "react";
import { DropTarget } from "react-dnd";
import {PlayingCard} from "./../PlayingCard/Card"
import { SpringGrid, layout } from 'react-stonecutter';
import "./../Column.css"
const cardHeight = 314
const cardWidth = 225

export const Column = ({values, index, addable, originCardSetter, section, iteration}) => {
  return (
    <div style= {{
      flexDirection: "column",
      alignItems: "center",
      zoom: "1", 
      verticalAlign: "top",
      height: "auto",
      backgroundColor: "black",
      padding: 4,
    }}
    
    className = {section==="oponent"? 'vhand-compact-reversed': 'vhand-compact'}
    
    onClick = {() => {
      if (!addable) {return}
      originCardSetter(index)
    }}
    >

      {values.map((value, cardIndex) => (
          <PlayingCard value={value} index={cardIndex} section={section}/>
        )
      )}
   
      
    </div>
  )
} 
