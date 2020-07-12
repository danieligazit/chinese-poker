import React, { useState } from "react";
import { DragDropContext, Draggable, Droppable } from "react-beautiful-dnd";
import uuid from "uuid/v4";
import PlayingCardsList from './PlayingCard/PlayingCardsList';

const minColumnHeight = "213px";


const columns = {
  "column1": {
    droppable: false,
    items: [{ id: 'card1', value: "1h", seq: 1 }]
  },
  "column2": {
    droppable: true,
    items: [{ id: 'card5', value: "8c", seq: 5 }]
  },
  "column3": {
    droppable: false,
    items: [{ id: 'card6', value: "kh", seq: 6 }]
  },
  "column4": {
    droppable: true,
    items: [{ id: 'card4', value: "jh", seq: 4 }]
  },
  "column5": {
    droppable: true,
    items: [{ id: 'card3', value: "ks", seq: 3 }]
  }
};


const currentTurnColumns= {
  "column1": {
    items: [{ id: 'card1', value: "1h", seq: 1 }]
  },
  "column2": {
    items: []
  },
  "column3": {
    items: []
  },
  "column4": {
    items: [{ id: 'card4', value: "jh", seq: 4 }]
  },
  "column5": {
    items: []
  }
}


const origin = {
  items: [{ id: 'card7', value: "qs", seq: 7 }]
}


const onDragEnd = (result, originColumn, setOriginColumn, columns, setColumns) => {
  if (!result.destination) return;
  const { source, destination } = result;
  
  if (destination.droppableId === source.droppableId) return;  
  
  const destColumn = columns[destination.droppableId];
  const destItems = [...destColumn.items];
  console.log(originColumn.items[source.index])
  const [removed] = originColumn.items.splice(source.index, 1);
  console.log(destItems)
  destItems.splice(destination.index, 0, removed);
  console.log(destItems)
  console.log(destination.droppableId)
  setColumns({
    ...columns,
    [destination.droppableId]: {
      ...destColumn,
      items: destItems
    }
  });
  setOriginColumn(originColumn)
  
};

function App() {
  
  const [columns, setColumns] = useState(currentTurnColumns);
  const [originColumn, setOriginColumn] = useState(origin); 
  return (
    <div style={{ display: "flex", justifyContent: "center", height: "100%" }}>
      <DragDropContext
        onDragEnd={result => onDragEnd(result, originColumn, setOriginColumn, columns, setColumns)}
      >
        <div>
          {Object.entries(columns).map(([columnId, column], index) => {
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
                key={columnId}
                
              >
                <div style={{ margin: 8 }}>
                  <Droppable droppableId={columnId} key={columnId} isDropDisabled={column.items.length !== 0}>
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
                          {column.items.map((item, index) => {
                            return (
                              <Draggable
                                key={item.id}
                                draggableId={item.id}
                                index={index}
                                isDragDisabled={true}
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
                                      alt={PlayingCardsList[item.value]} 
                                      src={PlayingCardsList[item.value]} 
                                      className="file-img" />
                                    </div>
                                  );
                                }}
                              </Draggable>
                            )})
                          }
                          {provided.placeholder}
                        </div>
                      );
                    }}
                  </Droppable>
                </div>
              </div>
            );
          })}
        </div>
        <div>
        <Droppable droppableId='origin' key='origin' isDropDisabled={false}>
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
                {origin.items.map((item, index) => { 
                  return (
                    <Draggable
                      key={item.id}
                      draggableId={item.id}
                      index={index}
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
                            alt={PlayingCardsList[item.value]} 
                            src={PlayingCardsList[item.value]} 
                            className="file-img" />
                          </div>
                        );
                      }}
                    </Draggable>
                  )
                })}
                {provided.placeholder}
              </div>
            );
          }}
        </Droppable>
        </div>
      </DragDropContext>
    </div>
  );
}

export default App;
