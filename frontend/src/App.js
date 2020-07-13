import React, { useState } from "react";
import { DragDropContext, Draggable, Droppable } from "react-beautiful-dnd";
import uuid from "uuid/v4";
import PlayingCardsList from './PlayingCard/PlayingCardsList';
import { PlayingCard } from './PlayingCard/Card'

const minColumnHeight = "213px";


const state = {
  "hands": [
    ["1h"],
    [],
    [],
    ["jh"],
    []
  ],
  "isCurrentTurn": true,
  "topCard": "qs",
  "iteration": 1
}


const onDragEnd = (result, origin, setOrigin, columns, setColumns) => {
  if (!result.destination) return;
  const { source, destination } = result;
  
  if (destination.droppableId === source.droppableId) return;  
    
  const destColumn = columns[destination.index];
  const removed = origin;
  destColumn.splice(destination.index, 0, removed);
  setColumns({
    ...columns,
    destColumn
  });
  setOrigin(null)
  
};

export const Column = ({values, id, isDropDisabled, isDragDisabled}) => {
  return (
    <div style={{ margin: 8 }}>
      <Droppable droppableId={id} key={id} isDropDisabled={isDropDisabled}>
        {(provided, snapshot) => {
          return (
            <div
              {...provided.droppableProps}
              ref={provided.innerRef}
              style={{
                background: snapshot.isDraggingOver
                  ? "lightblue"
                  : "lightgrey",
                padding: 4,
                width: 150,
                minHeight: minColumnHeight
              }}
            >
              { values.map((value, index) => {
                console.log("columnValues", values)
                return (
                  <PlayingCard
                    value = {value}
                    index = {index}
                    isDragDisabled = {isDragDisabled}
                  />
                  
                )})
              }
              {provided.placeholder}
            </div>
          );
        }}
      </Droppable>
    </div>
  )
}
function App() {
  
  const [columns, setColumns] = useState(state.hands);
  const [origin, setOrigin] = useState(state.topCard); 
  return (
    <div style={{ display: "flex", justifyContent: "center", height: "100%" }}>
      <DragDropContext
        onDragEnd={result => onDragEnd(result, origin, setOrigin, columns, setColumns)}
      > 
        
        <div>
          {Object.entries(columns).map((values, index) => {
            return (
              <div
                style={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  display: "inline-block", 
                  zoom: "1", 
                  verticalAlign: "top"
                }}
                key={"column-"+ index}
                
              >
                <Column id={"column-" + index} values={values} isDropDisabled={false} isDragDisabled={true}/>
              </div>
                
            );
          })}
        </div>
        <div>
          <Column id='origin' values={origin ? [origin] : []} isDropDisabled = {true}m isDragDisabled={false}/>
        </div>
      </DragDropContext>
    </div>
  );
}

export default App;
