

export let currentGameState = {}

export function stateSet(){
  window.onclick = () => {
    currentGameState = {
      hands: [
        [
          ["1h", "3s", "4s"],
          ["Qs", "Kc"],
          ["Td", "Qc"],
          ["Jh", "3c", "Ts"],
          ["8d", "1h"]
        ],
        [
          ["nocard", "7s", "2h"],
          ["7c", "Kd", "1c"],
          ["nocard", "7d", "Qh"],
          ["nocard", "Ts", "Jh"],
          ["1c", "Qh", "Tc"]
        ]
      ],
      isCurrentTurn: true,
      topCard: "Qs",
      iteration: 3,
      playerIndex: 0
    
    }
  }
}

function getOponentHands(gameState){
  let playerIndex = (gameState.playerIndex + 1) % 2
  let hands = gameState.hands[playerIndex]
  
  hands.map((hand, index) => {
    while (hand.length < gameState.iteration) {
      hand.push("nocard")
    }
    console.log(hand)
    hand = hand.reverse()
    console.log(hand)
  })
  
  return hands
}