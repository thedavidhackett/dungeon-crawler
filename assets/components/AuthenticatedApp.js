import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
  Link,
} from "react-router-dom";
import Navigation from "./Navigation";
import Home from "./Home";
import CreateDungeon from "./CreateDungeon";
import Dungeon from "./Dungeon";

export default function AuthenticatedApp() {
  return (
    <Router>
      <div>
        <Navigation />
        <Switch>
          <Route path="/" exact component={Home} />
          <Route path="/create-dungeon" exact component={CreateDungeon} />
          <Route path="/dungeon/:id" exact component={Dungeon} />
          <Redirect to="/" />
        </Switch>
      </div>
    </Router>
  );
}
