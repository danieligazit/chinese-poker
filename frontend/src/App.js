import React, {Component} from 'react';
import {Card} from "./PlayingCard/PlayingCard";
import { BrowserRouter as Router, Route, Link, DefaultRoute} from "react-router-dom";
import './App.css';
import { TouchBackend } from "react-dnd-touch-backend";
import { HTML5Backend } from 'react-dnd-html5-backend';
import { DndProvider } from 'react-dnd';


const isTouchDevice = () => {
  if ("ontouchstart" in window) {
    return true;
  }
  return false;
};

// Assigning backend based on touch support on the device
const backendForDND = isTouchDevice() ? TouchBackend : HTML5Backend;


class App extends Component {
  
  render (){
    return (
      <div className="App">
        <header className="App-header">
          
          <DndProvider backend={backendForDND}>
            <Card
              value = {"1h"}
              id = {"some_id"}
              flippder = {false}
              index = {0}
              />
            <Card
              value = {"kh"}
              id = {"some_id"}
              flippder = {false}
              index = {0}
              />
        
     
          </DndProvider>
  
           
        </header>
        
      </div>
    );
  }
}

export default App

     //   <PlayingCard
          //         key={ "example" }
          //         height={ 300 }
          //         card={ "1h" }
          //         flipped={ true }
          //         elevateOnClick={true}
          //         id={0}
          //         index={0}
          //         // onDragStart={this.onDragStart.bind(this)}
          //         // onDragStop={this.onDragStop.bind(this)}
          //         // onDrag={this.onDrag.bind(this)}
          //         // onClick={this.onClick.bind(this)}
          //     />
          //   <PlayingCard
          //         key={ "example" }
          //         height={ 300 }
          //         card={ "kh" }
          //         flipped={ false }
          //         elevateOnClick={true}
          //         id={1}
          //         index={0}
          //         // onDragStart={this.onDragStart.bind(this)}
          //         // onDragStop={this.onDragStop.bind(this)}
          //         // onDrag={this.onDrag.bind(this)}
          //     />