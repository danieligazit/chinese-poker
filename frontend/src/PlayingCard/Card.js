import React from "react";
import { Draggable} from "react-beautiful-dnd";
import PlayingCardsList from './PlayingCardsList';

export const PlayingCard = ({value, index, isDragDisabled}) => {
  console.log("value", value)
  return (
    <Draggable
      key = {value}
      draggableId = {value}
      index = {index}
      isDragDisabled = {isDragDisabled}
    >
      {(provided, snapshot) => {
        return (
          <div
            ref={provided.innerRef}
            {...provided.draggableProps}
            {...provided.dragHandleProps}
            style={{
             
              userSelect: "none",
              
              ...provided.draggableProps.style
            }}
          >
            <img 
            style={{
              maxHeight: '100%',
              maxWidth: '100%'
            }}
            alt={PlayingCardsList[value]} 
            src={PlayingCardsList[value]} 
            className="file-img" />
          </div>
        );
      }}
    </Draggable>
  );
}