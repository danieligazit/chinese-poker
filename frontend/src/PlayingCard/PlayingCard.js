import React, { Component, useRef} from 'react';
import { useDrag, useDrop } from 'react-dnd';
import './PlayingCard.css';
import PlayingCardsList from './PlayingCardsList';
import * as ReactDOM from "react-dom"; // Both at the same time

const type = "PlayingCard"



export const Card= ({ value, flipped, index, id }) => {
  const ref = useRef(null); // Initialize the reference
  
  // useDrop hook is responsible for handling whether any item gets hovered or dropped on the element
  const [, drop] = useDrop({
    // Accept will make sure only these element type can be droppable on this element
    accept: type,
    hover(item) {
      if (!ref.current) {
        return;
      }
      const dragIndex = item.index;
      // current element where the dragged element is hovered on
      const hoverIndex = index;
      // If the dragged element is hovered in the same place, then do nothing
      if (dragIndex === hoverIndex) { 
        return;
      }
      
      /*
        Update the index for dragged item directly to avoid flickering
        when the image was half dragged into the next
      */
      item.index = hoverIndex;
    }
  });
  // useDrag will be responsible for making an element draggable. It also expose, isDragging method to add any styles while dragging
  const [{ isDragging }, drag] = useDrag({
    // item denotes the element type, unique identifier (id) and the index (position)
    item: { type, id: id, index },
    // collect method is like an event listener, it monitors whether the element is dragged and expose that information
    collect: monitor => ({
      isDragging: monitor.isDragging()
    })
  });
  
  /* 
    Initialize drag and drop into the element using its reference.
    Here we initialize both drag and drop on the same element (i.e., Image component)
  */
  drag(drop(ref));
  
  return (
    <div
      ref={ref}
      className="file-item"
    >
    
                    
      <img 
        hidden={isDragging}
        alt={flipped === true ? 'Hidden Card' : PlayingCardsList[value]} 
        src={flipped === true ? PlayingCardsList.flipped : PlayingCardsList[value]} 
        className="file-img" />
    </div>
  );
  
}


const DraggablePlayingCard = ({ card, index, renderFunction }) => {
  const ref = useRef(null); // Initialize the reference
  
  // useDrop hook is responsible for handling whether any item gets hovered or dropped on the element
  const [, drop] = useDrop({
    // Accept will make sure only these element type can be droppable on this element
    accept: type,
    hover(item) {
      // if (!ref.current) {
      //   return;
      // }
      // const dragIndex = item.index;
      // // current element where the dragged element is hovered on
      // const hoverIndex = index;
      // // If the dragged element is hovered in the same place, then do nothing
      // if (dragIndex === hoverIndex) { 
      //   return;
      // }
      
      // /*
      //   Update the index for dragged item directly to avoid flickering
      //   when the image was half dragged into the next
      // */
      // item.index = hoverIndex;
    }
  });
  // useDrag will be responsible for making an element draggable. It also expose, isDragging method to add any styles while dragging
  const [{ isDragging }, drag] = useDrag({
    // item denotes the element type, unique identifier (id) and the index (position)
    item: { type, id: card.id, index },
    // collect method is like an event listener, it monitors whether the element is dragged and expose that information
    collect: monitor => ({
      isDragging: monitor.isDragging()
    })
  });
  
  /* 
    Initialize drag and drop into the element using its reference.
    Here we initialize both drag and drop on the same element (i.e., Image component)
  */
  drag(drop(ref));
  
  return renderFunction();
  
}

class PlayingCard extends Component {
  constructor(props){
    
    super(props);
    this.id = props.id
    this.state = {
      flipped : props.flipped || props.card === 'hide',
      card : props.card,
      height : props.height,
      flippable : props.flippable,
      elevated : props.elevated,
      style : this.props.style,
        position : {x : 0, y : 0},
        draggableDivStyle : {"zIndex":this.props.zIndex}
    }
  
  }
  componentWillReceiveProps(props) {
    this.setState({
        flipped : props.flipped,
        card : props.card,
        height : props.height,
        flippable : props.flippable,
        elevated : props.elevated,
        style : props.style,
        position : {x : 0, y : 0}

    })
  }
  elevate(percent){
    if(this.state.elevated) percent = -percent;
    let style = this.state.style;
    let translateY = style.transform.match(/translateY\((.*?)\)/); //pull out translateY
    if(translateY){
      let newTranslateY = Number(translateY[1].slice(0, -1)) - percent; //add 50%
      style.transform = style.transform.replace(/translateY(.*)/, `translateY(${newTranslateY}%)`)
    }else{
      style.transform += `\ntranslateY(${-percent}%)`
    }
    this.setState({style : style,
                    elevated : !this.state.elevated})
  }
  
  
  
  render() {

    return (
        
      <DraggablePlayingCard
          card = {this}
          index = {0}
          renderFunction = {
            () => {
              return (
                <div style={this.state.draggableDivStyle}>
                  <img ref={useRef(this.state.card)}
                    style={this.state.style}
                    height={this.state.height}
                    className='Playing-card'
                    src={this.state.flipped === true ? PlayingCardsList.flipped : PlayingCardsList[this.state.card]}
                    alt={this.state.flipped === true ? 'Hidden Card' : PlayingCardsList[this.state.card]}
                  />
                </div>
              )
            }
          }
          
        />
    )
  }
}

// /*this.state.flippable ? ()=> {
//             this.setState({flipped:this.state.flipped === true ? false : true,
//               height: this.state.height,
//               card: this.state.card});
//           } : null*/

export default PlayingCard;
