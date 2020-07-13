import React from "react";
import { Draggable} from "react-beautiful-dnd";
import PlayingCardsList from './PlayingCardsList';

export const PlayingCard = ({value, id, index, isDragDisabled}) => {
  return (
    <Draggable
      key = {id}
      draggableId = {id}
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