import React, { useState } from "react";
import { DragDropContext, Draggable, Droppable } from "react-beautiful-dnd";
import uuid from "uuid/v4";
import PlayingCardsList from './PlayingCard/PlayingCardsList';

const minColumnHeight = "213px";


const existingColumns = {
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
    item: { id: 'card1', value: "1h", seq: 1 }
  },
  "column2": {
    item: null
  },
  "column3": {
    item: null
  },
  "column4": {
    item: { id: 'card4', value: "jh", seq: 4 }
  },
  "column5": {
    item: null
  }
}


const origin = {
  item: { id: 'card7', value: "qs", seq: 7 }
}


const columns = {
  column1: currentTurnColumns.column1,
  column2: currentTurnColumns.column2,
  column3: currentTurnColumns.column3,
  column4: currentTurnColumns.column4,
  column5: currentTurnColumns.column5,
  origin: origin.item
}

const onDragEnd = (result, originColumn, setOriginColumn, columns, setColumns) => {
  if (!result.destination) return;
  const { source, destination } = result;
  console.log(source, destination)
  if (source.droppableId !== destination.droppableId) {
    const sourceColumn = columns[source.droppableId];
    console.log(columns)
    const destColumn = columns[destination.droppableId];
    const sourceItems = [sourceColumn.item];
    const destItems = [...destColumn.items];
    const [removed] = sourceItems.splice(source.index, 1);
    destItems.splice(destination.index, 0, removed);
    console.log(sourceItems)
    setColumns({
      ...columns,
      [source.droppableId]: {
        ...sourceColumn,
        items: sourceItems
      },
      [destination.droppableId]: {
        ...destColumn,
        items: destItems
      }
    });
  } else {

    const column = columns[source.droppableId];
    const copiedItems = [column.item];
    const [removed] = copiedItems.splice(source.index, 1);
    copiedItems.splice(destination.index, 0, removed);
    setColumns({
      ...columns,
      [source.droppableId]: {
        ...column,
        items: copiedItems
      }
    });
  }
};

function App() {
  
  const [columns, setColumns] = useState(currentTurnColumns);
  const [originColumn, setOriginColumn] = useState(origin); 
  return (
    <div style={{ display: "flex", justifyContent: "center", height: "100%" }}>
      <DragDropContext
        onDragEnd={result => onDragEnd(result, originColumn, setOriginColumn, columns, setColumns)}
      >
        {Object.entries(columns).map(([columnId, column], index) => {
          return (
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center"
              }}
              key={columnId}
            >
              <div style={{ margin: 8 }}>
                <Droppable droppableId={columnId} key={columnId} isDropDisabled={typeof column.item !== 'undefined' && column.item !== null}>
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
                          'min-height': minColumnHeight
                        }}
                      >
                        {column.item && 
                          
                          <Draggable
                            key={column.item.id}
                            draggableId={column.item.id}
                            index={0}
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
                                    'max-height': '100%',
                                    'max-width': '100%'
                                  }}
                                  alt={PlayingCardsList[column.item.value]} 
                                  src={PlayingCardsList[column.item.value]} 
                                  className="file-img" />
                                </div>
                              );
                            }}
                          </Draggable>
                          
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
        <break/>
        <Droppable droppableId='origin' key='origin' isDropDisabled='false'>
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
                  'min-height': minColumnHeight
                }}
              >
                {origin.item && 
                  
                  <Draggable
                    key={origin.item.id}
                    draggableId={origin.item.id}
                    index={0}
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
                            'max-height': '100%',
                            'max-width': '100%'
                          }}
                          alt={PlayingCardsList[origin.item.value]} 
                          src={PlayingCardsList[origin.item.value]} 
                          className="file-img" />
                        </div>
                      );
                    }}
                  </Draggable>
                  
                }
                {provided.placeholder}
              </div>
            );
          }}
        </Droppable>
      </DragDropContext>
    </div>
  );
}

export default App;
