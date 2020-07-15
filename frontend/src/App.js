import React from "react";
import CreateLobby from "./MainPage/CreateLobby"
import {BrowserRouter, Route} from "react-router-dom"
import Game from "./Games"
class App extends React.Component {
  
  constructor(props){
    super(props)
    this.state =  {
      
      gameSelection: "chinese-poker"
    }
  }
  
  
  render() {
    return (
      <BrowserRouter>
        <Route exact path={"/:game/:lobbyId"} component={Game} />
        <Route exact path={"/"} component={() => <CreateLobby gameSelection={this.state.gameSelection}/>}/>
        
      </BrowserRouter>
    )
  }
}


export default App;
