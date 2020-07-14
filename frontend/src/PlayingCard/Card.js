import React from "react";
import { DragSource, DragPreviewImage } from "react-dnd";
import PlayingCardsList from './PlayingCardsList';

export const naturalWidth = 225
export const naturalHeight = 134
export const PlayingCard = ({state}) => {
  const cardImage = state.value? PlayingCardsList[state.value] : PlayingCardsList['b']
  return (
    <div>
      <DragPreviewImage src={cardImage} />
      <div>
        <img
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