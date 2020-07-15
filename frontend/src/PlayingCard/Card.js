import React from "react";
import { DragSource, DragPreviewImage } from "react-dnd";
import PlayingCardsList from './PlayingCardsList';
import "./PlayingCard.css"
export const naturalWidth = 225
export const naturalHeight = "134px"

export const PlayingCard = ({value}) => {
  const cardImage = value ? PlayingCardsList[value] : PlayingCardsList['b']
  return (
    <div>
      <DragPreviewImage src={PlayingCardsList[value]} />
      <div>
        <img
          className="Playing-card"
          style = {{
            width: "100%"
          }}
          src = {cardImage}
        />
      </div>
    </div>
  );
}

export default DragSource(
  "PlayingCard",
  {
    beginDrag: () => ({}),
  },
  (connect, monitor) => ({
    connectDragSource: connect.dragSource(),
    connectDragPreview: connect.dragPreview(),
    isDragging: monitor.isDragging(),
  }),
)(PlayingCard)