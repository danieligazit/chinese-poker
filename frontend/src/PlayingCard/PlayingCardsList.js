let PlayingCardsList = {};
let suits = ['c', 'd', 'h', 's'];
let faces = ['T', 'J', 'Q', 'K'];

let addSuits = (i, PlayingCardsList) => {
	for(let suit of suits){
		PlayingCardsList[i + suit] = require('./CardImages/' + i + suit + '.svg');
	}
}

for(let i = 1; i < 9; i++){
	addSuits(i, PlayingCardsList);
}

for(let i of faces){
	addSuits(i, PlayingCardsList);
}
PlayingCardsList.flipped = require('./CardImages/b.svg');
PlayingCardsList.nocard = require('./CardImages/png/nocard.png')

export default PlayingCardsList;