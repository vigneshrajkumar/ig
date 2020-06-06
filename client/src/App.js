import React from 'react';

import {
  BrowserRouter,
  Switch,
  Route
} from 'react-router-dom'

import FeedPage from "./components/FeedPage"
import "./app.css"

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/"><FeedPage /></Route>
      </Switch>
    </BrowserRouter>
    )
}

export default App;
