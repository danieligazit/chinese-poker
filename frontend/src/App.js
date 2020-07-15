import React from "react";
import CreateLobby from "./MainPage/CreateLobby"


class App extends React.Component {
  
  constructor(props){
    super(props)
    this.state =  {
      gameSelection: "chinese-poker"
    }
  }
  
  
  render() {
    return (
      <div>
        <CreateLobby gameSelection={this.state.gameSelection}/>
      </div>
    )
  }
}


export default App;
