import React from "react";
import { DragSource, DragPreviewImage } from "react-dnd";
import PlayingCardsList from './PlayingCardsList';
import "./PlayingCard.css"
import "./../Column.css"

export const naturalWidth = 225
export const naturalHeight = 314


export const PlayingCard = ({value, index=0, section, text}) => {
  if (!value){value = "flipped"}
  return (
    <div
    className="Playing-card" style = {{zIndex: section === "oponent" ? 10 - index : index, height: naturalHeight}}>
      {value !== "nocard" && 
      <img draggable={false}
        src = {PlayingCardsList[value]}
      />}
    </div>
    
  );
  
  
}